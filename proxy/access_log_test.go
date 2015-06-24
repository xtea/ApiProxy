package proxy

import (
	"github.com/EE-Tools/ApiProxy/proxy"
	"testing"
)

func TestInit(t *testing.T) {
	proxy.InitAccessLogger("")
	//
	proxy.WriteAccessLog("one night in beijing.")
}
