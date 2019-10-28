package main

import (
	"fmt"

	"github.com/pkossyfas/go-lab/slicediff/slicediff"
)

func main() {
	a1 := []int{8, 7, 23, 9, 2, 0, 5, 24, 24, 25, 25, 20, -1, 0}
	a2 := []int{0, 23, 8, 9, 2, 7, 24, 24, 29, 35, 35, 3535, 3535, 100, 0, 0, -40, -10000}

	fmt.Println(slicediff.FindMissingElement(a1, a2))
	fmt.Println(slicediff.FindMissingElement2(a1, a2))
}
