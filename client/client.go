package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8081")

	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")

	Conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)

	defer Conn.Close()

	for {

		replHead()

		sendMsg, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		msg := sendMsg

		buf := []byte(msg)

		_, err := Conn.Write(buf)

		if err != nil {
			fmt.Println(msg, err)
		}

		time.Sleep(time.Millisecond * 200) // 200 millisecond buffer

	}

}
