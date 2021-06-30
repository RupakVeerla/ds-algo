//Google Question
//Given an array = [2,5,1,2,3,5,1,2,4]:
//It should return 2

//Given an array = [2,1,1,2,3,5,1,2,4]:
//It should return 1

//Given an array = [2,3,4,5]:
//It should return undefined

//Bonus... What if we had this:
// [2,5,5,2,3,5,1,2,4]
// return 5 because the pairs are before 2,2

package main

import "fmt"

func firstRecChar(arr []int) (n int) {
	m := map[int]bool{}
	for _, v := range arr {
		if m[v] {
			return v
		}
		m[v] = true
	}
	return 0
}

func main() {
	array := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	arra := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	arr := []int{2, 3, 4, 5}
	ar := []int{2, 5, 5, 2, 3, 5, 1, 2, 4}
	fmt.Println(firstRecChar(array), firstRecChar(arra), firstRecChar(arr), firstRecChar(ar))
}
