package go_error

import (
	"errors"
	"fmt"
	"github.com/pefish/go-test-assert"
	"testing"
)

func TestErrorInfo_Test(t *testing.T) {
	defer Recover(func(errorInfo *ErrorInfo) {
		test.Equal(t, true, errorInfo.Code == 1000)
		test.Equal(t, `hehe`, errorInfo.Err.Error())
	})
	test.Equal(t, `ErrorInfo -> msg: haha, code: 1, data: <nil>, err: &errors.errorString{s:"haha"}`, fmt.Sprint(&ErrorInfo{
		Code: INTERNAL_ERROR_CODE,
		Err:  errors.New(`haha`),
	}))
	ThrowWithCode(errors.New(`hehe`), 1000)
}

func TestErrorInfo_Test1(t *testing.T) {
	defer Recover(func(errorInfo *ErrorInfo) {
		test.Equal(t, true, errorInfo.Code == 1)
		test.Equal(t, `interface conversion: interface {} is string, not float64`, errorInfo.Err.Error())
	})

	var a interface{} = `aa`
	b := a.(float64)
	fmt.Println(b)
}
