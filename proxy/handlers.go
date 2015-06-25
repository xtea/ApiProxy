package proxy

import (
	"log"
	"net/http"
)

// print debug log handler
type DebugHandler struct {
	On bool
}

func (this *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}
