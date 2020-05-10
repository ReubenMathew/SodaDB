package main

import (
	"fmt"
	"strings"
)

func eval(raw string) {

	splitStr := strings.Split(raw, " ")

	fmt.Println(splitStr, len(splitStr))

}
