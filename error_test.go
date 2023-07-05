package go_error

import (
	"fmt"
	"github.com/pefish/go-test-assert"
	"testing"
)

func TestErrorInfo_Test1(t *testing.T) {
	defer Recover(func(errorInfo *ErrorInfo) {
		test.Equal(t, true, errorInfo.Code == 1)
		test.Equal(t, `interface conversion: interface {} is string, not float64`, errorInfo.Err.Error())
	})

	var a interface{} = `aa`
	b := a.(float64)
	fmt.Println(b)
}
