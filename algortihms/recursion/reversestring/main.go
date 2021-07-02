package main

import "fmt"

func main() {
	fmt.Println(recursive("Rupak sai nag veerla"))
	fmt.Println(iteration("Rupak sai nag veerla"))
}

func recursive(s string) string {
	if len(s) == 1 {
		return s
	}
	return string(s[len(s)-1]) + recursive(s[:len(s)-1])
}

func iteration(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}
