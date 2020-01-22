package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kuberlab/tfservable-proxy/pkg/tfhttp"
)

var (
	timeout        = 300
	port           int
	URIPrefix      string
	DefaultPort    int
	DefaultAddress string
	staticRoot     string
)

func main() {
	flag.IntVar(&port, "port", 8082, "Proxy port")
	flag.IntVar(&timeout, "timeout", 900, "Timeout for model call in sec")

	flag.StringVar(&DefaultAddress, "default-addr", "", "Default target address if applicable")
	flag.IntVar(&DefaultPort, "default-port", 9000, "Default target server port")
	flag.StringVar(&URIPrefix, "uri-prefix", "proxy", "URI path for proxy")
	flag.StringVar(&staticRoot, "static-root", "./static", "Path for static content")
	flag.Parse()

	proxy := tfhttp.NewProxy("/", URIPrefix, staticRoot)
	proxy.DefaultAddress = DefaultAddress
	proxy.DefaultPort = DefaultPort
	proxy.Timeout = time.Duration(timeout) * time.Second

	if proxy.DefaultAddress != "" {
		log.Printf("Default target address: %v", proxy.DefaultAddress)
	}
	log.Printf("Default target port: %v", proxy.DefaultPort)
	log.Printf("Proxy will be available on: http://0.0.0.0:%v%v", port, proxy.URIPrefix)
	log.Printf("Proxy will be available on: http://0.0.0.0:%v/serving-proxy", port)

	log.Printf("Listen on :%d\n", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Printf("Shutdown proxy...")
			os.Exit(0)
		}
	}()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), proxy))
}
