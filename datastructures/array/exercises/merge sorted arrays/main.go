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

func mergeSortedArrays1(arr1, arr2 []int) (r []int) {
	arr1val, arr2val := arr1[0], arr2[0]
	i, j := 1, 1
	for i <= len(arr1) || j <= len(arr2) {
		if arr1val < arr2val {
			r = append(r, arr1val)
			arr1val = arr1[i]
			i++
		} else {
			r = append(r, arr2val)
			arr2val = arr2[j]
			j++
		}
	}
	return
}

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 10, 43}, []int{5, 8, 31, 42}))
	fmt.Println(mergeSortedArrays1([]int{0, 3, 10, 43}, []int{5, 8, 31, 42, 69}))
}
