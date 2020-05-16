package model

import (
	"os"
	"testing"
)

var db *DB

func setup() {
	db = NewDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestSetError(t *testing.T) {
	key := "key1"
	value := "value1"
	err := db.Set(key, value)

	if err != nil {
		t.Fail()
	}
}

func TestSetDuplicate(t *testing.T) {
	key := "key1"
	value := 123
	err := db.Set(key, value)

	if err == nil {
		t.Fail()
	}
}

func TestSetGet(t *testing.T) {
	key := "foo"
	value := "bar"
	err := db.Set(key, value)

	if err != nil {
		t.Fail()
	}

	valueRet := db.Get(key)

	if value != valueRet {
		t.Fail()
	}
}

// Tests must have
// 1. t *testing.T
// 2. lowercase filename (seperate words with _)
// 3. must have t.Error or t.Fail
// 4. testing functions must start with "Test"
