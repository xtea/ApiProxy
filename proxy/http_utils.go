package proxy

import (
	"fmt"
	"net/http"
)

const (
	HTTP_CODE_NOTFOUND              = 404
	HTTP_CODE_INTERNAL_SERVER_ERROR = 500
)

// write http error message to client.
func WriteHttpError(code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	fmt.Fprintf(w, http.StatusText(code))
}
