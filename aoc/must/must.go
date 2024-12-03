package must

func Succeed(err error) {
	if err != nil {
		panic(err)
	}
}

func Return[T any](v T, err error) T {
	Succeed(err)
	return v
}
