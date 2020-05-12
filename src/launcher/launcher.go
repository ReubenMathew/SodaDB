package launcher

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ReubenMathew/sodadb/launcher/services"
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

	return answers.Service
}

// Launch : Launches a SodaDB service from an interactive CLI selector
func Launch() {
	selected := ServiceSelector()

	switch selected {
	case "client":
		fmt.Printf("\nYou chose to launch SodaDB %s.\n", selected)
		services.ClientLaunch()
	case "server":
		fmt.Printf("\nYou chose to launch SodaDB %s.\n", selected)
		services.ServerLaunch()
	case "exit":
		fmt.Printf("\nYou chose to exit SodaDB\n")
	default:
		fmt.Printf("\nYou chose to exit SodaDB\n")
	}
	fmt.Println()
}
