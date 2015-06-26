package proxy

import (
	"errors"
	"strings"
)

// sample given: /hetu/product/1
// code :  hetu
// path :  /product/1
type ApiUrlFileds struct {
	Code string
	Path string
}

// Parse url to ApiUrlFileds
func ParseCodeAndPath(urlPath string) (ApiUrlFileds, error) {
	arrays := strings.Split(urlPath, "/")
	if len(arrays) == 0 {
		return ApiUrlFileds{}, errors.New("url doesnot contains '/' ")
	}
	code := arrays[0]
	path := strings.Replace(urlPath, code, "", 1)
	return ApiUrlFileds{code, path}, nil
}
