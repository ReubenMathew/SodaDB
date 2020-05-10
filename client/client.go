package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	welcome()

	ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8081")

	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")

	Conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)

	defer fmt.Println("Soda client closing...")
	defer Conn.Close()

	for {

		replHead()

		sendMsg, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		msg := strings.TrimSpace(sendMsg)
		// sendMsg

		if exitCase(msg) {
			break // if exit , break out of main loop and end program
		}

		buf := []byte(msg)

		_, err := Conn.Write(buf) // returns string length of write and potential write errors

		if err != nil {
			fmt.Println(msg, err)
		}

		time.Sleep(time.Millisecond * 100) // 200 millisecond buffer

	}

}
