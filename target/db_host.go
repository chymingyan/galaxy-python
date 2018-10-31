package target

const dbHostsFileName = "conf/DbHosts.xml"

var DbHostList DbHosts

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

//加载数据库和主机关联信息
func LoadDbHost() {
	var isOk = true
	if _, err := os.Stat(dbHostsFileName); os.IsNotExist(err) {
		isOk = false
	}
	if isOk {
		//存在文件读取到全局变量中
		file, err := os.Open(dbHostsFileName)
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
		err = xml.Unmarshal(data, &DbHostList)
		if err != nil {
			isOk = false
		}
		fmt.Printf("DataBaseRules %s", DbHostList.Encoding)

	} else {
		//不存在需要创建文件
		file, err := os.Create(dbHostsFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		DbHostList.Version = "1.0"
		DbHostList.Encoding = "utf-8"
		output, err := xml.MarshalIndent(DbHostList, "  ", "    ")
		err = ioutil.WriteFile(dbHostsFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}

	}
	return isOk
}

//添加一条数据库和主机关联信息
func (this *DbHosts) AddDbHost(dbhost DbHost) bool {
	var isOk = true
	DbHostList.Relation = append(DbHostList.Relation, dbhost)
	output, err := xml.MarshalIndent(DbHostList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbHostsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//删除一条数据库和主机的关联信息
func (this *DbHosts) DelDbHost(dbhostId string) bool {
	var isOk = true
	tempDbHosts := []DbHost{}
	relationCount := len(DbHostList.Relation)
	for i := 0; i < relationCount; i++ {
		if DbHostList.Relation[i].DbHostId != dbhostId {
			tempDbHosts = append(tempDbHosts, DbHostList.Relation[i])
		}
	}
	DbHostList.Relation = tempDbHosts
	output, err := xml.MarshalIndent(DbHostList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbHostsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//修改一条数据库和主机的关联信息
func (this *DbHosts) ModifyDbHost(dbhost DbHost) bool {
	var isOk = true
	tempDbHosts := []DbHost{}
	relationCount := len(DbHostList.Relation)
	for i := 0; i < relationCount; i++ {
		if DbHostList.Relation[i].DbHostId != dbhost.DbHostId {
			tempDbHosts = append(tempDbHosts, DbHostList.Relation[i])
		}
	}
	tempDbHosts = append(tempDbHosts, cmd)
	DbHostList.Relation = tempDbHosts
	output, err := xml.MarshalIndent(DbHostList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbHostsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//查询一条数据库和主机关联信息
func (this *DbHosts) QueryDbHost(dbhostId string) (dbhost DbHost) {
	dbhost = DbHost{}
	relationCount := len(DbHostList.Relation)
	for i := 0; i < relationCount; i++ {
		if DbHostList.Relation[i].DbHostId == dbhostId {
			dbhost = DbHostList.Relation[i]
			break
		}
	}
	return
}

//查询全部数据库和主机关联信息
func (this *DbHosts) QueryDbHosts() (dbhosts []DbHost) {
	dbhosts = DbHostList.Relation
	return
}
