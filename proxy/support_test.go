package proxy

import (
	"github.com/EE-Tools/ApiProxy/proxy"
	"strings"
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
		a, err := proxy.ParseCodeAndPath(c.code + c.path)
		if err != nil {
			t.Fatal(err)
		}
		expectC := strings.Replace(c.code, "/", "", -1)
		if a.Code != expectC {
			t.Errorf("expect %s,but got %s", expectC, a.Code)
		}
		if a.Path != c.path {
			t.Errorf("expect %s,but got %s", c.path, a.Path)
		}
	}

}
