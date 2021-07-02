package main

import "fmt"

func main() {
	fmt.Println(recursive(7))
	fmt.Println(cache)
	fmt.Println(iteration(7))
	fmt.Println(cache)
}

var cache []int = []int{0, 1}

func recursive(n int) int {
	if n < len(cache) {
		return cache[n]
	}
	cache = append(cache, recursive(n-2)+recursive(n-1))
	return cache[n]
}

func iteration(num int) int {
	n := len(cache) - 1
	for k := n; k < num; k++ {
		cache = append(cache, cache[k]+cache[k-1])
	}
	return cache[num]
}
