package device

type ID string

type EmptyString struct {
	s string
	v bool
}

func NewEmptyString(s string, v bool) EmptyString {
	return EmptyString{s, v}
}

func (e EmptyString) IsValid() bool {
	return e.v
}

func (e EmptyString) Get() string {
	if !e.v {
		return ""
	}
	return e.s
}

type Record struct {
	Port        string
	Description EmptyString
	VLAN        EmptyString
}

type Data struct {
	Id      ID
	Records []Record
}
