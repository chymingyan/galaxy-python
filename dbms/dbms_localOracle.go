package dbms

import (
	"database/sql"
	"log"

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
func SaveDbInspData(sqlStr, connectionString string) bool {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	result, err := db.Exec(sqlStr)
	if err != nil {
		log.Fatalln(err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	if lastId > 0 {
		return true
	}
	return false
}

//保存主机巡检结果
func SaveHostInspData(sqlStr, connectionString string) bool {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	result, err := db.Exec(sqlStr)
	if err != nil {
		log.Fatalln(err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	if lastId > 0 {
		return true
	}
	return false
}
