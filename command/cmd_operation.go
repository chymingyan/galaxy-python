package command

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
)

var dbRulesFileName = "conf/DatabaseRules.xml"
var aixRulesFileName = "conf/AixRules.xml"
var linuxRulesFileName = "conf/LinuxRules"
var DatabaseRules OracleCommands
var AixRules AixCommands
var LinuxRules LinuxCommands

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//加载OracleCommands命令
func LoadDataBaseRules(filename string) bool {
	var isOk = true
	if checkFileIsExist(filename) {
		//存在文件读取到全局变量中
		file, err := os.Open(filename)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			isOk = false
		}
		DatabaseRules = OracleCommands{}
		err = xml.Unmarshal(data, &v)
		if err != nil {
			isOk = false
		}
	} else {
		//不存在需要创建文件
		file, err := os.Create(filename)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		DatabaseRules = OracleCommands{}
	}

	return isOk
}

//加载Aix规则命令
func LoadAixRules(filename string) boo {
	var isOk = true
	if checkFileIsExist(filename) {
		//存在文件读取到全局变量中
		file, err := os.Open(filename)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			isOk = false
		}
		AixRules = AixCommands{}
		err = xml.Unmarshal(data, &v)
		if err != nil {
			isOk = false
		}
	} else {
		//不存在需要创建文件
		file, err := os.Create(filename)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		AixRules = AixCommands{}
	}

	return isOk
}

//加载Linux规则命令
func LoadLinuxRules(filename string) bool {
	var isOk = true
	if checkFileIsExist(filename) {
		//存在文件读取到全局变量中
		file, err := os.Open(filename)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			isOk = false
		}
		LinuxRules = LinuxCommands{}
		err = xml.Unmarshal(data, &v)
		if err != nil {
			isOk = false
		}
	} else {
		//不存在需要创建文件
		file, err := os.Create(filename)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		LinuxRules = LinuxCommands{}
	}

	return isOk
}

//操作xml文件
func OperationXmlFile(filename, filetype string) bool {
	var output []byte
	if filetype == "oracle" {
		//转换要保存的数据
		output, err := xml.MarshalIndent(DatabaseRules, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	} else if filetype == "aix" {
		//转换要保存的数据
		output, err := xml.MarshalIndent(AixRules, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	} else if filetype == "linux" {
		//转换要保存的数据
		output, err := xml.MarshalIndent(LinuxRules, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
	//写入文件(字节数组)
	err := ioutil.WriteFile(filename, output, 0666)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
