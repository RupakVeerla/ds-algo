package main

import (
	"fmt"
	"sort"
)

func mergeSortedArrays(arr1, arr2 []int) (r []int) {
	r = append(r, append(arr1, arr2...)...)
	sort.Ints(r)
	return
}

func mergeSortedArrays2(arr1, arr2 []int) (r []int) {
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			r = append(r, arr1[i])
			i++
			if i == len(arr1) {
				r = append(r, arr2[j:]...)
			}
		} else {
			r = append(r, arr2[j])
			j++
			if j == len(arr2) {
				r = append(r, arr1[i:]...)
			}
		}
	}
	return r
}

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 10, 43}, []int{5, 8, 31, 42}))
	fmt.Println(mergeSortedArrays2([]int{0, 3, 10, 43}, []int{5, 8, 31, 42, 69}))
}
