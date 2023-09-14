package must

import (
	"fmt"
	"io"
)

func NotError(err error) {
	if err != nil {
		panic(fmt.Errorf("must: unexpected error: %w", err))
	}
}

func Succeed[T any](value T, err error) T {
	NotError(err)

	return value
}

func Close(c io.Closer) {
	NotError(c.Close())
}
