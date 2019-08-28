package main

import (
	"fmt"
	"net/http"
	"os"
)

var requestCount int

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr)

	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	requestCount++
	if requestCount >= 5 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Some internal error has occurred! This is pod '%s'\n", host)
		return
	}

	fmt.Fprintf(w, "This is v3 running in pod '%s'\n", host)
}

func main() {
	http.HandleFunc("/", HelloWorld)

	fmt.Println("Kubia server starting ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
