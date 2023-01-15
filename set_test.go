package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	testPatterns := map[string]struct {
		a      []int
		b      []int
		expect []int
	}{
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
			assert.Equal(t, v.expect, actual)
		})
	}
}
