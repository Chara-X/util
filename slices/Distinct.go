package slices

func Distinct[T comparable](slice []T) []T {
	var res []T
	var set = map[T]struct{}{}
	for _, v := range slice {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}
