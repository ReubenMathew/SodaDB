package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ReubenMathew/sodadb/model"
)

// SetCmd structure type
type SetCmd struct {
	db    *model.DB
	list  bool
	key   interface{}
	value interface{}
}

// GetCmd structure type
type GetCmd struct {
	db  *model.DB
	key interface{}
}

func eval(raw string, dB *model.DB) (interface{}, error) {

	splitStr := strings.Split(raw, " ")

	fmt.Println(splitStr, len(splitStr))

	isList := false
	cmdRaw := splitStr[0]

	switch cmdRaw {
	case "get":
		cmd := &GetCmd{
			db:  dB,
			key: splitStr[1],
		}

		value := cmd.db.Get(cmd.key)
		return value, nil

	case "set":
		if splitStr[1] == "list" && len(splitStr) == 4 {
			isList = true
		}

		cmd := &SetCmd{
			db:   dB,
			list: isList,
		}

		if cmd.list == true {
			cmd.key = splitStr[2]
			cmd.value = splitStr[3]
		} else {
			cmd.key = splitStr[1]
			cmd.value = splitStr[2]
		}

		cmd.db.Set(cmd.key, cmd.value)
		return nil, nil

	default:
		err := errors.New("Invalid command: make sure you use 'get' or 'set' with valid arguments")
		return nil, err

	}

	// return nil, nil

}
