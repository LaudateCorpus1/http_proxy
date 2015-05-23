package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/bountylabs/log"
	"github.com/tonnerre/golang-pretty"
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

	if destination == "" {
		log.Panicln("requires destination")
	}

	url, err := url.Parse(destination)
	if err != nil {
		log.Panicln(err)
	}
	log.Infoln("Redirecting to", url)

	client := httputil.NewSingleHostReverseProxy(url)
	orig := client.Director
	client.Director = func(req *http.Request) {
		orig(req)
		log.Infoln(req.URL)
		pretty.Log(req.Header)
	}

	if err := http.ListenAndServe(":"+port, client); err != nil {
		log.Panicln(err)
	}
}
