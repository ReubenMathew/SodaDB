package main

import (
	"os"

	"github.com/ReubenMathew/sodadb/launcher"
)

func main() {
	selected := launcher.ServiceSelector()

	switch selected {
	case "client":
	case "server":
		launcher.ServerLaunch()
	case "exit":
		os.Exit(0)
	}

}
