package target

type TargetGroup struct {
}

type TargetGroups struct {
	XMLName  xml.Name      `xml:"groups"`
	Version  string        `xml:"version,attr"`
	Encoding string        `xml:"encoding,attr"`
	Group    []TargetGroup `xml:"targetgroup"`
}
