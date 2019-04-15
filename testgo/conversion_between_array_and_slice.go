package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	var a [3]int

	//fmt.Println(copy(a, s)) //오류 발생 [first argument to copy should be slice; have [3]int]
	fmt.Println(copy(a[:3], s))
	fmt.Println(a)

	a1 := [...]int{1, 2, 3}
	s1 := make([]int, 3)

	fmt.Println(copy(s1, a1[:2]))
	fmt.Println(s1)
}
