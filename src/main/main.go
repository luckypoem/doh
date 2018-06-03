package main

import (
	"net"
)

func main() {
	CheckPermissions() //Checking permissions of runtime, upgrades if necessairy
	ServerConn := Server() //create server
	defer ServerConn.Close() //close server if program exits
	go CheckDNS() //check DNS settings of device to prevent DNS leaking
	Listener(ServerConn) //start listening for infinite time
}

func Listener(ServerConn *net.UDPConn) { //listen to incomming packets
	for { //infinite loop
		buf := make([]byte, 1024) //create buffer
		n,addr,_ := ServerConn.ReadFromUDP(buf) //read incoming packets into buffer
		go ProcessRequest(n,addr,buf, ServerConn) //process request async
	}
}

func CheckDNS() { //Check dns settings of device
	/*
	//CHECK SYSTEM DNS SETUP //TODO recheck every minute?
	if runtime.GOOS == "darwin" { // macos
		//TODO // TODO //TODO NSLOOKUP WITHOUT EXTRA RESOURCES TO FIGURE OUT NS SERVERS
		output, _ := exec.Command("sh", "-c", "nslookup empty | grep Server | grep 127.0.0.1").CombinedOutput()
		if len(output) != 0 { //good
		} else { //bad
			fmt.Println("WARN To use Doh, the system dns server must be 127.0.0.1") //TODO color bash
			exec.Command("sh", "-c", "osascript -e 'tell app \"System Events\" to display dialog \"DoH can not work for you because your system DNS is not set to 127.0.0.1\"'").Run()
		}
	}*/
}

func Server() *net.UDPConn { //create server
	ServerAddr,err := net.ResolveUDPAddr("udp",":53") //create address object for server to bind to
	CheckError(err) //check error, if found, log and exit
	ServerConn, err := net.ListenUDP("udp", ServerAddr) //bind server to address
	CheckError(err) //check error, if found, log and exit
	Notification("display notification \"Successfully started!\" with title \"DoH\"") //display a notification when succesful
	return ServerConn
}
