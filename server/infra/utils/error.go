package utils

func Chain(fns ...func()error) (err error) {
	for _, fn := range fns {
		if err = fn(); err != nil {
			break
		}
	}
	return
}
