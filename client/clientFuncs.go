package main

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

func exit() {
	fmt.Print("Soda Terminated\n")
}
