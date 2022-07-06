package torrc

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type Field struct {
	name  string
	value []string
}

func (f *Field) Get() []string {
	return f.value
}

func (f *Field) GetString() string {
	return strings.Join(f.value, " ")
}

func (f *Field) GetInt() (int, error) {
	if len(f.value) == 0 {
		return 0, errors.New("Field is empty")
	}
	if len(f.value) > 1 {
		val, err := strconv.Atoi(f.value[0])
		if err != nil {
			return 0, err
		}
		return val, errors.New("Field has multiple values")
	}
	return strconv.Atoi(f.value[0])
}

func (f *Field) GetInts() ([]int, error) {
	if len(f.value) == 0 {
		return nil, errors.New("Field is empty")
	}
	var ints []int
	for _, v := range f.value {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func (f *Field) Set(value ...string) {
	f.value = value
}

func (f *Field) Compare(name string) bool {
	if f.name == name {
		return true
	}
	if strings.ReplaceAll(f.name, "#", "") == name {
		return true
	}
	return false
}

func NewField(name string, value ...string) *Field {
	return &Field{name: name, value: value}
}

type TorRC struct {
	fields []*Field
}

func (t *TorRC) GetField(name string) *Field {
	for _, f := range t.fields {
		if f.Compare(name) {
			return f
		}
	}
	return nil
}

func (t *TorRC) GetFields() []*Field {
	return t.fields
}

func (t *TorRC) GetFieldsByName(name string) []*Field {
	var fields []*Field
	for _, f := range t.fields {
		if f.Compare(name) {
			fields = append(fields, f)
		}
	}
	return fields
}

func (t *TorRC) String() string {
	var str string
	for _, f := range t.fields {
		str += f.name + " " + f.GetString() + "\n"
	}
	return str
}

func (t *TorRC) Bytes() []byte {
	return []byte(t.String())
}

func (t *TorRC) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t, err = ParseTorRC(data)
	return err
}

func (t *TorRC) Read(filename string) error {
	return t.Load(filename)
}

func (t *TorRC) Write(path string) error {
	return ioutil.WriteFile(path, t.Bytes(), 0644)
}

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
		torrc.fields = append(torrc.fields, NewField(fields[0], fields[1:]...))
	}
	return &torrc, nil

}

func ReadTorRC(path string) (*TorRC, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseTorRC(bytes)
}
