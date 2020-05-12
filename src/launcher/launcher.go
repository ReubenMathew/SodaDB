package launcher

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ReubenMathew/sodadb/tcp/server"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "service",
		Prompt: &survey.Select{
			Message: "Choose a SodaDB service to launch:",
			Options: []string{"client", "server", "exit"},
			Default: "client",
		},
	},
}

// ServiceSelector : CLI selector for SodaDB services
func ServiceSelector() string {
	answers := struct {
		Service string `survey:"service"`
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if answers.Service == "exit" {
		fmt.Printf("\nYou chose to exit SodaDB\n")
	} else {
		fmt.Printf("\nYou chose to launch SodaDB %s.\n", answers.Service)
	}
	return answers.Service
}

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
