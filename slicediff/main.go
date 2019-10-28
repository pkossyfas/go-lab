package main

import (
	"fmt"

	"github.com/pkossyfas/go-lab/slicediff/slicediff"
)

func main() {
	a1 := []int{2, 2}
	a2 := []int{2}
	fmt.Println(a1, a2)
	fmt.Println(slicediff.FindMissingElement(a1, a2))
	fmt.Println(slicediff.FindMissingElement2(a1, a2))
	fmt.Println(slicediff.FindMissingElement3(a1, a2))
	fmt.Println("-------------------------------------")
}
