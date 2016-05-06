package sorting_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/sasha-s/go-sorting"
)

func TestIdxFloats(t *testing.T) {
	s := []float64{1., 3., 2.}
	idx := sorting.SortIdx(sort.Float64Slice(s))
	e := []int{0, 2, 1}
	for i := range e {
		if actual, expected := idx[i], e[i]; actual != expected {
			t.Error(i, actual, expected)
		}
	}
	es := []float64{1., 2., 3.}
	for i := range es {
		if actual, expected := s[i], es[i]; actual != expected {
			t.Error(i, actual, expected)
		}
	}
}

func TestIdxStrings(t *testing.T) {
	s := []string{"a1", "b", "a"}
	idx := sorting.SortIdx(sort.StringSlice(s))
	e := []int{2, 0, 1}
	for i := range e {
		if actual, expected := idx[i], e[i]; actual != expected {
			t.Error(i, actual, expected)
		}
	}
	es := []string{"a", "a1", "b"}
	for i := range es {
		if actual, expected := s[i], es[i]; actual != expected {
			t.Error(i, actual, expected)
		}
	}
}

func TestLex(t *testing.T) {
	name := []string{"Max", "Ann", "Alex", "Miles", "Ann"}
	age := []int{7, 5, 12, 3, 3}
	sort.Sort(sorting.Lex(sort.StringSlice(name), sort.IntSlice(age)))
	expectedNames := []string{`Alex`, `Ann`, `Ann`, `Max`, `Miles`}
	expectedAges := []int{12, 3, 5, 7, 3}
	if !reflect.DeepEqual(name, expectedNames) {
		t.Errorf("expected %v, got %v", expectedNames, name)
	}
	if !reflect.DeepEqual(age, expectedAges) {
		t.Errorf("expected %v, got %v", expectedAges, age)
	}
}

func TestLexReverse(t *testing.T) {
	name := []string{"Max", "Ann", "Alex", "Miles", "Ann"}
	age := []int{7, 5, 12, 3, 3}
	sort.Sort(sorting.Lex(sort.Reverse(sort.StringSlice(name)), sort.IntSlice(age)))
	expectedNames := []string{`Miles`, `Max`, `Ann`, `Ann`, `Alex`}
	expectedAges := []int{3, 7, 3, 5, 12}
	if !reflect.DeepEqual(name, expectedNames) {
		t.Errorf("expected %v, got %v", expectedNames, name)
	}
	if !reflect.DeepEqual(age, expectedAges) {
		t.Errorf("expected %v, got %v", expectedAges, age)
	}
}

func TestLexReverse2(t *testing.T) {
	name := []string{"Max", "Ann", "Alex", "Miles", "Ann"}
	age := []int{7, 5, 12, 3, 3}
	sort.Sort(sort.Reverse(sorting.Lex(sort.StringSlice(name), sort.IntSlice(age))))
	expectedNames := []string{`Miles`, `Max`, `Ann`, `Ann`, `Alex`}
	expectedAges := []int{3, 7, 5, 3, 12}
	if !reflect.DeepEqual(name, expectedNames) {
		t.Errorf("expected %v, got %v", expectedNames, name)
	}
	if !reflect.DeepEqual(age, expectedAges) {
		t.Errorf("expected %v, got %v", expectedAges, age)
	}
}

func TestLexAge(t *testing.T) {
	name := []string{"Max", "Ann", "Alex", "Miles", "Ann"}
	age := []int{7, 5, 12, 3, 3}
	sort.Sort(sorting.Lex(sort.IntSlice(age), sort.StringSlice(name)))
	expectedNames := []string{`Ann`, `Miles`, `Ann`, `Max`, `Alex`}
	expectedAges := []int{3, 3, 5, 7, 12}
	if !reflect.DeepEqual(name, expectedNames) {
		t.Errorf("expected %v, got %v", expectedNames, name)
	}
	if !reflect.DeepEqual(age, expectedAges) {
		t.Errorf("expected %v, got %v", expectedAges, age)
	}
}
