# go-sorting
sorting parallel slices in golang. [GoDoc](https://godoc.org/github.com/sasha-s/go-sorting).

Lexicographical ordering:
```go
package main

import (
	"fmt"
	"sort"

	"github.com/sasha-s/si-srv/sorting"
)

func main() {
	name := []string{"Max", "Ann", "Alex", "Miles", "Ann"}
	age := []int{7, 5, 12, 3, 3}

	sort.Sort(sorting.Lex(sort.IntSlice(age), sort.StringSlice(name)))

	fmt.Println(name, age)

	// Output:
	// [Ann Miles Ann Max Alex] [3 3 5 7 12]
}
```

Sort by grade, reordering names as we go:
```go
names := [][]string{[]string{"Max"}, []string{"Ann", "Alex"}, []string{"Miles", "Ann"}}
grade := []int{12, 7, 9}

sorting.By(sort.IntSlice(grade)).Sort(sorting.SwapFunc(func(i, j int) { names[i], names[j] = names[j], names[i] }))

fmt.Println("By grade:", names, grade)

// Output:
// By grade: [[Ann Alex] [Miles Ann] [Max]] [7 9 12]

```

Indices:
```go
s := []string{"a1", "b", "a"}
idx := sorting.SortIdx(sort.StringSlice(s))
fmt.Println(s, idx)

// Output:
[a a1 b] [2 0 1]
```

Also see [example](https://github.com/sasha-s/go-sorting/blob/master/example_multi_test.go).
