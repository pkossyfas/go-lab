package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}
	temp := []int{}
	for i := 0; i < len(a)-1; i++ {
		if a[i] != a[i+1] {
			temp = append(temp, a[i])
		}
	}
	temp = append(temp, a[len(a)-1])
	return temp
}

func iterateSlices(a *[]int, b *[]int) []int {
	s := []int{}
	for _, v := range *a {
		isEqual := true
		for _, k := range *b {
			if v == k {
				isEqual = false
			}
		}
		if isEqual == true {
			s = append(s, v)
		}
	}
	return s
}

// FindMissingElement returns the missing element from two different slices
func FindMissingElement(a []int, b []int) []int {
	// var s []int
	// for _, v := range a {
	// 	isEqual := true
	// 	for _, k := range b {
	// 		if v == k {
	// 			isEqual = false
	// 		}
	// 	}
	// 	if isEqual == true {
	// 		s = append(s, v)
	// 	}
	// }
	// for _, v := range b {
	// 	isEqual := true
	// 	for _, k := range a {
	// 		if v == k {
	// 			isEqual = false
	// 		}
	// 	}
	// 	if isEqual == true {
	// 		s = append(s, v)
	// 	}
	// }
	var s []int
	s = iterateSlices(&a, &b)
	s = append(s, iterateSlices(&b, &a)...)
	sort.Ints(s)
	s = removeDuplicates(s)
	return s
}

func main() {
	a1 := []int{8, 7, 23, 9, 2, 0, 5, 24, 24, 25, 25, 20, -1, 0}
	a2 := []int{0, 23, 8, 9, 2, 7, 24, 24, 29, 35, 35, 3535, 3535, 100, 0, 0, -40, -10000}

	fmt.Println(FindMissingElement(a1, a2))
}
