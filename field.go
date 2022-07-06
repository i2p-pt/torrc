package torrc

import (
	"errors"
	"strconv"
	"strings"
)

// Field is a structure that holds one line of a TorRC file.
type Field struct {
	name  string
	value []string
}

// Get the value of a field in the TorRC file.
func (f *Field) Get() []string {
	return f.value
}

// GetString returns the value of a field as a string.
func (f *Field) GetString() string {
	return strings.Join(f.value, " ")
}

// GetInt returns the value of a field as an integer. If it is a multiple-value field,
// it returns the first value. If it is not possible to convert the value to an integer,
// it returns an error.
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

// GetInts returns the values of a field as integers. If it is a multiple-value
// field, it returns all values. If it is not possible to convert a value to an
// integer, it returns an error.
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

// IsComment returns true if the field is a comment.
func (f *Field) IsComment() bool {
	return strings.HasPrefix(f.name, "#")
}

// Set the value of a field in the TorRC file.
func (f *Field) Set(value []string) {
	f.value = value
}

// Compare matches a name, whether it is commented out or not.
func (f *Field) Compare(name string) bool {
	if f.name == name {
		return true
	}
	if strings.TrimLeft(strings.TrimLeft(f.name, "#"), " ") == name {
		return true
	}
	return false
}

// NewField creates a new Field structure.
func NewField(name string, value []string) *Field {
	return &Field{name: name, value: value}
}
