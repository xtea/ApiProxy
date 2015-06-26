package main

import (
	"flag"
	"github.com/EE-Tools/ApiProxy/proxy"
	"log"
	"net/http"
	"time"
)

// listen http port
var port = flag.String("port", "8080", "http serve port")
var mode = flag.String("mode", "debug", "project run mode,default is debug")

func init() {
	flag.Parse()
}

// entry point.
func main() {
	// startup info.
	log.Println("run mode is", *mode)
	log.Println("startup and listen", *port)

	hlist := []proxy.Handler{
		// print debug log handler
		&proxy.DebugHandler{},
	}

	// add handle
	s := &http.Server{
		Addr: ":" + *port,
		Handler: &DefaultHandleChain{
			HandleList: hlist,
		},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type DefaultHandleChain struct {
	HandleList []proxy.Handler
}

func (this *DefaultHandleChain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handle := range this.HandleList {
		handle.ServeHTTP(w, r)
	}
}
