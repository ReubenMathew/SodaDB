package model

import (
	"fmt"
)

// Db is a Database key-value structure
type DB struct {
	keys   []interface{}               // key list
	values map[interface{}]interface{} // key-value map
}

func NewDB() *DB {

	db := &DB{
		keys:   make([]interface{}, 0),
		values: make(map[interface{}]interface{}, 0),
	}

	return db
}

// Set : sets a value indexed by a key
func (db *DB) Set(key, value interface{}) {

	// if _, ok := db.values[key]; ok {
	// 	err := errors.New("Key already found")
	// 	return err
	// }
	db.values[key] = value
	db.keys = append(db.keys, key)

}

// Get : gets a value from datastore, indexed by key
func (db *DB) Get(key interface{}) interface{} {

	value := db.values[key]

	return value
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

func (db *DB) dbPrint() {

	fmt.Println(db.keys)
	fmt.Println(db.values)

}
