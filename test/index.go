package main

import (
	"errors"
	"fmt"
	"github.com/pefish/go-error"
)

func main() {
	defer go_error.Recover(func(msg string, internalMsg string, code uint64, data interface{}, err interface{}) {
		fmt.Println(msg, internalMsg, code, data, err)
	})

	go_error.ThrowErrorWithData(`test`, 12, map[string]interface{}{
		`a`: 1,
	}, errors.New(`haha`))
}
