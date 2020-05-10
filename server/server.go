package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

/*
CheckError : function to handle any errors and terminate server
*/
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":8081")
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		rawMessage := strings.TrimSpace(string(buf[0:n]))

		eval(rawMessage)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
