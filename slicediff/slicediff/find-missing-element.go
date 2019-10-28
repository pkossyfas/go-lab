package slicediff

import (
	"sort"
)

// keepUniques removes duplicate values from a given slice of ints
func keepUniques(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}
	temp := []int{}
	if a[0] != a[1] {
		temp = append(temp, a[0])
	}
	i := 1
	for ; i < len(a)-1; i++ {
		if a[i] != a[i-1] && a[i] != a[i+1] {
			temp = append(temp, a[i])
		}
	}
	if a[i] != a[i-1] {
		temp = append(temp, a[i])
	}

	return temp
}

// removeDuplicates removes duplicate values from a given slice of ints
func removeDuplicates(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}
	temp := []int{}
	i := 0
	for ; i < len(a)-1; i++ {
		if a[i] != a[i+1] {
			temp = append(temp, a[i])
		}
	}
	temp = append(temp, a[i])

	return temp
}

// iterateSlices returns a slice with integers from an array which are not
// present to the second given array
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

// FindMissingElement returns a slice with the missing elements from
// two different slices of ints
func FindMissingElement(a []int, b []int) []int {
	s := []int{}
	// iterate slice a over slice b
	s = iterateSlices(&a, &b)
	// iterate slice b over slice a and add missing elements to slice s
	s = append(s, iterateSlices(&b, &a)...)
	sort.Ints(s)
	s = removeDuplicates(s)

	return s
}

// FindMissingElement2 returns a slice with the missing elements from
// two different slices of ints
func FindMissingElement2(a []int, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)
	a = removeDuplicates(a)
	b = removeDuplicates(b)
	s := []int{}
	s = append(a, b...)
	sort.Ints(s)

	return keepUniques(s)
}
