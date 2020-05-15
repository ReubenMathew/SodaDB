package server

import (
	"fmt"
	"strings"
)

// type keyAndValue struct {
// 	key   string
// 	value struct{}
// }

func eval(raw string) (interface{}, interface{}) {

	splitStr := strings.Split(raw, " ")

	fmt.Println(splitStr, len(splitStr))

	if isSet(splitStr) {
		return splitStr[1], splitStr[2]
	}

	return nil, nil
}

func isSet(args []string) bool {

	if len(args) == 3 && args[0] == "set" {
		return true
	}
	return false

}
