package slices

func Intersect[T comparable](slices ...[]T) []T {
	var counts = map[T]int{}
	for _, slice := range slices {
		for _, value := range slice {
			counts[value]++
		}
	}
	var res []T
	for value, count := range counts {
		if count == len(slices) {
			res = append(res, value)
		}
	}
	return res
}
