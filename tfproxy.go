package main

import (
	"flag"
	"fmt"
	"github.com/kuberlab/tfservable-proxy/pkg/tfhttp"
	"log"
	"net/http"
	"time"
)

var (
	prefix  = "proxy"
	timeout = 300
	port    int
)

func main() {
	flag.IntVar(&port, "port", 8082, "Proxy port")
	flag.IntVar(&timeout, "timeout", 300, "Timeout for model call in sec")
	flag.Parse()
	proxy := tfhttp.TFHttpProxy{
		URIPrefix: prefix,
		Timeout:   time.Duration(timeout) * time.Second,
	}
	http.Handle("/", proxy)
	log.Printf("Listen on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
