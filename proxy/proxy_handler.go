package proxy

import (
	"net/http"
)

type ProxyHandler struct {
}

func (this *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	//
	return true

}
