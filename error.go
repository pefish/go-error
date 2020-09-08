package go_error

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorInfo struct {
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
	Err  error       `json:"err"`
}

// 供外部使用
var (
	INTERNAL_ERROR             = `internal error`
	INTERNAL_ERROR_CODE uint64 = 1
)

func (ei *ErrorInfo) String() string {
	return fmt.Sprintf(`ErrorInfo -> msg: %s, code: %d, data: %#v, err: %#v`, ei.Err.Error(), ei.Code, ei.Data, ei.Err)
}

func (ei *ErrorInfo) Error() string {
	return ei.String()
}

func (ei *ErrorInfo) Unwrap() error {
	return ei.Err
}

func Recover(fun func(*ErrorInfo)) {
	if err := recover(); err != nil {
		switch e := err.(type) {
		case *ErrorInfo:
			fun(e)
		case error:
			fun(WrapWithAll(e, INTERNAL_ERROR_CODE, nil))
		default:
			fun(WrapWithAll(errors.New(fmt.Sprint(err)), INTERNAL_ERROR_CODE, nil))
		}
	}
}

func ThrowInternal(err error) {
	var errorInfo_ = ErrorInfo{INTERNAL_ERROR_CODE, nil, err}
	panic(&errorInfo_)
}

func ThrowWithCode(err error, code uint64) {
	var errorInfo_ = ErrorInfo{code, nil, err}
	panic(&errorInfo_)
}

func ThrowWithCodeAndData(err error, code uint64, data error) {
	var errorInfo_ = ErrorInfo{code, data, err}
	panic(&errorInfo_)
}

func WrapWithAll(err error, code uint64, data interface{}) *ErrorInfo {
	return &ErrorInfo{code, data, err}
}

func WrapWithErr(err error) *ErrorInfo {
	return &ErrorInfo{INTERNAL_ERROR_CODE, nil, err}
}

func Wrap(err error) *ErrorInfo {
	return &ErrorInfo{INTERNAL_ERROR_CODE, nil, err}
}

func WrapWithCode(err error, code uint64) *ErrorInfo {
	return &ErrorInfo{code, nil, err}
}
