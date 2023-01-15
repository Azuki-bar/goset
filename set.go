package goset

func Difference[T any, R comparable](aSlice, bSlice []T, key func(key T) R) []T {
	aMap := make(map[R]T)
	bMap := make(map[R]T)
	for _, a := range aSlice {
		key := key(a)
		if _, ok := aMap[key]; !ok {
			aMap[key] = a
		}
	}
	for _, b := range bSlice {
		key := key(b)
		if _, ok := bMap[key]; !ok {
			bMap[key] = b
		}
	}
	diff := make([]T, 0)
	for ka, va := range aMap {
		if _, ok := bMap[ka]; !ok {
			diff = append(diff, va)
		}
	}
	return diff
}
