package proxy

import (
	"log"
	"net/http"
)

// server http handler
type Handler interface {
	// return if continue to next handler
	ServeHTTP(w http.ResponseWriter, r *http.Request) bool
}

// print debug log handler
type DebugHandler struct {
	On bool
}

func (this *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	format := "[%s] %s form: %v Connection: %s"
	log.Printf(format, r.Method, r.URL, r.Form, r.Header.Get("Connection"))
	return true
}
