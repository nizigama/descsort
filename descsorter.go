package descsort

import "sort"

// DescInts sorts a slice of ints in descending order
func DescInts(xi []int) {

	sort.Sort(sort.Reverse(sort.IntSlice(xi))) // working
	// sort.Sort(sort.Reverse(sort.IntSlice([]int{2, 8, 0, 3, 5}))) // failing
}

// DescFloat64s sorts a slice of float64s in descending order
func DescFloat64s(xf []float64) {

	sort.Sort(sort.Reverse(sort.Float64Slice(xf))) //working
	// sort.Sort(sort.Reverse(sort.Float64Slice([]float64{2.1, 8.5, 0.2, 3.4, 5.9}))) // failing
}

// DescStrings sorts a slice of strings in descending order
func DescStrings(xs []string) {

	sort.Sort(sort.Reverse(sort.StringSlice(xs))) //working
	// sort.Sort(sort.Reverse(sort.StringSlice([]string{"a", "b", "z", "r", "f"}))) // failing
}
