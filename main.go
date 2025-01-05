package main

import (
	"github.com/hellolib/apifox-proxy/proxy"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", proxy.ReverseProxy()))
}
