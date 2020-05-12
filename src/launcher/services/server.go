package services

import (
	"fmt"
	"time"

	"github.com/ReubenMathew/sodadb/tcp/server"
)

// ServerLaunch : Launches a server instance
func ServerLaunch() {
	fmt.Println("Launching Server!")
	time.Sleep(500)
	fmt.Println("--- Ctrl+C to close server ---")
	// client.Launch()

	server := server.NewServer(8081)

	defer server.Close()

	server.Listen()
}
