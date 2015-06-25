package proxy

import (
	_ "log"
	"net/http"
)

// oauth check handler
type OauthCheckHandler struct {
}

func (this *OauthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	return true
}
