package services

import (
	"fmt"
	"time"

	"github.com/ReubenMathew/sodadb/tcp/client"
)

// ClientLaunch : Launches a server instance
func ClientLaunch() {

	fmt.Println("Launching Client!")
	time.Sleep(500)
	fmt.Println("--- Ctrl+C or type 'exit' to close server ---")
	// client.Launch()

	client := client.NewClient()
	client.Run()

}
