package main

import "fmt"

func main() {
	fmt.Println(recursive(43))
	// fmt.Println(iteration(43))
}

func recursive(n int) int {
	if n < 2 {
		return n
	}
	return recursive(n-2) + recursive(n-1)
}

func iteration(num int) int {
	i, j := 0, 1
	for k := 1; k < num; k++ {
		i, j = j, i+j
	}
	return j
}
