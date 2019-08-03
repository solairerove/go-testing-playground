package main

import "testing"

// go test -v
func TestMySum(t *testing.T) {
	sum := mySum(1, 2, 3, 4, 5)
	if sum != 15 {
		t.Error("Expected", 15, "Got", sum)
	}
}
