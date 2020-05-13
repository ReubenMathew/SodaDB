package server

import (
	"fmt"
	"strings"
)

// type keyAndValue struct {
// 	key   string
// 	value struct{}
// }

func eval(raw string) {

	splitStr := strings.Split(raw, " ")

	fmt.Println(splitStr, len(splitStr))

}
