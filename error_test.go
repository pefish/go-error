package go_error

import (
	"fmt"
	go_test_ "github.com/pefish/go-test"
	"testing"
)

func TestErrorInfo_Test1(t *testing.T) {
	defer Recover(func(errorInfo *ErrorInfo) {
		go_test_.Equal(t, true, errorInfo.Code == 1)
		go_test_.Equal(t, `interface conversion: interface {} is string, not float64`, errorInfo.Err.Error())
	})

	var a interface{} = `aa`
	b := a.(float64)
	fmt.Println(b)
}
