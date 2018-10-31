package target

type TargetDb struct {
}

type TargetDbs struct {
	XMLName  xml.Name   `xml:"databases"`
	Version  string     `xml:"version,attr"`
	Encoding string     `xml:"encoding,attr"`
	Database []TargetDb `xml:"targetdb"`
}
