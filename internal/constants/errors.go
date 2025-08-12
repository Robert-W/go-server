package constants

import "net/http"

const (
	ERROR_FORBIDDEN int = iota
	ERROR_INTERNAL
	ERROR_NO_CONTENT
)

func ErrorCode(error int) int {
	switch error {
	case ERROR_FORBIDDEN:
		return http.StatusForbidden
	case ERROR_INTERNAL:
		return http.StatusInternalServerError
	case ERROR_NO_CONTENT:
		return http.StatusNoContent
	default:
		return http.StatusInternalServerError
	}
}
