package proxy

import (
	"testing"
)

func TestInit(t *testing.T) {
	InitAccessLogger("")
	//
	WriteAccessLog("one night in beijing.")
}
