package main

import (
	"bufio"
	"os"
)

func main() {

	defer exitMsg() // Exit Message

	buffReader := bufio.NewReader(os.Stdin)

	for {
		replHead()

		line := getLine(buffReader)

		eval(line)

	}

}
