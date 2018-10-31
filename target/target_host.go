package target

type TargetHost struct {
}

type TargetHosts struct {
	XMLName  xml.Name     `xml:"hosts"`
	Version  string       `xml:"version,attr"`
	Encoding string       `xml:"encoding,attr"`
	Host     []TargetHost `xml:"targethost"`
}
