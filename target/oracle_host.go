package target

type OracleHost struct {
	HostId         string
	HostIp         string
	HostProtocol   string
	HostPort       string
	HostType       string
	HostOsVer      string
	DbVer          string
	RootPwd        string
	OracleUsername string
	OraclePassword string
	OracleHome     string
	GridUsername   string
	GridPassword   string
	GridHome       string
}

type OracleHosts struct {
	XMLName  xml.Name     `xml:"hosts"`
	Version  string       `xml:"version,attr"`
	Encoding string       `xml:"encoding,attr"`
	Host     []OracleHost `xml:"oraclehost"`
}
