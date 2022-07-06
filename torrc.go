// Package torrc provides a simple way to read, edit, and write torrc files.
//
// It takes a torrc file and parses it into a TorRC structure, which can be
// manipulated and then written back to a file.
//
// Example(set up an I2P Bridge Client):
//
//	torrc, err := ReadTorRC("torrc.test")
//	if err != nil {
//		panic(err)
//	}
//	torrc.SetField("UseBridges", []string{"1"})
//	torrc.SetField("Bridge", []string{"i2p", "i2pbase32addressesarefiftytwocharacterslongenoughok.b32.i2p"})
//	torrc.SetField("ClientTransportPlugin", []string{"i2p", "exec", "i2p-client"})
//	if err = torrc.Write("torrc.test.i2p-client"); err != nil {
//		panic(err)
//	}
//
// Example(set up an I2P Bridge Server):
//
//	torrc, err := ReadTorRC("torrc.test")
//	if err != nil {
//		panic(err)
//	}
//	torrc.SetField("BridgeRelay", []string{"1"})
//	torrc.SetField("ORPort", []string{"9001"})
//	torrc.SetField("ExtORPort", []string{"7689"})
//	torrc.SetField("ServerTransportPlugin", []string{"i2p", "exec", "i2p-server"})
//	if err = torrc.Write("torrc.test.i2p-server"); err != nil {
//		panic(err)
//	}
package torrc

import (
	"io/ioutil"
	"strings"
)

// TorRC is a structure that holds the contents of a torrc file.
type TorRC struct {
	fields []*Field
}

// GetFieldIndex returns the index of the field with the given name.
func (t *TorRC) GetFieldIndex(name string) int {
	for i, f := range t.fields {
		if f.Compare(name) {
			return i
		}
	}
	return -1
}

// GetField returns the value of the field with the given name.
func (t *TorRC) GetField(name string) *Field {
	index := t.GetFieldIndex(name)
	if index == -1 {
		return nil
	}
	return t.fields[index]
}

// GetFields returns a slice of all the fields in the TorRC
func (t *TorRC) GetFields() []*Field {
	return t.fields
}

// GetFieldsByName returns a slice of all the fields in the TorRC matching
// the given name. This includes all commented-out fields.
func (t *TorRC) GetFieldsByName(name string) []*Field {
	var fields []*Field
	for _, f := range t.fields {
		if f.Compare(name) {
			fields = append(fields, f)
		}
	}
	return fields
}

// SetField sets the value of a field in the TorRC structure. If the field
// is commented out or non-existent, it places a new entry at the end of the
// torrc file.
func (t *TorRC) SetField(name string, value []string) {
	index := t.GetFieldIndex(name)
	if index == -1 {
		t.fields = append(t.fields, NewField(name, value))
	} else {
		if t.fields[index].IsComment() {
			// if the field is a comment, we need to append, not set
			t.fields = append(t.fields, NewField(name, value))
		} else {
			t.fields[index].Set(value)
		}
	}
}

// String returns the TorRC structure as a string.
func (t *TorRC) String() string {
	var str string
	for _, f := range t.fields {
		str += f.name + " " + f.GetString() + "\n"
	}
	return str
}

// Bytes returns the TorRC structure as a byte slice.
func (t *TorRC) Bytes() []byte {
	return []byte(t.String())
}

// Load takes the path to a torrc file and loads it into the TorRC structure.
func (t *TorRC) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t, err = ParseTorRC(data)
	return err
}

// Read takes the path to a TorRC file and loads it into the TorRC structure.
func (t *TorRC) Read(filename string) error {
	return t.Load(filename)
}

// Write writes the TorRC structure to a file specified by path
func (t *TorRC) Write(path string) error {
	return ioutil.WriteFile(path, t.Bytes(), 0644)
}

// ParseTorRC takes a byte slice and returns a TorRC structure.
func ParseTorRC(path []byte) (*TorRC, error) {
	var torrc TorRC
	lines := strings.Split(string(path), "\n")
	for _, line := range lines {
		// trim away all leading spaces
		line = strings.TrimLeft(line, " ")
		// trim away spaces between a # and the rest of the line
		if strings.HasPrefix(line, "#") {
			templine := strings.TrimLeft(line, "#")
			templine = strings.TrimLeft(templine, " ")
			line = "#" + templine
		}
		fields := strings.Split(line, " ")
		if len(fields) == 0 {
			continue
		}
		torrc.fields = append(torrc.fields, NewField(fields[0], fields[1:]))
	}
	return &torrc, nil
}

// ReadTorRC takes the path to a torrc file(or torrc.d file) and returns a TorRC
// structure.
func ReadTorRC(path string) (*TorRC, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseTorRC(bytes)
}
