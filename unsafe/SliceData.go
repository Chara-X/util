package unsafe

func SliceData[T any](slice []T) *T { return &slice[0] }
