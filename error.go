package p_error

import (
	"fmt"
	"github.com/pefish/go-logger"
)

type ErrorInfo struct {
	ErrorMessage string
	ErrorCode uint64
	Data interface{}
	Err error
}

var (
	INTERNAL_ERROR = `internal error`

	CODE_ERROR uint64 = 0
)

func (e *ErrorInfo) Error() string {
	return e.ErrorMessage
}


func ThrowInternal(text string) {
	var errorInfo_ = ErrorInfo{text, 1, nil,nil}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowInternalError(text string, err error) {
	var errorInfo_ = ErrorInfo{text, 1, nil,err}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func Throw(text string, code uint64) {
	var errorInfo_ = ErrorInfo{text, code, nil,nil}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowWithData(text string, code uint64, data interface{}) {
	var errorInfo_ = ErrorInfo{text, code, data,nil}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowError(text string, code uint64, err error) {
	var errorInfo_ = ErrorInfo{text, code, nil,err}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

func ThrowErrorWithData(text string, code uint64, data interface{}, err error) {
	var errorInfo_ = ErrorInfo{text, code, data,err}
	p_logger.Logger.Error(fmt.Sprintf(`ERROR: %v`, errorInfo_))
	panic(errorInfo_)
}

