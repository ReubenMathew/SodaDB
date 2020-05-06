package main

import (
	"bufio"
	"fmt"
	"os"
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

func exitMsg() {
	fmt.Print("Soda Terminated\n")
}

func getLine(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	t = strings.TrimSpace(t)
	return t
}

func eval(line string) {
	if exitCase(line) {
		fmt.Print("Exiting...\n")
		os.Exit(0)
	}

	fmt.Println(line)
}
