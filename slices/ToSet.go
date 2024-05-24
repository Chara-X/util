package slices

func ToSet[T comparable](slice []T) map[T]bool {
	var set = map[T]bool{}
	for _, v := range slice {
		set[v] = true
	}
	return set
}
