package slices

import (
	"cmp"
)

func MinBy[T any, K cmp.Ordered, V any](slice []T, keySelector func(T) K, valueSelector func(int, T) V) V {
	var key = keySelector(slice[0])
	var value = valueSelector(0, slice[0])
	for i, v := range slice {
		var k = keySelector(v)
		if k < key {
			key = k
			value = valueSelector(i, v)
		}
	}
	return value
}
