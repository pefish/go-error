package go_error

type ErrorInfo struct {
	ErrorMessage string
	ErrorCode uint64
	Data interface{}
	Err error
}

// 供外部使用
var (
	INTERNAL_ERROR = `internal error`
	INTERNAL_ERROR_CODE uint64 = 1
)


func (e *ErrorInfo) Error() string {
	return e.ErrorMessage
}


func ThrowInternal(text string) {
	var errorInfo_ = ErrorInfo{text, INTERNAL_ERROR_CODE, nil,nil}
	panic(errorInfo_)
}

func ThrowInternalError(text string, err error) {
	var errorInfo_ = ErrorInfo{text, INTERNAL_ERROR_CODE, nil,err}
	panic(errorInfo_)
}

func Throw(text string, code uint64) {
	var errorInfo_ = ErrorInfo{text, code, nil,nil}
	panic(errorInfo_)
}

func ThrowWithData(text string, code uint64, data interface{}) {
	var errorInfo_ = ErrorInfo{text, code, data,nil}
	panic(errorInfo_)
}

func ThrowError(text string, code uint64, err error) {
	var errorInfo_ = ErrorInfo{text, code, nil,err}
	panic(errorInfo_)
}

func ThrowErrorWithData(text string, code uint64, data interface{}, err error) {
	var errorInfo_ = ErrorInfo{text, code, data,err}
	panic(errorInfo_)
}

