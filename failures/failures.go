package failures

import (
	"fmt"
	"net/http"
)

// Error can be any error.
type Error struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

var ErrNotFound = Error{
	Code:    http.StatusNotFound,
	Message: http.StatusText(http.StatusNotFound),
}
