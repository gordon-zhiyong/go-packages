package main

import (
	"fmt"
	"sort"
)

type myIntType []int

func (m myIntType) Len() int {
	return len(m)
}

func (m myIntType) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m myIntType) Less(i, j int) bool {
	return m[i] < m[j]
}

func main() {
	intSlice := myIntType{
		4, 7, 1, 0, 9, 1, 6, 3, 4,
	}

	// IsSorted reports whether data is sorted.
	if !sort.IsSorted(intSlice) {
		fmt.Println("intSlice is not sorted:", intSlice)
	}

	// Sort the intSlice which is of type myIntType
	//
	// Sort sorts data.
	// It makes one call to data.Len to determine n, and
	// O(n*log(n)) calls to data.Less and data.Swap.
	// The sort is not guaranteed to be stable.
	sort.Sort(intSlice)

	if sort.IsSorted(intSlice) {
		fmt.Println("intSlice is now sorted:", intSlice)
	}
}
