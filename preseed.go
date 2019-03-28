package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var config string

func main() {
	// Handle first argument
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: ./preseed <configfile>")
	}

	// Read config file into variable
	filename := args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	config = string(content)

	// Launch webserver
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Output formatted config
	fmt.Fprintf(w, editConfig(r.URL.Path[1:]))
}

func editConfig(hostname string) string {
	// Replace the hostname on config file based on url path
	return strings.Replace(config, "forced-hostname", hostname, 1)
}
