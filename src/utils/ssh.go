package utils;

import (
    "fmt"
)

struct SshClient {
	ip string
}

func SshInit(ip) {
	fmt.Println("ssh init");
	instance := SshClient{ip}
	return *instance
}

func IpParser(ipString) {
	
}

func Connect(sshClient){
	
}

func Disconnect(sshClient){
	
}