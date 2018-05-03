package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp",":53")
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	defer ServerConn.Close()
	buf := make([]byte, 1024)

	for {
		n,addr,_ := ServerConn.ReadFromUDP(buf)
		go func(n int, addr *net.UDPAddr, buf []byte) {
			record:=""
			for _, element := range buf[13:n-5] {
				if element < 31 { //first this was 8
					record+="."
				} else {
					record+=string(element)
				}
			}
			res, err := http.Get("https://1.1.1.1/dns-query?ct=application/dns-json&name="+record+"&type=A")
			CheckError(err)

			body, err := ioutil.ReadAll(res.Body)
			CheckError(err)

			fmt.Println(string(body))
			var f interface{}
			json.Unmarshal(body, &f)
			m := f.(map[string]interface{})
			fmt.Println(m["Answer"]) Tis is a type interface {}

			//var g interface{}
			//g := m["Answer"].(map[string]interface{})

			//fmt.Println(g)

			//TODO catch empty response (failed)
			//TODO carc 

		}(n,addr,buf)
	}

}
