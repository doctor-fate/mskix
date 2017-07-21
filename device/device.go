package device

type ID string

type Data struct {
	Id      ID
	Entries []struct {
		Port        string
		Description string
		VLAN        string
	}
}
