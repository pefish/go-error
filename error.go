package go_error

import (
	"errors"
	"fmt"
)

type ErrorInfo struct {
	Msg  string      `json:"msg"`
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
	Err  error       `json:"err,omitempty"`
}

// 供外部使用
var (
	INTERNAL_ERROR             = `internal error`
	INTERNAL_ERROR_CODE uint64 = 1
)

func (ei *ErrorInfo) String() string {
	return fmt.Sprintf(`ErrorInfo -> msg: %s, code: %d, data: %#v, err: %#v`, ei.Msg, ei.Code, ei.Data, ei.Err)
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
			fun(WrapWithAll(e.Error(), INTERNAL_ERROR_CODE, nil, e))
		default:
			msg := fmt.Sprint(err)
			fun(WrapWithAll(msg, INTERNAL_ERROR_CODE, nil, errors.New(msg)))
		}
	}
}

func ThrowInternal(text string) {
	var errorInfo_ = ErrorInfo{text, INTERNAL_ERROR_CODE, nil, nil}
	panic(&errorInfo_)
}

func ThrowInternalWithError(text string, err error) {
	panic(WrapWithErr(text, err))
}

func ThrowWithCode(text string, code uint64) {
	var errorInfo_ = ErrorInfo{text, code, nil, nil}
	panic(&errorInfo_)
}

func ThrowWithCodeAndData(text string, code uint64, data error) {
	var errorInfo_ = ErrorInfo{text, code, data, nil}
	panic(&errorInfo_)
}

func ThrowWithCodeAndErr(text string, code uint64, err error) {
	var errorInfo_ = ErrorInfo{text, code, nil, err}
	panic(&errorInfo_)
}

func ThrowWithAll(text string, code uint64, data interface{}, err error) {
	panic(WrapWithAll(text, code, data, err))
}

func WrapWithAll(text string, code uint64, data interface{}, err error) *ErrorInfo {
	return &ErrorInfo{text, code, data, err}
}

func WrapWithErr(text string, err error) *ErrorInfo {
	return &ErrorInfo{text, INTERNAL_ERROR_CODE, nil, err}
}
