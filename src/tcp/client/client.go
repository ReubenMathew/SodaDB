package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// SodaClient : Client struct
type SodaClient struct {
	server *net.UDPAddr
	local  *net.UDPAddr
	conn   *net.UDPConn
}

// NewClient : creates a new UDP client
func NewClient() *SodaClient {

	ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8081")

	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")

	client := &SodaClient{
		server: ServerAddr,
		local:  LocalAddr,
	}

	Conn, _ := net.DialUDP("udp", client.local, client.server)

	client.conn = Conn

	return client
}

/*
Close : Closes UDP client
*/
func (c *SodaClient) Close() {
	c.conn.Close()
}

/*
Write : Writes bytes to a UDP server, remember to close connection!
*/
func (c *SodaClient) Write(sendMsg string) {
	msg := strings.TrimSpace(sendMsg)

	buf := []byte(msg)

	_, err := c.conn.Write(buf) // returns string length of write and potential write errors

	if err != nil {
		fmt.Println(msg, err)
	}
}

var wg sync.WaitGroup

/*
Run : Launches UDP client to listen to port 8081
*/
func (c *SodaClient) Run() {

	welcome()

	wg.Add(1)
	go c.repl()
	wg.Wait()

}

func (c *SodaClient) repl() {

	defer c.Close()
	defer wg.Done()

	for {

		replHead()

		sendMsg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		sendMsg = strings.TrimSpace(sendMsg)

		if exitCase(sendMsg) {
			fmt.Println("Exiting SodaDB Client...")
			c.Close()
			time.Sleep(time.Second)
			break // if exit , break out of main loop and end program
		}

		c.Write(sendMsg)

		time.Sleep(time.Millisecond * 100) // 200 millisecond buffer

	}

}
