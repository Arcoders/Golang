package main

import (
    "net/http"
    "io"
    "fmt"
)

func main() {
    http.HandleFunc("/", handler)

    http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hay una nueva petición...")
    io.WriteString(w, "<h1>Arcoder...<h1>")
}
