package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

const (
	dataFile    = "/var/data/kubia.txt"
	serviceName = "kubia.default.svc.cluster.local"
	port        = 8080
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr)

	host, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "failed to get hostname: %v\n", err)
		return
	}

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "failed to read request body: %v\n", err)
			return
		}

		if err := ioutil.WriteFile(dataFile, body, 0644); err != nil {
			fmt.Fprintf(w, "failed to save request body: %v\n", err)
			return
		}

		fmt.Println("New data has been received and stored.")
		fmt.Fprintln(w, "Data stored on pod "+host)
		return
	}

	if r.URL.Path == "/data" {
		fmt.Fprintf(w, "You've hit '%s'\n", host)

		data, err := ioutil.ReadFile(dataFile)
		if err == nil {
			fmt.Fprintf(w, "Data stored on this pod: %s\n", data)
		} else {
			fmt.Fprintln(w, "Data stored on this pod: No data posted yet", data)
		}

		return
	}

	fmt.Fprintln(w, "You've hit "+host)
	fmt.Fprintln(w, "Data stored in the cluster:")

	_, addresses, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		fmt.Fprintf(w, "Could not look up DNS SRV records: %v\n", err)
		return
	}

	if len(addresses) == 0 {
		fmt.Fprintln(w, "No peers discovered")
		return
	}

	for _, addr := range addresses {
		url := fmt.Sprintf("http://%s:%d/data", addr.Target, port)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(w, "failed to POST to %s: %v\n", url, err)
			return
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "err reading response from %s\n", url)
			resp.Body.Close()
			return
		}
		fmt.Fprintf(w, "- %s: %s\n", addr.Target, data)
		resp.Body.Close()
	}

	//w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", HelloWorld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
