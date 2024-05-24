package sync

// Context:
//
//	sync.WaitGroup: RWLocker
//	chan struct{}: Bounded concurrency
func Actor[T any](jobs <-chan func(ctx T), ctx T) {
	go func() {
		for job := range jobs {
			job(ctx)
		}
	}()
}
