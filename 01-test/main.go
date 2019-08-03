package main

import "fmt"

func main() {
	sum := mySum(1, 2, 3, 4, 5)
	fmt.Println(sum)
}

func mySum(x ...int) int {
	var sum int

	for _, v := range x {
		sum += v
	}

	return sum
}
