package slicediff

import "testing"

type TestData struct {
	id   int
	in1  []int
	in2  []int
	want []int
}

var cases = []TestData{
	{
		1,
		[]int{},
		[]int{},
		[]int{},
	},
	{
		2,
		[]int{-1, -1, 2, 2, -10000, 500},
		[]int{-1, -1, 2, 3, -1000, 0, -10000},
		[]int{-1000, 0, 3, 500},
	},
	{
		3,
		[]int{-1000},
		[]int{1000},
		[]int{-1000, 1000},
	},
	{
		4,
		[]int{0, 1, 2, 3},
		[]int{1, 2, 3, 0},
		[]int{},
	},
	{
		5,
		[]int{1, 1, 1},
		[]int{2, 2, 2},
		[]int{1, 2},
	},
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFindMissingElements(t *testing.T) {

	for _, c := range cases {
		got := FindMissingElements(c.in1, c.in2)
		if !isEqual(got, c.want) {
			t.Errorf("Test No%v: FindMissingElements(%v, %v) == %v, want %v", c.id, c.in1, c.in2, got, c.want)
		}
	}

}

func TestFindMissingElements1(t *testing.T) {

	for _, c := range cases {
		got := FindMissingElements1(c.in1, c.in2)
		if !isEqual(got, c.want) {
			t.Errorf("Test No%v: FindMissingElements1(%v, %v) == %v, want %v", c.id, c.in1, c.in2, got, c.want)
		}
	}

}

func TestFindMissingElements2(t *testing.T) {

	for _, c := range cases {
		got := FindMissingElements2(c.in1, c.in2)
		if !isEqual(got, c.want) {
			t.Errorf("Test No%v: FindMissingElements2(%v, %v) == %v, want %v", c.id, c.in1, c.in2, got, c.want)
		}
	}

}
