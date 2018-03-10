package jsend

// Status constants
const (
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "error"
)

// JSend type signature
type JSend struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// Success JSend response
func Success(data interface{}) *JSend {
	return &JSend{Status: StatusSuccess, Data: data}
}

// Fail JSend response
func Fail(data interface{}) *JSend {
	return &JSend{Status: StatusFail, Data: data}
}

// Error JSend response
func Error(message string) *JSend {
	return &JSend{Status: StatusError, Message: message}
}

// ErrorCode JSend with error code
func ErrorCode(message string, code int) *JSend {
	return &JSend{Status: StatusError, Message: message, Code: code}
}

// ErrorData JSend with error data
func ErrorData(message string, data interface{}) *JSend {
	return &JSend{Status: StatusError, Message: message, Data: data}
}

// ErrorCodeWithData JSend with error code and data
func ErrorCodeWithData(message string, code int, data interface{}) *JSend {
	return &JSend{Status: StatusError, Message: message, Code: code, Data: data}
}
