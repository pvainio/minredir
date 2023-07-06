package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func getRedirect(redirect *url.URL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		location := redirect.JoinPath(r.URL.Path)
		w.Header().Add("Location", location.String())
		w.WriteHeader(http.StatusMovedPermanently)
	}
}

func main() {
	redirect := os.Getenv("REDIRECT")

	if len(redirect) == 0 {
		fmt.Print("REDIRECT environment not set")
		os.Exit(1)
	}

	redirectUrl, err := url.Parse(redirect)
	if err != nil {
		fmt.Printf("error parsing REDIRECT url: %s\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", getRedirect(redirectUrl))

	err = http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
