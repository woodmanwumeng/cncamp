package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http server starting ...")
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	requestHeader := r.Header
	for k, vs := range requestHeader {
		fmt.Printf("head key : %s ; header value: %s\n", k, vs)
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
}
