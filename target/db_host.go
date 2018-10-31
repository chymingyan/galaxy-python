package target

type DbHost struct {
	DbHostId string
	DbId     string
	HostId   string
}

type DbHosts struct {
	XMLName  xml.Name `xml:"dbhosts"`
	Version  string   `xml:"version,attr"`
	Encoding string   `xml:"encoding,attr"`
	Relation []DbHost `xml:"dbhost"`
}
