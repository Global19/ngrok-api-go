package ngrok

import (
	"fmt"
	"net/http"
)

// Returns true if the error is a not found response from the ngrok API.
func IsNotFound(err error) bool {
	if ee, ok := err.(*Error); ok {
		return int(ee.StatusCode) == http.StatusNotFound
	}
	return false
}

// Returns true if the given error is caused by any of the specified ngrok error codes.
func IsErrorCode(err error, codes ...int) bool {
	if ee, ok := err.(*Error); ok {
		for _, code := range codes {
			if ee.ErrorCode == fmt.Sprintf("ERR_NGROK_%d", code) {
				return true
			}
		}
	}
	return false
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Msg)
	if e.ErrorCode != "" {
		msg += fmt.Sprintf(" [%s]", e.ErrorCode)
	}
	if e.operationID() != "" {
		msg += fmt.Sprintf("\n\nOperation ID: %s", e.operationID())
	}
	return msg
}

func (e *Error) operationID() string {
	return e.Details["operation_id"]
}