package proxy

import (
	"errors"
	"fmt"
	"strings"
)

type ApiUrlFileds struct {
	Code string
	Path string
}

// Parse url to ApiUrlFileds
// given:  /hetu/product/1
// return: code :  hetu , path :  /product/1
func ParseCodeAndPath(urlPath string) (ApiUrlFileds, error) {
	arrays := strings.Split(urlPath, "/")
	if len(arrays) == 0 {
		return ApiUrlFileds{}, errors.New("url doesnot contains '/' ")
	}
	code := ""
	// get first not empty string.
	for _, str := range arrays {
		if str != "" {
			code = str
			break
		}
	}
	if code == "" {
		return ApiUrlFileds{}, errors.New("code is empty from " + urlPath)
	}
	// substring start with end of code
	pathBegin := strings.Index(urlPath, code) + len(code)
	path := urlPath[pathBegin:]
	// replace '//' with '/'
	path = strings.Replace(path, "//", "/", -1)
	if path == "" {
		return ApiUrlFileds{}, errors.New("path code is empty from " + urlPath)
	}
	return ApiUrlFileds{code, path}, nil
}

func (a *ApiUrlFileds) String() string {
	return fmt.Sprintf("code:%s,path:%s", a.Code, a.Path)
}
