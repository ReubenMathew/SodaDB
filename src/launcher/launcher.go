package launcher

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "service",
		Prompt: &survey.Select{
			Message: "Choose a SodaDB service to launch:",
			Options: []string{"client", "server", "tester"},
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

	fmt.Printf("\nYou chose to launch SodaDB %s.\n", answers.Service)

	return answers.Service
}

// func serverLaunch() {
// 	fmt.Println("Launching Server!")
// 	fmt.Println("--- Ctrl+C to close server ---")
// 	// client.Launch()

// 	server := server.NewServer(8081)

// 	defer server.Close()

// 	server.Listen()

// }
