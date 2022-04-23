package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("http server starting ...")
	os.Setenv("VERSION", "1.0.0")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/localhost/healthz", healthzHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	requestHeader := r.Header
	for k, vs := range requestHeader {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
	log.Println(r.RemoteAddr)
}
