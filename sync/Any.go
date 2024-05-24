package sync

import "reflect"

func Any[T any](channels ...chan T) (T, bool) {
	var cases = []reflect.SelectCase{}
	for _, c := range channels {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c)})
	}
	var _, value, ok = reflect.Select(cases)
	return value.Interface().(T), ok
}
