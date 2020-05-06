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

func main() {

	defer exitMsg() // Exit Message

	buffReader := bufio.NewReader(os.Stdin)

	for {
		replHead()
		line := getLine(buffReader)

		if exitCase(line) {
			fmt.Print("Exiting...\n")
			return
		}

		fmt.Println(line)

	}

}
