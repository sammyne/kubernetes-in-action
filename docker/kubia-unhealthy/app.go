package main

import (
	"fmt"
	"net/http"
	"os"
)

var throttle int

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr)

	if throttle >= 3 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "failed to serve request")
		return
	}

	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "You've hit '%s'\n", host)
	throttle++
}

func main() {
	http.HandleFunc("/", HelloWorld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
