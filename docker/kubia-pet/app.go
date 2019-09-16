package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const dataFile = "/var/data/kubia.txt"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr)

	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(dataFile, body, 0644); err != nil {
			panic(err)
		}

		fmt.Println("New data has been received and stored.")
		fmt.Fprintln(w, "Data stored on pod "+host)
	} else {
		data, err := ioutil.ReadFile(dataFile)
		if err != nil {
			data = []byte("No data posted yet")
		}

		fmt.Fprintf(w, "You've hit '%s'\n", host)
		fmt.Fprintf(w, "Data stored on this pod: %s\n", data)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", HelloWorld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
