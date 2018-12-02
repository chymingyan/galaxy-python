package dgssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

//SSH远程连接
func SshConnect(user, password, hostip string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	log.Printf("%s ssh connected.", addr)
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

//SFTP远程连接
func SftpConnect(user, password, hostip string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {

		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

//执行shell
//@param shell shell脚本命令
func RunShell(shell string, session *ssh.Session) (string, error) {

	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	shellout := string(buf)
	return shellout, err
}

//执行带交互的命令
func RunTerminal(shell string, session *ssh.Session) error {

	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	// excute command
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		log.Fatal(err)
	}

	session.Run(shell)
	return nil
}

//向远程服务器发送文件
func PutFileForRemote(host, user, pwd, oraclehome, localFilePath, remotefilePath string, port int) {

	//连接远程服务器
	sftp, err := SftpConnect(host, user, pwd, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()

	//打开本地文件
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)
	//	remotepath := path.Join(remotefilePath, remoteFileName)
	//远程写入文件
	dstFile, err := sftp.Create(path.Join(remotefilePath, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()
	srcFileByte, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	srcfilelen := len(srcFileByte)

	buf := make([]byte, srcfilelen)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	log.Println("upload file to remote server finished!")
}

//在远程服务器得到文件
func GetFileForRemote(host, user, pwd, oraclehome, localfilePath, remoteFilePath string, port int) {
	//连接远程sftp服务器
	sftp, err := SftpConnect(host, user, pwd, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()
	//打开远程目录的文件
	srcFile, err := sftp.Open(remoteFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var localFileName = path.Base(remoteFilePath)
	//生成本地文件
	dstFile, err := os.Create(path.Join(localfilePath, localFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Fatal(err)
	}
	log.Println("down load file from remote server finished!")
}
