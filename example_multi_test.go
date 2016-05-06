// Adapted from https://golang.org/src/sort/example_multi_test.go
// Original copyright:
//
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sorting_test

import (
	"fmt"
	"sort"

	"github.com/sasha-s/go-sorting"
)

// A Change has records of source code changes, recording user, language, and delta size.
type Changes struct {
	user     sort.StringSlice
	language sort.StringSlice
	lines    sort.IntSlice
}

// A Change is a record of source code changes, recording user, language, and delta size.
type Change struct {
	user     string
	language string
	lines    int
}

var changes = Changes{
	user:     sort.StringSlice{`gri`, `ken`, `glenda`, `rsc`, `r`, `ken`, `dmr`, `r`, `gri`},
	language: sort.StringSlice{`Go`, `C`, `Go`, `Go`, `Go`, `Go`, `C`, `C`, `Smalltalk`},
	lines:    sort.IntSlice{100, 150, 200, 200, 100, 200, 100, 150, 80},
}

func ExampleSortMultiKeys() {
	// Simple use: Sort by user.
	sorting.By(changes.user).Sort(changes.language, changes.lines)
	fmt.Println("By user:", changes)

	// Another way to do the same.
	sorting.Sort(changes.language, changes.lines).By(changes.user)
	fmt.Println("By user:", changes)

	// More examples.
	sorting.By(changes.user, changes.lines).Sort(changes.language)
	fmt.Println("By user,<lines:", changes)

	sorting.By(changes.user, sort.Reverse(changes.lines)).Sort(changes.language)
	fmt.Println("By user,>lines:", changes)

	sorting.By(changes.language, changes.lines).Sort(changes.user)
	fmt.Println("By language,<lines:", changes)

	sorting.By(changes.language, changes.lines, changes.user).Sort()
	fmt.Println("By language,<lines,user:", changes)

	// Another way to do the same.
	sort.Sort(sorting.Lex(changes.language, changes.lines, changes.user))
	fmt.Println("By language,<lines,user:", changes)

	sort.Sort(sorting.Lex(changes.user, changes.lines, changes.language))
	fmt.Println("By user,<lines,language:", changes)

	// Output:
	// By user: {[dmr glenda gri gri ken ken r r rsc] [C Go Go Smalltalk C Go Go C Go] [100 200 100 80 150 200 100 150 200]}
	// By user: {[dmr glenda gri gri ken ken r r rsc] [C Go Go Smalltalk C Go Go C Go] [100 200 100 80 150 200 100 150 200]}
	// By user,<lines: {[dmr glenda gri gri ken ken r r rsc] [C Go Smalltalk Go C Go Go C Go] [100 200 80 100 150 200 100 150 200]}
	// By user,>lines: {[dmr glenda gri gri ken ken r r rsc] [C Go Go Smalltalk Go C C Go Go] [100 200 100 80 200 150 150 100 200]}
	// By language,<lines: {[dmr ken r r gri ken glenda rsc gri] [C C C Go Go Go Go Go Smalltalk] [100 150 150 100 100 200 200 200 80]}
	// By language,<lines,user: {[dmr ken r gri r glenda ken rsc gri] [C C C Go Go Go Go Go Smalltalk] [100 150 150 100 100 200 200 200 80]}
	// By language,<lines,user: {[dmr ken r gri r glenda ken rsc gri] [C C C Go Go Go Go Go Smalltalk] [100 150 150 100 100 200 200 200 80]}
	// By user,<lines,language: {[dmr glenda gri gri ken ken r r rsc] [C Go Smalltalk Go C Go Go C Go] [100 200 80 100 150 200 100 150 200]}
}
