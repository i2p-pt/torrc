# [torrc](/)

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/i2p-pt/torrc)
[![Go Report Card](https://goreportcard.com/badge/github.com/i2p-pt/torrc)](https://goreportcard.com/report/github.com/i2p-pt/torrc)

Package torrc provides a simple way to read, edit, and write torrc files.

It takes a torrc file and parses it into a TorRC structure, which can be
manipulated and then written back to a file.

Example(set up an I2P Bridge Client):

```go
torrc, err := ReadTorRC("torrc.test")
if err != nil {
	panic(err)
}
torrc.SetField("UseBridges", []string{"1"})
torrc.SetField("Bridge", []string{"i2p", "i2pbase32addressesarefiftytwocharacterslongenoughok.b32.i2p"})
torrc.SetField("ClientTransportPlugin", []string{"i2p", "exec", "i2p-client"})
if err = torrc.Write("torrc.test.i2p-client"); err != nil {
	panic(err)
}
```

Example(set up an I2P Bridge Server):

```go
torrc, err := ReadTorRC("torrc.test")
if err != nil {
	panic(err)
}
torrc.SetField("BridgeRelay", []string{"1"})
torrc.SetField("ORPort", []string{"9001"})
torrc.SetField("ExtORPort", []string{"7689"})
torrc.SetField("ServerTransportPlugin", []string{"i2p", "exec", "i2p-server"})
if err = torrc.Write("torrc.test.i2p-server"); err != nil {
	panic(err)
}
```

## Types

### type [Field](/field.go#L10)

`type Field struct { ... }`

Field is a structure that holds one line of a TorRC file.

#### func [NewField](/field.go#L82)

`func NewField(name string, value []string) *Field`

NewField creates a new Field structure.

#### func (*Field) [Compare](/field.go#L71)

`func (f *Field) Compare(name string) bool`

Compare matches a name, whether it is commented out or not.

#### func (*Field) [Get](/field.go#L16)

`func (f *Field) Get() []string`

Get the value of a field in the TorRC file.

#### func (*Field) [GetInt](/field.go#L28)

`func (f *Field) GetInt() (int, error)`

GetInt returns the value of a field as an integer. If it is a multiple-value field,
it returns the first value. If it is not possible to convert the value to an integer,
it returns an error.

#### func (*Field) [GetInts](/field.go#L45)

`func (f *Field) GetInts() ([]int, error)`

GetInts returns the values of a field as integers. If it is a multiple-value
field, it returns all values. If it is not possible to convert a value to an
integer, it returns an error.

#### func (*Field) [GetString](/field.go#L21)

`func (f *Field) GetString() string`

GetString returns the value of a field as a string.

#### func (*Field) [IsComment](/field.go#L61)

`func (f *Field) IsComment() bool`

IsComment returns true if the field is a comment.

#### func (*Field) [Set](/field.go#L66)

`func (f *Field) Set(value []string)`

Set the value of a field in the TorRC file.

### type [TorRC](/torrc.go#L40)

`type TorRC struct { ... }`

TorRC is a structure that holds the contents of a torrc file.

#### func [ParseTorRC](/torrc.go#L136)

`func ParseTorRC(path []byte) (*TorRC, error)`

ParseTorRC takes a byte slice and returns a TorRC structure.

#### func [ReadTorRC](/torrc.go#L159)

`func ReadTorRC(path string) (*TorRC, error)`

ReadTorRC takes the path to a torrc file(or torrc.d file) and returns a TorRC
structure.

#### func (*TorRC) [Bytes](/torrc.go#L107)

`func (t *TorRC) Bytes() []byte`

Bytes returns the TorRC structure as a byte slice.

#### func (*TorRC) [GetField](/torrc.go#L55)

`func (t *TorRC) GetField(name string) *Field`

GetField returns the value of the field with the given name.

#### func (*TorRC) [GetFieldIndex](/torrc.go#L45)

`func (t *TorRC) GetFieldIndex(name string) int`

GetFieldIndex returns the index of the field with the given name.

#### func (*TorRC) [GetFields](/torrc.go#L64)

`func (t *TorRC) GetFields() []*Field`

GetFields returns a slice of all the fields in the TorRC

#### func (*TorRC) [GetFieldsByName](/torrc.go#L70)

`func (t *TorRC) GetFieldsByName(name string) []*Field`

GetFieldsByName returns a slice of all the fields in the TorRC matching
the given name. This includes all commented-out fields.

#### func (*TorRC) [Load](/torrc.go#L112)

`func (t *TorRC) Load(filename string) error`

Load takes the path to a torrc file and loads it into the TorRC structure.

#### func (*TorRC) [Read](/torrc.go#L126)

`func (t *TorRC) Read(filename string) error`

Read takes the path to a TorRC file and loads it into the TorRC structure.

#### func (*TorRC) [SetField](/torrc.go#L83)

`func (t *TorRC) SetField(name string, value []string)`

SetField sets the value of a field in the TorRC structure. If the field
is commented out or non-existent, it places a new entry at the end of the
torrc file.

#### func (*TorRC) [String](/torrc.go#L98)

`func (t *TorRC) String() string`

String returns the TorRC structure as a string.

#### func (*TorRC) [Write](/torrc.go#L131)

`func (t *TorRC) Write(path string) error`

Write writes the TorRC structure to a file specified by path

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
