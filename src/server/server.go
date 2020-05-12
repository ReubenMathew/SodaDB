package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// SodaServer : Server interface
type SodaServer struct {
	addr *net.UDPAddr
	conn *net.UDPConn
}

// New : creates a new server object
func New() *SodaServer {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":8081")
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

// Listen : server starts to listen on assigned port
func (s *SodaServer) Listen() {

	buf := make([]byte, 1024)

	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		rawMessage := strings.TrimSpace(string(buf[0:n]))

		eval(rawMessage)

		if err != nil {
			fmt.Println("Error: ", err)
		}
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
