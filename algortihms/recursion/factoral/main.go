package main

import "fmt"

func main() {
	fmt.Println(recursion(5))
	fmt.Println(iteration(5))
}

func recursion(num int) int {
	if num == 2 {
		return num
	}
	return num * recursion(num-1)
}

func iteration(num int) int {
	r := num
	for i := num - 1; i > 1; i-- {
		r *= i
	}
	return r
}
