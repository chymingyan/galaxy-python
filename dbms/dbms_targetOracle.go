package dbms

import (
	"database/sql"
	"ipp/ssh"
	"log"
	"strconv"

	_ "github.com/mattn/go-oci8"
)

//返回目标数据库连接
func TargetDbSqlConnection(username, password, ip, servicename string, port int) string {
	connectionString := username + "/" + password + "@" + ip + ":" + strconv.Itoa(port) + "/" + servicename
	return connectionString
}

//(ARCHIVELOG_MODE)归档模式
func TargetDbArchiveLogMode(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	//	for rows.Next() {

	//		var THREAD string
	//		var GROUP string
	//		var MEMBER string
	//		var TYPE string
	//		var MB string
	//		if err := rows.Scan(&THREAD, &GROUP, &MEMBER, &TYPE, &MB); err != nil {
	//			log.Fatal(err)
	//		}
	//		log.Println("%s ,%s ,%s, %s, %s\n", THREAD, GROUP, MEMBER, TYPE, MB)

	//		group, err := strconv.Atoi(GROUP)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		mb, err := strconv.Atoi(MB)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		logfile := new(LogFile)
	//		logfile.Thread = THREAD
	//		logfile.Group = group
	//		logfile.Member = MEMBER
	//		logfile.Type = TYPE
	//		logfile.Mb = mb
	//		logFiles = append(logFiles, *logfile)

	//	}
	defer rows.Close()
	return
}

//(ARCHIVELOG_PATH)归档路径
func TargetDbArchiveLogPath(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	var archivelogpath string
	for rows.Next() {
		if err := rows.Scan(&archivelogpath); err != nil {
			log.Fatal(err)
		}
		return archivelogpath
	}
	defer rows.Close()
	return ""
}

//(ASM_DISK_INFO)ASM_DISK_INFO磁盘组信息
func TargetDbAsmDiskInfo(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	defer rows.Close()
	return ""
}

//(AUDIT_TRAIL)审计
func TargetDbAuditTrail(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	defer rows.Close()
	return ""
}

//(BROKEN_JOB)失效的任务
func TargetDbBrokenJob(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	defer rows.Close()
	return ""
}

//(DBID)数据库DBID
func TargetDbId(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	var dbid string
	for rows.Next() {
		if err := rows.Scan(&dbid); err != nil {
			log.Fatal(err)
		}
		return dbid
	}
	defer rows.Close()
	return ""
}

//(DB_NAME)数据库DBName
func TargetDbName(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	var dbname string
	for rows.Next() {
		if err := rows.Scan(&dbname); err != nil {
			log.Fatal(err)
		}
		return dbname
	}
	defer rows.Close()
	return ""
}

//(DB_PARAMETERS)数据库参数
func TargetDbParameters(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	defer rows.Close()
	return ""
}

//(INSTANCE_NAME)数据库实例名称
func TargetDbInstanceName(sqlStr, connectionString string) string {
	db, err := sql.Open("oci8", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	defer rows.Close()
	return ""
}

//(ASM)查询数据库是否为ASM
func TargetDbIsAsm(shell, username, password, hostip string, port int) bool {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//验证output结果是否为asm（是返回true，否返回false）
	return true
}

//(ASM_DISK)查询目标数据库的磁盘组信息
func TargetDbAsmDiskInfo(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的磁盘组信息
	return ""
}

//(RDBMS)RDBMS版本补丁
func TargetDbRdbmsBugs(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的RDBMS版本补丁
	return ""
}

//(HOST_MODEL)主机型号
func TargetDbHostModel(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的主机型号
	return ""
}

//(ALERT_LOG)数据库告警日志
func TargetDbAlertLog(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的数据库告警日志
	return ""
}
