package goset

import "github.com/samber/lo"

// Sum returns set by sum.
// expressed by `a or b`.
func Sum[T any, R comparable](aSlice, bSlice []T, key func(key T) R) []T {
	flat := lo.Flatten([][]T{aSlice, bSlice})
	ma := SetBy(flat, key)
	return lo.Values(ma)
}

// Intersect returns intersection set expressed by slice.
// This function expressed by `aSlice and bSlice`.
func Intersect[T any, R comparable](aSlice, bSlice []T, key func(key T) R) []T {
	aMap := SetBy(aSlice, key)
	bMap := SetBy(bSlice, key)
	diff := make([]T, 0)
	for ka, va := range aMap {
		if _, ok := bMap[ka]; ok {
			diff = append(diff, va)
		}
	}
	return diff
}

// Difference returns set difference calclated by aSlice - bSlice expressed by slice.
func Difference[T any, R comparable](aSlice, bSlice []T, key func(key T) R) []T {
	aMap := SetBy(aSlice, key)
	bMap := SetBy(bSlice, key)
	diff := make([]T, 0)
	for ka, va := range aMap {
		if _, ok := bMap[ka]; !ok {
			diff = append(diff, va)
		}
	}
	return diff
}

// SetBy returns map structured by no duplicated elements.
func SetBy[T any, R comparable](slice []T, key func(key T) R) map[R]T {
	mapped := make(map[R]T)
	for _, elem := range slice {
		key := key(elem)
		if _, ok := mapped[key]; !ok {
			mapped[key] = elem
		}
	}
	return mapped
}
