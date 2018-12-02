package dbms

import (
	"database/sql"
	"ipp/ssh"
	"log"
	"strconv"

	_ "github.com/mattn/go-oci8"
)

//(HOST_NAME)主机名称
func TargetHostName(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的主机名称
	return ""
}

//(SYSTEM_LOG)系统日志
func TargetHostSystemLog(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的系统日志
	return ""
}

//(SCHEDULED_TASK_LOG)计划任务日志
func TargetHostSystemLog(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的计划任务日志
	return ""
}

//(CPU_USAGE_RATE)cup使用率
func TargetHostCpuUsageRate(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的cup使用率
	return ""
}

//(MEMORY_USAGE_RATE)内存使用率
func TargetHostMemoryUsageRate(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的内存使用率
	return ""
}

//(FILE_SYSTEM)文件系统
func TargetHostFileSystem(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的文件系统
	return ""
}

//(SWAP_USAGE)SWAP使用率
func TargetHostSwapUsage(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的SWAP使用率
	return ""
}

//(SWAP_USAGE)SWAP使用率
func TargetHostSwapUsage(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的SWAP使用率
	return ""
}

//(ROOTVG)根卷组镜像
func TargetHostRootVg(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的根卷组镜像
	return ""
}

//(MULTISOFTPATH_STATE)多路径软件运行状态 有返回结果即为正确
func TargetHostMultiSoftPathState(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的多路径软件运行状态
	return ""
}

//(MULTIPATH_STATE)多路径运行状态
func TargetHostMultiPathState(shell, username, password, hostip string, port int) string {
	sshSession, err := ssh.SshConnect(username, password, hostip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sshSession.Close()
	output, err := sshSession.RunShell(shell, sshSession)
	if err != nil {
		log.Fatal(err)
	}
	//解析output结果中的多路径运行状态
	return ""
}
