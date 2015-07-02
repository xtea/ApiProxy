package proxy

import (
	"fmt"
	"net/http"
)

const (
	HTTP_CODE_NOTFOUND              = 404
	HTTP_CODE_INTERNAL_SERVER_ERROR = 500
	HTTP_CODE_FORBIDDEN             = 403
)

// write http error message to client.
func WriteHttpError(code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	fmt.Fprintf(w, http.StatusText(code))
}

// write http error message to client.
func WriteHttpErrorMessage(code int, message string, w http.ResponseWriter) {
	w.WriteHeader(code)
	fmt.Fprintf(w, message)
}
