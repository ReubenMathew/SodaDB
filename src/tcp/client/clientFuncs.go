package client

import (
	"fmt"
	"strings"
)

func replHead() {
	fmt.Printf("soda> ")
}

func exitCase(text string) bool {
	if strings.EqualFold("exit", text) {
		return true
	}
	return false
}

func welcome() {
	fmt.Printf("\nWelcome to SodaDB Client\n")
	fmt.Printf("Type 'h' for help or visit the docs at https://github.com/ReubenMathew/SodaDB\n\n")
}
