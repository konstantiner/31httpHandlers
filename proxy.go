package main

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
)

func main() {
	firstHost, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	secondHost, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	counter := 0
	
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{})
	proxy.Director = func(req *http.Request) {
		if counter <= 0 {
			counter++
			fmt.Println("8080")
			req.URL.Scheme = firstHost.Scheme
			req.URL.Host = firstHost.Host
		} else {
			counter--
			fmt.Println("8081")
			req.URL.Scheme = secondHost.Scheme
			req.URL.Host = secondHost.Host
		}
	}

	log.Fatal(http.ListenAndServe("localhost:8877", proxy))
}