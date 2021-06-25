package main

import "fmt"

func reverse(str string) string {
	bstr := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		bstr = append(bstr, str[i])
	}
	return string(bstr)
}

func reverse1(str string) string {
	var r []rune
	for _, v := range str {
		r = append([]rune{v}, r...)
	}
	return string(r)
}

func reverse2(str string) (r string) {
	for _, v := range str {
		r = string(v) + r
	}
	return
}

func reverse3(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverse("Hello World!"))
	fmt.Println(reverse1("Hello World!"))
	fmt.Println(reverse2("Hello World!"))
	fmt.Println(reverse3("Hello World!"))
}
