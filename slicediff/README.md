# Slicediff package

Package **slicediff** provides three different functions which provide solution to the same problem:

>Given two slices of integers, find the elements that are unique per slice.

*Example:*

>Input: [8, 7, 23, 9, 2, 0] [0, 23, 8, 9, 2] -> Output: [7]

The three signature functions are the following:

```go
func FindMissingElements([]int, []int) []int
// time complexity: O(nlogn)

func FindMissingElements1([]int, []int) []int
// time complexity: O(nlogn)

func FindMissingElements2([]int, []int) []int
// time complexity: O(n*n)
```

## Installation

Use the `go` command:

```bash
go get -u gitlab.internal.upstreamsystems.com/polykarpos.kossyfas/go-lab/slicediff
```

## Tests

You may test the package by running (inside the same folder):

```bash
go test
```

## Example

```go
package main

import (
    "fmt"

    "gitlab.internal.upstreamsystems.com/polykarpos.kossyfas/go-lab/slicediff"
)

func main() {
    // Creating two slices of
    s1 := []int{1, 3, 5}
    s2 := []int{-2, 1, 1, 3}

    fmt.Println("Input:")
    fmt.Println("------")
    fmt.Println(s1)
    fmt.Println(s2)

    // Find the missing elements (values that are not exist on both slices)
    s := slicediff.FindMissingElements(s1, s2)

    fmt.Println()
    fmt.Println("Output (missing elements (values that are not exist on both slices)):")
    fmt.Println("---------------------------------------------------------------------")
    fmt.Println(s)
}
```
