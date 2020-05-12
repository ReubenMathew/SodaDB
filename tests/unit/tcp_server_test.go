package main

import "testing"

func TestServerUp(t *testing.T) {
	total := 5 + 5
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

// Tests must have
// 1. t *testing.T
// 2. lowercase filename (seperate words with _)
// 3. must have t.Error or t.Fail
// 4. testing functions must start with "Test"
