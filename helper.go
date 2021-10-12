package config

import (
	"io"
)

func CloseResources(closer io.Closer) {
	if err := closer.Close(); err != nil {
		panic(err)
	}
}
