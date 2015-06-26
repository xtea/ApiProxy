package proxy

import (
	"reflect"
	"testing"
)

func TestParseCodeAndPath(t *testing.T) {

	cases := []struct {
		code, path string
	}{
		{"/hetu", "/project/1"},
		{"/hetu/", "/project/adsfasdf/asdfasfd12/1"},
		{"/hetu/", "/project/adsfasdf/asdfasfd12/1/"},
		{"/s_sdfu/", "/project/adsfasdf/asdfasfd12/1/"},
	}

	for _, c := range cases {
		a, err := ParseCodeAndPath(c.code + c.path)
		if err != nil {
			t.Fatal(err)
		}
		reflect.DeepEqual(a, ApiUrlFileds{c.code, c.path})
	}

}
