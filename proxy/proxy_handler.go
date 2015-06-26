package proxy

import (
	_ "errors"
	"github.com/EE-Tools/regapi/models"
	"net/http"
	"net/url"
	_ "strings"
)

// define delegate handle method.
//type HandleMethod func(a models.ApiInfo, w http.ResponseWriter, r *http.Request)

type ProxyHandler struct {
}

func (this *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	//

	return true
}

// // Parse api info from url.URL.
func ParseApiInfo(u *url.URL) (models.ApiInfo, error) {

	return models.ApiInfo{}, nil
}
