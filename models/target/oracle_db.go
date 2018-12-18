package target

const oracleDbsFileName = "conf/OracleDatabases.xml"

var OracleDbList OracleDbs

type OracleDb struct {
	DbId       string
	DbIp       string
	DbPort     string
	DbSid      string
	DbUserName string
	DbPassword string
	GroupId    string
	Alias      string
	DbVer      string
	ExpPath    string
	Remark     string
}

type OracleDbs struct {
	XMLName  xml.Name   `xml:"databases"`
	Version  string     `xml:"version,attr"`
	Encoding string     `xml:"encoding,attr"`
	Database []TargetDb `xml:"targetdb"`
}

//加载Oracle数据库
func LoadOracleDbs() bool {
	var isOk = true
	if _, err := os.Stat(oracleDbsFileName); os.IsNotExist(err) {
		isOk = false
	}
	if isOk {
		//存在文件读取到全局变量中
		file, err := os.Open(oracleDbsFileName)
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
		err = xml.Unmarshal(data, &OracleDbList)
		if err != nil {
			isOk = false
		}

	} else {
		//不存在需要创建文件
		file, err := os.Create(oracleDbsFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		OracleDbList.Version = "1.0"
		OracleDbList.Encoding = "utf-8"
		output, err := xml.MarshalIndent(OracleDbList, "  ", "    ")
		err = ioutil.WriteFile(oracleDbsFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}

	}
	return isOk
}

//新增一条数据库信息
func (this *OracleDbs) AddDb(db OracleDb) bool {
	var isOk = true
	OracleDbList.Database = append(OracleDbList.Database, cmd)
	output, err := xml.MarshalIndent(OracleDbList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(oracleDbsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//删除一条数据库信息
func (this *OracleDbs) DelDb(dbId string) bool {
	var isOk = true
	tempDbs := []OracleDb{}
	dbCount := len(OracleDbList.Database)
	for i := 0; i < dbCount; i++ {
		if OracleDbList.Database[i].DbId != dbId {
			tempDbs = append(tempDbs, OracleDbList.Database[i])
		}
	}
	OracleDbList.Database = tempDbs
	output, err := xml.MarshalIndent(OracleDbList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(oracleDbsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//修改一条数据库信息
func (this *OracleDbs) ModifyDb(db OracleDb) bool {
	var isOk = true
	tempDbs := []OracleCommand{}
	dbCount := len(OracleDbList.Database)
	for i := 0; i < dbCount; i++ {
		if OracleDbList.Database[i].DbId != db.DbId {
			tempDbs = append(tempDbs, OracleDbList.Database[i])
		}
	}
	tempDbs = append(tempDbs, db)
	OracleDbList.Database = tempDbs
	output, err := xml.MarshalIndent(OracleDbList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(oracleDbsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//查询一条数据库信息
func (this *OracleDbs) QueryDb(dbId string) (db OracleDb) {
	db = OracleDb{}
	dbCount := len(OracleDbList.Database)
	for i := 0; i < dbCount; i++ {
		if OracleDbList.Database[i].DbId == dbId {
			cmd = OracleDbList.Database[i]
			break
		}
	}
	return
}

//查询全部的数据库信息
func (this *OracleDbs) QueryDbs() (dbs []OracleDb) {
	dbs = OracleDbList.Database
	return
}

