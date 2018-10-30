package command

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)


const dbRulesFileName = "conf/DatabaseRules.xml"
const aixRulesFileName = "conf/AixRules.xml"
const linuxRulesFileName = "conf/LinuxRules"

var DatabaseRules OracleCommands
var AixHostRules AixHostCommands

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//加载OracleCommands命令
func LoadDataBaseRules() bool {
	fmt.Println("Database file Name: " + dbRulesFileName)
	var isOk = true
	if checkFileIsExist(dbRulesFileName) {
		//存在文件读取到全局变量中
		file, err := os.Open(dbRulesFileName)
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
		err = xml.Unmarshal(data, &DatabaseRules)
		if err != nil {
			isOk = false
		}
		fmt.Printf("DataBaseRules %s", DatabaseRules.Encoding)

	} else {
		//不存在需要创建文件
		file, err := os.Create(dbRulesFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		DatabaseRules.Version = "1.0"
		DatabaseRules.Encoding = "utf-8"
		output, err := xml.MarshalIndent(DatabaseRules, "  ", "    ")
		err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}

	}

	return isOk
}

//加载Aix规则命令
func LoadAixRules() bool {
	var isOk = true
	if checkFileIsExist(aixRulesFileName) {
		//存在文件读取到全局变量中
		file, err := os.Open(aixRulesFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			isOk = false
		}
		err = xml.Unmarshal(data, &AixHostRules)
		if err != nil {
			isOk = false
		}
	} else {
		//不存在需要创建文件
		file, err := os.Create(aixRulesFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		AixHostRules.Version = "1.0"
		AixHostRules.Encoding = "utf-8"
		output, err := xml.MarshalIndent(AixHostRules, "  ", "    ")
		err = ioutil.WriteFile(aixRulesFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}
	}

	return isOk
}

//加载Linux规则命令
func LoadLinuxRules() bool {
	var isOk = true
	if checkFileIsExist(LinuxHostRules) {
		//存在文件读取到全局变量中
		file, err := os.Open(LinuxHostRules)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			isOk = false
		}
		err = xml.Unmarshal(data, LinuxHostRules)
		if err != nil {
			isOk = false
		}
	} else {
		//不存在需要创建文件
		file, err := os.Create(linuxRulesFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		LinuxHostRules.Version = "1.0"
		LinuxHostRules.Encoding = "utf-8"
		output, err := xml.MarshalIndent(LinuxHostRules, "  ", "    ")
		err = ioutil.WriteFile(linuxRulesFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}
	}

	return isOk
}

//操作xml文件
func OperationXmlFile(filename, cmd string) bool {
	var output []byte
	var err error
	var isOk = true
	if filename == dbRulesFileName {

		//转换要保存的数据
		output, err = xml.MarshalIndent(DatabaseRules, "  ", "    ")
		if err != nil {
			isOk = false
		}
		err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
	//	 else if filetype == "aix" {
	//		//转换要保存的数据
	//		output, err := xml.MarshalIndent(AixRules, "  ", "    ")
	//		if err != nil {
	//			fmt.Printf("error: %v\n", err)
	//		}
	//	} else if filetype == "linux" {
	//		//转换要保存的数据
	//		output, err := xml.MarshalIndent(LinuxRules, "  ", "    ")
	//		if err != nil {
	//			fmt.Printf("error: %v\n", err)
	//		}
	//	}
	//写入文件(字节数组)
	err = ioutil.WriteFile(filename, output, 0666)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return true
}
