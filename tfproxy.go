package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kuberlab/tfservable-proxy/pkg/tfhttp"
)

var (
	prefix  = "proxy"
	timeout = 300
	port    int
)

func main() {
	flag.IntVar(&port, "port", 8082, "Proxy port")
	flag.IntVar(&timeout, "timeout", 300, "Timeout for model call in sec")

	proxy := tfhttp.TFHttpProxy{
		URIPrefix: prefix,
		Timeout:   time.Duration(timeout) * time.Second,
	}

	flag.StringVar(&proxy.DefaultAddress, "default-addr", "", "Default target address if applicable")
	flag.IntVar(&proxy.DefaultPort, "default-port", 9000, "Default target server port")
	flag.Parse()

	if proxy.DefaultAddress != "" {
		log.Printf("Default target address: %v", proxy.DefaultAddress)
	}
	log.Printf("Default target port: %v", proxy.DefaultPort)

	http.Handle("/", proxy)
	log.Printf("Listen on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
