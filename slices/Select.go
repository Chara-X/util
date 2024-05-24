package slices

func Select[T, S any](slice []T, selector func(value T) S) []S {
	var res = []S{}
	for _, v := range slice {
		res = append(res, selector(v))
	}
	return res
}
