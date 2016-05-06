package sorting_test

import (
	"fmt"
	"sort"

	"github.com/sasha-s/go-sorting"
)

func ExampleSwap() {
	names := [][]string{[]string{"Max"}, []string{"Ann", "Alex"}, []string{"Miles", "Ann"}}
	grade := []int{12, 7, 9}

	sorting.By(sort.IntSlice(grade)).Sort(sorting.SwapFunc(func(i, j int) { names[i], names[j] = names[j], names[i] }))

	fmt.Println("By grade:", names, grade)

	// Output:
	// By grade: [[Ann Alex] [Miles Ann] [Max]] [7 9 12]
}
