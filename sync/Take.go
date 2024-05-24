package sync

func Take[T any](channel <-chan T, count int) []T {
	var values = make([]T, count)
	for i := 0; i < count; i++ {
		values[i] = <-channel
	}
	return values
}
