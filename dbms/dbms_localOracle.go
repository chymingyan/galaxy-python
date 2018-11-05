package dbms

import (
	"database/sql"

	_ "github.com/mattn/go-oci8"
)

//本地数据库连接字符串
var LocalDatabaseStr string

//登陆资料库（登陆系统）
func SingIn(ip, port, sid, username, pwd string) bool {

	//写入本地字符串
	LocalDatabaseStr = username + "/" + pwd + "@" + ip + ":" + port + "/" + sid
	return true
}

//保存数据库巡检结果
func SaveDbInspData() {

}

//保存主机巡检结果
func SaveHostInspData() {

}
