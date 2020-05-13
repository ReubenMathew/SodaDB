package main

import "fmt"

type Db struct {
	keys   []interface{}               // key
	values map[interface{}]interface{} // key value
}

func new() *Db {

	db := &Db{
		keys:   make([]interface{}, 0),
		values: make(map[interface{}]interface{}, 0),
	}

	return db
}

// func (db *Db) Set(key, value interface{}) error {

// 	return nil
// }

func valType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float")
	}
}

func main() {
	a := new()
	fmt.Println(a.keys)
	fmt.Println(a.values)

	b := make([]int, 2)

	a.values["hello"] = b
	a.keys = append(a.keys, "hello")

	fmt.Println(a.keys)
	fmt.Println(a.values)

}
