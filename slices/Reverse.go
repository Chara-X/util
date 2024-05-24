package slices

func Reverse[T any](slice []T) []T {
	var res = make([]T, len(slice))
	for i, v := range slice {
		res[len(slice)-i-1] = v
	}
	return res
}
