// package device contains useful types to be used by drawer
package device

// ID is the device name
type ID string

// EmptyString is the string which also contains valid flag
type EmptyString struct {
	s string
	v bool
}

// NewEmptyString creates EmptyString from string and boolean flag
func NewEmptyString(s string, v bool) EmptyString {
	return EmptyString{s, v}
}

// IsValid checks if e is a valid string
func (e EmptyString) IsValid() bool {
	return e.v
}

// Get returns actual string if e is valid, empty string otherwise
func (e EmptyString) Get() string {
	if !e.v {
		return ""
	}
	return e.s
}

// Record represents one line of input
type Record struct {
	Port        string
	Description EmptyString
	VLAN        EmptyString
}

// Data represents parsed device input
type Data struct {
	Id      ID
	Records []Record
}
