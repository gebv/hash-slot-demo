package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

var seq = 0

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	seq++
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, World %d!\n", seq)
	fmt.Fprintln(w, "Headers:")
	for k := range r.Header {
		v := r.Header.Get(k)
		fmt.Fprintf(w, " %s=%q\n", k, v)
	}

	res, _ := http.Get("http://metadata.google.internal")
	defer res.Body.Close()
	resBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintln(w, "Metadata:", string(resBytes))
}
