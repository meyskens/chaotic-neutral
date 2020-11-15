package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
)

var destination *url.URL

const rick = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

func main() {
	if os.Getenv("DESTINATION") == "" {
		fmt.Println("You need to set the DESTINATION env var")
		os.Exit(1)
		return
	}

	var err error
	destination, err = url.Parse(os.Getenv("DESTINATION"))
	if err != nil {
		log.Panicln(err)
	}

	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/*", requestHandler)

	log.Println("Starting webserver...")
	err = http.ListenAndServe(":80", nil)
	log.Println(err)
}

// requestHandler does a rickroll on 1% of all requests
func requestHandler(w http.ResponseWriter, r *http.Request) {
	if rn := rand.Intn(100); rn == 42 {
		log.Println("Rick")
		http.Redirect(w, r, rick, http.StatusTemporaryRedirect)
	}

	log.Println("Good")
	// deep copy the url.URL
	dst, _ := url.Parse(destination.String())
	dst.Path = path.Join(destination.Path, r.URL.Path)
	http.Redirect(w, r, dst.String(), http.StatusTemporaryRedirect)
}
