package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://www.google.com", 301)
    // or
    // http.Redirect(w, r, "http://www.google.com", http.StatusMovedPermanently);
}
