package proxy

import (
	"errors"
	"fmt"
	"github.com/EE-Tools/regapi/models"
	"io"
	"log"
	"net/http"
	_ "strings"
)

// define delegate handle method.
type HandleMethod func(a models.ApiInfo, w http.ResponseWriter, r *http.Request)

type ProxyHandler struct {
	Log string
}

func (this *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	//
	apiInfo, err := ParseApiInfo(r.URL)
	if err != nil {
		// parse error
		log.Printf("parse api fail, write 404 error,%s", err)
		WriteHttpError(HTTP_CODE_NOTFOUND, w)
		return false
	}
	// find adapte http method function
	httpMethodHandle, err := findHandleMethod(apiInfo)
	if err != nil {
		log.Printf("register api %v method not support", apiInfo)
		WriteHttpError(HTTP_CODE_INTERNAL_SERVER_ERROR, w)
		return false
	}
	// call handle proxy
	httpMethodHandle(apiInfo, w, r)

	// save access log
	WriteAccessLogByApiInfo(
		LogWrapper{
			AppId:    apiInfo.ApiId,
			ClientIp: r.RemoteAddr,
		})
	return true
}

// Find support handle method by ApiInfo.
func findHandleMethod(a models.ApiInfo) (HandleMethod, error) {
	switch m := a.Method; m {
	case "GET":
		// GET method.
		return HandleGetMethod, nil
	case "POST":
		// POST method
		return HandlePostMethod, nil
	default:
		return nil, errors.New("http method not support.")
	}
}

// Hanle http get method to remote api.
func HandleGetMethod(a models.ApiInfo, w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(buildRemoteApiUrl(a) + "?" + r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(w, "error is %q", err)
		return
	}
	afterHandleMethod(w, resp)
}

// // For sending http post method to remote api.
func HandlePostMethod(a models.ApiInfo, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := http.PostForm(buildRemoteApiUrl(a), r.PostForm)
	if err != nil {
		fmt.Fprintf(w, "error is %q", err)
		return
	}
	afterHandleMethod(w, resp)
}

// build remote api url
func buildRemoteApiUrl(a models.ApiInfo) string {
	r := a.MainUrl + a.Path
	log.Printf("access remote api %s", r)
	return r
}

// After common opeations.
func afterHandleMethod(w http.ResponseWriter, resp *http.Response) {
	// if status is not ok , direct write error to client.
	if resp.StatusCode != 200 {
		WriteHttpError(resp.StatusCode, w)
		return
	}
	defer resp.Body.Close()
	_, err := io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("copy io err %s", err)
	}
}
