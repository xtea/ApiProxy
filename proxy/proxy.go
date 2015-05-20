package main

import (
	"errors"
	"fmt"
	"github.com/xtea/app"
	"io"
	"log"
	"net/http"
	"net/url"
)

// entry point.
func main() {
	// add handle
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// parse api info from URL.
		api, err := ParseApiInfo(r.URL)
		if err != nil {
			ErrorMessage(w, err.Error())
			return
		}
		// save log for statistic.
		log.Println("access api", api)
		// find mapping mehtod by app.ApiInfo
		handleApiMethod, err := findHandleMethod(api)
		if err != nil {
			ErrorMessage(w, "api method not support.")
			return
		}
		// call real handle function
		handleApiMethod(api, w, r)
	})
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// Parse api info from url.URL.
func ParseApiInfo(u *url.URL) (app.ApiInfo, error) {
	api, ok := app.GetApiInfoById(u.Path)
	if !ok {
		return app.ApiInfo{}, errors.New("api not found.")
	}
	return api, nil
}

// Response error message to client.
func ErrorMessage(w http.ResponseWriter, format string, a ...interface{}) {
	if len(a) > 0 {
		fmt.Fprintf(w, format, a)
	} else {
		fmt.Fprintf(w, format)
	}

}

// Find support handle method by ApiInfo.
func findHandleMethod(a app.ApiInfo) (HandleMethod, error) {
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

// define delegate handle method.
type HandleMethod func(a app.ApiInfo, w http.ResponseWriter, r *http.Request)

// Hanle http get method to remote api.
func HandleGetMethod(a app.ApiInfo, w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(a.Url + "?" + r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(w, "error is %q", err)
		return
	}
	afterHandleMethod(w, resp)
}

// For sending http post method to remote api.
func HandlePostMethod(a app.ApiInfo, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := http.PostForm(a.Url, r.PostForm)
	if err != nil {
		fmt.Fprintf(w, "error is %q", err)
		return
	}
	afterHandleMethod(w, resp)
}

// After custom handle method.
func afterHandleMethod(w http.ResponseWriter, resp *http.Response) {
	defer resp.Body.Close()
	_, err := io.Copy(w, resp.Body)
	if err != nil {
		log.Println("copy error %q", err)
	}
}
