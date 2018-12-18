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
	timeout = 300
	port    int
)

func main() {
	flag.IntVar(&port, "port", 8082, "Proxy port")
	flag.IntVar(&timeout, "timeout", 300, "Timeout for model call in sec")

	proxy := tfhttp.Proxy{
		Timeout: time.Duration(timeout) * time.Second,
	}

	flag.StringVar(&proxy.DefaultAddress, "default-addr", "", "Default target address if applicable")
	flag.IntVar(&proxy.DefaultPort, "default-port", 9000, "Default target server port")
	flag.StringVar(&proxy.URIPrefix, "uri-prefix", "proxy", "URI path for proxy")
	flag.Parse()

	if proxy.DefaultAddress != "" {
		log.Printf("Default target address: %v", proxy.DefaultAddress)
	}
	log.Printf("Default target port: %v", proxy.DefaultPort)
	log.Printf("Proxy will be available on: http://0.0.0.0:%v/%v", port, proxy.URIPrefix)

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
