package device

type ID string

type Record struct {
	Port        string
	Description string
	VLAN        string
}

type Data struct {
	Id      ID
	Records []Record
}
