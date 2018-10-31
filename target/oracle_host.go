package target

const oracleHostsFileName = "conf/OracleHosts.xml"

var OracleHostList OracleHosts

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

//加载Oracle主机信息
func LoadOracleHosts() {
	var isOk = true
	if _, err := os.Stat(oracleHostsFileName); os.IsNotExist(err) {
		isOk = false
	}
	if isOk {
		//存在文件读取到全局变量中
		file, err := os.Open(oracleHostsFileName)
		if err != nil {
			fmt.Printf("file: %s", err)
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("data %s", err)
			isOk = false
		}
		err = xml.Unmarshal(data, &OracleHostList)
		if err != nil {
			isOk = false
		}

	} else {
		//不存在需要创建文件
		file, err := os.Create(oracleHostsFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		OracleHostList.Version = "1.0"
		OracleHostList.Encoding = "utf-8"
		output, err := xml.MarshalIndent(OracleHostList, "  ", "    ")
		err = ioutil.WriteFile(oracleHostsFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}
	}
	return isOk
}

//添加oracle主机信息
func (this *OracleHosts) AddHost(host OracleHost) bool {
	var isOk = true
	OracleHostList.Host = append(OracleHostList.Host, host)
	output, err := xml.MarshalIndent(OracleHostList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(oracleHostsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//删除oracle主机信息
func (this *OracleHosts) DelHost(hostId string) bool {
	var isOk = true
	tempHosts := []OracleHost{}
	hostCount := len(OracleDbList.Host)
	for i := 0; i < hostCount; i++ {
		if OracleDbList.Host[i].HostId != hostId {
			tempHosts = append(tempHosts, OracleDbList.Host[i])
		}
	}
	OracleDbList.Host = tempHosts
	output, err := xml.MarshalIndent(OracleDbList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(oracleHostsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//修改Oracle主机信息
func (this *OracleHosts) ModifyHost(host OracleHost) bool {
	var isOk = true
	tempHosts := []OracleHost{}
	hostCount := len(OracleHostList.Host)
	for i := 0; i < hostCount; i++ {
		if OracleHostList.Host[i].HostId != host.HostId {
			tempHosts = append(tempHosts, OracleHostList.Host[i])
		}
	}
	tempHosts = append(tempHosts, host)
	OracleHostList.Host = tempHosts
	output, err := xml.MarshalIndent(OracleHostList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(oracleHostsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//查询一条Oracle主机信息
func (this *OracleHosts) QueryHost(hostId string) (host OracleHost) {
	host = OracleHost{}
	hostCount := len(OracleHosts.Host)
	for i := 0; i < hostCount; i++ {
		if OracleHosts.Host[i].HostId == hostId {
			host = OracleHosts.Host[i]
			break
		}
	}
	return
}

//查询全部的主机信息
func (this *OracleHosts) QueryHosts() (hosts []OracleHost) {
	hosts = OracleHostList.Host
	return
}
