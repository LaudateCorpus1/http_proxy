package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	port        string
	destination string
)

func init() {
	flag.StringVar(&port, "port", "8080", "port to listen to")
	flag.StringVar(&destination, "dest", "", "destination addr")
}

func main() {

	//parse flags
	if !flag.Parsed() {
		flag.Parse()
	}

	url, err := url.Parse(destination)
	if err != nil {
		log.Panicln(err)
	}

	client := httputil.NewSingleHostReverseProxy(url)
	log.Fatal(http.ListenAndServe(":"+port, client))
}
