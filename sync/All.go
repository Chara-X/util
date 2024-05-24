package sync

func All[T any](channels ...chan T) []T {
	var values = make([]T, len(channels))
	for i, ch := range channels {
		values[i] = <-ch
	}
	return values
}
