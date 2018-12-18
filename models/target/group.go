package target

const groupsFileName = "Groups.xml"

var GroupList TargetGroups

type TargetGroup struct {
	GroupId   string
	GroupName string
}

type TargetGroups struct {
	XMLName  xml.Name      `xml:"groups"`
	Version  string        `xml:"version,attr"`
	Encoding string        `xml:"encoding,attr"`
	Group    []TargetGroup `xml:"targetgroup"`
}

//加载分组信息
func LoadGroups() bool {
	var isOk = true
	if _, err := os.Stat(groupsFileName); os.IsNotExist(err) {
		isOk = false
	}
	if isOk {
		//存在文件读取到全局变量中
		file, err := os.Open(groupsFileName)
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
		err = xml.Unmarshal(data, &GroupList)
		if err != nil {
			isOk = false
		}

	} else {
		//不存在需要创建文件
		file, err := os.Create(groupsFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		GroupList.Version = "1.0"
		GroupList.Encoding = "utf-8"
		output, err := xml.MarshalIndent(GroupList, "  ", "    ")
		err = ioutil.WriteFile(groupsFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}

	}
	return isOk
}

//新增一个分组
func (this *TargetGroups) AddGroup(group TargetGroup) bool {
	var isOk = true
	GroupList.Group = append(GroupList.Group, group)
	output, err := xml.MarshalIndent(GroupList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(groupsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//删除一条分组信息
func (this *TargetGroups) DelGroup(groupId string) bool {
	var isOk = true
	tempGroups := []TargetGroup{}
	goupCount := len(GroupList.Group)
	for i := 0; i < goupCount; i++ {
		if GroupList.Group[i].GroupId != groupId {
			tempGroups = append(tempGroups, GroupList.Group[i])
		}
	}
	GroupList.Group = tempGroups
	output, err := xml.MarshalIndent(GroupList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//修改一条分组信息
func (this *TargetGroups) ModifyGroup(group TargetGroup) bool {
	var isOk = true
	tempGroups := []TargetGroup{}
	groupCount := len(GroupList.Group)
	for i := 0; i < groupCount; i++ {
		if GroupList.Group[i].GroupId != group.GroupId {
			tempGroups = append(tempGroups, GroupList.Group[i])
		}
	}
	tempGroups = append(tempGroups, cmd)
	GroupList.Group = tempGroups
	output, err := xml.MarshalIndent(GroupList, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(groupsFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//查询一条分组信息
func (this *TargetGroups) QueryGroup(groupId string) (group TargetGroup) {
	group = TargetGroup{}
	groupCount := len(GroupList.Group)
	for i := 0; i < groupCount; i++ {
		if GroupList.Group[i].GroupId == groupId {
			cmd = GroupList.Group[i]
			break
		}
	}
	return
}

//查询全部分组信息
func (this *TargetGroups) QueryGroups() (groups []TargetGroup) {
	groups = GroupList.Group
	return
}
