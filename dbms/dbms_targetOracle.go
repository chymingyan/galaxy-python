package dbms

import (
	"database/sql"

	_ "github.com/mattn/go-oci8"
)

//返回目标数据库连接
func TargetSqlDb(connectionStr string) (db sql.DB) {

}
