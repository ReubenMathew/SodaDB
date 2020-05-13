package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// SodaServer : Server struct
type SodaServer struct {
	addr *net.UDPAddr
	conn *net.UDPConn
}

// NewServer : creates a new server object
func NewServer(port int) *SodaServer {

	portStr := ":" + strconv.Itoa(port)

	ServerAddr, err := net.ResolveUDPAddr("udp", portStr)
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	server := &SodaServer{
		addr: ServerAddr,
		conn: ServerConn,
	}

	return server
}

// Close : closes server connection
func (s *SodaServer) Close() {
	fmt.Println("Closing Server Connection...")
	s.conn.Close()
}

var wg sync.WaitGroup

// Listen : server starts to concurrently listen on assigned port
func (s *SodaServer) Listen() {
	wg.Add(1)
	go s.listenFunc()
	wg.Wait()
}

func (s *SodaServer) listenFunc() {
	defer wg.Done()
	for {
		s.read()
		time.Sleep(time.Second)
	}
}

func (s *SodaServer) read() {
	buf := make([]byte, 1024)

	n, addr, err := s.conn.ReadFromUDP(buf)
	fmt.Println("Received ", string(buf[0:n]), " from ", addr)

	rawMessage := strings.TrimSpace(string(buf[0:n]))

	eval(rawMessage)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

/*
CheckError : function to handle any errors and terminate server
*/
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
