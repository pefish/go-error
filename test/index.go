package main

import (
	"errors"
	"fmt"
	"github.com/pefish/go-error"
)

func main() {
	defer go_error.Recover(func(msg string, internalMsg string, code uint64, data interface{}, err interface{}) {
		fmt.Println(msg, internalMsg, code, data)
		fmt.Println(fmt.Sprint(err))
	})
	fmt.Printf("%s\n", &go_error.ErrorInfo{
		ErrorCode:    go_error.INTERNAL_ERROR_CODE,
		ErrorMessage: `post error`,
		Err: errors.New(`haha`),
	})
	go_error.ThrowError(`haha`, 1000, errors.New(`hehe`))
	//go_error.ThrowErrorWithData(`test`, 12, map[string]interface{}{
	//	`a`: 1,
	//}, errors.New(`haha`))
	//panic(`111`)
	//panic(errors.New(`haha`))
	var a interface{} = `aa`
	b := a.(float64)
	fmt.Println(b)
}
