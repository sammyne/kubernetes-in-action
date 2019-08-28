package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr)

	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "This is v4 running in pod '%s'\n", host)
}

func main() {
	http.HandleFunc("/", HelloWorld)

	fmt.Println("Kubia server starting ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
