package main

import (
	"fmt"
	"github.com/pefish/go-error"
)

func main() {
	defer go_error.Recover(func(msg string, internalMsg string, code uint64, data interface{}) {
		fmt.Println(msg, internalMsg, code, data)
	})

	//go_error.ThrowErrorWithData(`test`, 12, map[string]interface{}{
	//	`a`: 1,
	//}, errors.New(`haha`))
	//panic(`111`)
	//panic(errors.New(`haha`))
	var a interface{} = `aa`
	b := a.(float64)
	fmt.Println(b)
}
