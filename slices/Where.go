package slices

func Where[T any](slice []T, predicate func(T) bool) []T {
	var res = []T{}
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}
