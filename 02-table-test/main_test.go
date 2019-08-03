package main

import "testing"

// go test -v
func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{21, 21}, 42},
	}

	for _, v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Error("Expected", v.answer, "Got", sum)
		}
	}
}
