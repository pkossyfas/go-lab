/*
Package slicediff provides three different functions (FindMissingElements, FindMissingElements1, FindMissingElements2)
that return the diff of two given slices (unique elements that missing from the one or the other).
Functions take as input two slices of ints and return one slice of ints.
*/
package slicediff

import "sort"

// keepUniques returns a slice with the unique items of a given sorted slice of ints
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

// removeDuplicates removes duplicate values from a given sorted slice of ints
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

// FindMissingElements2 returns a slice with the elements that are not present on both given slices of ints
// FindMissingElements2 has O(n*k) time complexity
func FindMissingElements2(a []int, b []int) []int {
	s := []int{}
	// iterate slice a over slice b
	s = iterateSlices(&a, &b)
	// iterate slice b over slice a and add missing elements to slice s
	s = append(s, iterateSlices(&b, &a)...)
	sort.Ints(s)
	s = removeDuplicates(s)

	return s
}

// FindMissingElements1 returns a slice with the elements that are not present on both given slices of ints
// FindMissingElements1 has O(nlogn + klogk) time complexity
func FindMissingElements1(a []int, b []int) []int {
	// sort and remove duplicates values the two slices
	sort.Ints(a)
	sort.Ints(b)
	a = removeDuplicates(a)
	b = removeDuplicates(b)

	// create a new slice by adding the two slices and return a slice with the unique values
	s := []int{}
	s = append(a, b...)
	sort.Ints(s)

	return keepUniques(s)
}

// FindMissingElements returns a slice with the elements that are not present on both given slices of ints
// FindMissingElements has O(nlogn + klogk) time complexity and it is consider the best solution
func FindMissingElements(a []int, b []int) []int {
	// sort the two given slices
	sort.Ints(a)
	sort.Ints(b)
	// remove the duplicates
	a = removeDuplicates(a)
	b = removeDuplicates(b)

	s := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else if a[i] > b[j] {
			s = append(s, b[j])
			j++
		} else {
			s = append(s, a[i])
			i++
		}
	}
	s = append(s, a[i:]...)
	s = append(s, b[j:]...)

	return s
}
