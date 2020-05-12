package main

import (
	"fmt"

	"./server"
)

func main() {

	fmt.Println("Launching Server!")
	fmt.Println("--- Ctrl+C to close server ---")
	// client.Launch()

	server := server.New()

	defer server.Close()

	server.Listen()

}
