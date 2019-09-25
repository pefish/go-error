package go_error

type ErrorInfo struct {
	InternalErrorMessage string // 开发人员在测试环境才能看到的消息
	ErrorMessage         string // 用户也可以看到的消息，不可以带敏感信息
	ErrorCode            uint64
	Data                 interface{}
	Err                  error
}

// 供外部使用
var (
	INTERNAL_ERROR             = `internal error`
	INTERNAL_ERROR_CODE uint64 = 1
)

func (e *ErrorInfo) Error() string {
	return e.ErrorMessage
}

func Recover(fun func(msg string, internalMsg string, code uint64, data interface{}, err interface{})) {
	if err := recover(); err != nil {
		if _, ok := err.(ErrorInfo); !ok {
			msg := ``
			if _, ok := err.(error); !ok {
				msg = err.(string)
			} else {
				msg = err.(error).Error()
			}
			fun(msg, msg, INTERNAL_ERROR_CODE, nil, err)
		} else {
			errorInfoStruct := err.(ErrorInfo)
			fun(errorInfoStruct.ErrorMessage, errorInfoStruct.InternalErrorMessage, errorInfoStruct.ErrorCode, errorInfoStruct.Data, err)
		}
	}
}

func ThrowInternalWithInternalMsg(msg string, internalMsg string) {
	var errorInfo_ = ErrorInfo{internalMsg, msg, INTERNAL_ERROR_CODE, nil, nil}
	panic(errorInfo_)
}

func ThrowInternal(text string) {
	var errorInfo_ = ErrorInfo{text, text, INTERNAL_ERROR_CODE, nil, nil}
	panic(errorInfo_)
}

func ThrowInternalError(text string, err error) {
	var errorInfo_ = ErrorInfo{text, text, INTERNAL_ERROR_CODE, nil, err}
	panic(errorInfo_)
}

func ThrowInternalErrorWithInternalMsg(msg string, internalMsg string, err error) {
	var errorInfo_ = ErrorInfo{internalMsg, msg, INTERNAL_ERROR_CODE, nil, err}
	panic(errorInfo_)
}

func Throw(text string, code uint64) {
	var errorInfo_ = ErrorInfo{text, text, code, nil, nil}
	panic(errorInfo_)
}

func ThrowWithInternalMsg(msg string, internalMsg string, code uint64) {
	var errorInfo_ = ErrorInfo{internalMsg, msg, code, nil, nil}
	panic(errorInfo_)
}

func ThrowWithData(text string, code uint64, data interface{}) {
	var errorInfo_ = ErrorInfo{text, text, code, data, nil}
	panic(errorInfo_)
}

func ThrowWithDataInternalMsg(msg string, internalMsg string, code uint64, data interface{}) {
	var errorInfo_ = ErrorInfo{internalMsg, msg, code, data, nil}
	panic(errorInfo_)
}

func ThrowError(text string, code uint64, err error) {
	var errorInfo_ = ErrorInfo{text, text, code, nil, err}
	panic(errorInfo_)
}

func ThrowErrorWithInternalMsg(msg string, internalMsg string, code uint64, err error) {
	var errorInfo_ = ErrorInfo{internalMsg, msg, code, nil, err}
	panic(errorInfo_)
}

func ThrowErrorWithData(text string, code uint64, data interface{}, err error) {
	var errorInfo_ = ErrorInfo{text, text, code, data, err}
	panic(errorInfo_)
}

func ThrowErrorWithDataInternalMsg(msg string, internalMsg string, code uint64, data interface{}, err error) {
	var errorInfo_ = ErrorInfo{internalMsg, msg, code, data, err}
	panic(errorInfo_)
}
