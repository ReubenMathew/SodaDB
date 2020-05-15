package main

import "fmt"

// Db is a Database key-value structure
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

// Set : sets a value indexed by a key
func (db *Db) Set(key, value interface{}) error {

	return nil
}

// ValType : Returns the datatype of a variable
func ValType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float")
	case byte:
		fmt.Println("Byte")
	}
}

func main() {
	a := new()
	fmt.Println(a.keys)
	fmt.Println(a.values)

	b := make([]int, 2)

	c := "hello"
	a.values[c] = b
	a.keys = append(a.keys, c)

	d := 89
	a.values[d] = b
	a.keys = append(a.keys, d)

	fmt.Println(a.keys)
	fmt.Println(a.values)
}
