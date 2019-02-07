package p_error

import (
	"fmt"
	"gitee.com/pefish/p-go-logger"
)

type ErrorInfo struct {
	ErrorMessage string
	ErrorCode int64
	Data interface{}
	Err error
}

func (e *ErrorInfo) Error() string {
	return e.ErrorMessage
}


func ThrowInternal(text string) {
	var errorInfo_ = ErrorInfo{text, 0, nil,nil}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func Throw(text string, code int64) {
	var errorInfo_ = ErrorInfo{text, code, nil,nil}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowWithData(text string, code int64, data interface{}) {
	var errorInfo_ = ErrorInfo{text, code, data,nil}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowError(text string, code int64, err error) {
	var errorInfo_ = ErrorInfo{text, code, nil,err}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowErrorWithData(text string, code int64, data interface{}, err error) {
	var errorInfo_ = ErrorInfo{text, code, data,err}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

