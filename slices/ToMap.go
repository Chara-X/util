package slices

func ToMap[K comparable, V any](slice []V, keySelector func(V) K) map[K]V {
	var set = make(map[K]V)
	for _, v := range slice {
		set[keySelector(v)] = v
	}
	return set
}
