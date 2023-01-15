package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testPatterns := map[string]struct {
		a      []int
		b      []int
		expect []int
	}{
		"no elems": {
			expect: []int{},
		},
		"normal": {
			a:      []int{1, 3, 5},
			b:      []int{2, 4, 6},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		"a==b": {
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3},
			expect: []int{1, 2, 3},
		},
		"a is empty": {
			a:      []int{},
			b:      []int{1, 2, 3},
			expect: []int{1, 2, 3},
		},
		"b is empty": {
			a:      []int{1, 2, 3},
			b:      []int{},
			expect: []int{1, 2, 3},
		},
	}
	for k, v := range testPatterns {
		t.Run(k, func(t *testing.T) {
			actual := Sum(v.a, v.b, func(i int) int { return i })
			assert.ElementsMatch(t, v.expect, actual)
		})
	}
}

func TestIntersect(t *testing.T) {
	testPatterns := map[string]struct {
		a      []int
		b      []int
		expect []int
	}{
		"no elems": {
			expect: []int{},
		},
		"a==b": {
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3},
			expect: []int{1, 2, 3},
		},
		"odd and even": {
			a:      []int{1, 3, 5},
			b:      []int{2, 4, 6},
			expect: []int{},
		},
	}
	for k, v := range testPatterns {
		t.Run(k, func(t *testing.T) {
			actual := Intersect(v.a, v.b, func(item int) int { return item })
			assert.ElementsMatch(t, v.expect, actual)
		})
	}
}
func TestDifference(t *testing.T) {
	testPatterns := map[string]struct {
		a      []int
		b      []int
		expect []int
	}{
		"no elems": {
			a:      []int{},
			b:      []int{},
			expect: []int{},
		},
		"a==b": {
			a:      []int{1, 2, 3, 4},
			b:      []int{1, 2, 3, 4},
			expect: []int{},
		},
		"|a-b|=1": {
			a:      []int{1, 2, 3},
			b:      []int{1, 2},
			expect: []int{3},
		},
		"|a|-|b|<0": {
			a:      []int{1, 2, 3, 4},
			b:      []int{1, 2, 3, 4, 5},
			expect: []int{},
		},
	}
	for k, v := range testPatterns {
		t.Run(k, func(t *testing.T) {
			actual := Difference(v.a, v.b, func(key int) int { return key })
			assert.ElementsMatch(t, v.expect, actual)
		})
	}
}

func TestSetBy(t *testing.T) {
	testPatterns := map[string]struct {
		input  []int
		expect map[int]int
	}{
		"no elems": {
			input:  []int{1, 2, 3},
			expect: map[int]int{1: 1, 2: 2, 3: 3},
		},
	}
	for k, v := range testPatterns {
		t.Run(k, func(t *testing.T) {
			actual := SetBy(v.input, func(i int) int { return i })
			assert.Equal(t, v.expect, actual)
		})
	}

}
