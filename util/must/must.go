package must

import "fmt"

func Succeed[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must: unexpected error: %w", err))
	}

	return value
}
