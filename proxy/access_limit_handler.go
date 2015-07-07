package proxy

import (
	//"github.com/EE-Tools/goauth/ext/frequency"
	"log"
	"net/http"
)

// Access limit handler
type AccessLimitHandler struct {
	On bool
}

func (a *AccessLimitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	if a.On {
		// get api info.
		apiInfo, err := ParseApiInfo(r.URL)
		if err != nil {
			// parse error
			log.Printf("parse api fail, write 404 error,%s", err)
			WriteHttpError(HTTP_CODE_NOTFOUND, w)
			return false
		}
		log.Println(apiInfo)
		// call check limit entry method.
		// func CheckAccessLimit(entry *AccessLimitEntry) (bool, error)
	}
	return true
}
