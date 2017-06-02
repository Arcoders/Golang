package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "encoding/json"
)

type Response struct {
    Message string `json:"message"`
    Status bool `json:"status"`
}

func Test(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte ("<h1>Test Go Server</h1>"))
}

func TestJson(w http.ResponseWriter, r *http.Request) {
    response := CreateResponse()
    json.NewEncoder(w).Encode(response)
}

func CreateResponse() Response {
    return Response{"Test del formato JSON", true}
}

func LoadStatic(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/index.html")
}

func main()  {

    cssHandle := http.FileServer(http.Dir("./public/css/"))

    mux := mux.NewRouter()

    mux.HandleFunc("/test", Test).Methods("GET")
    mux.HandleFunc("/testJson", TestJson).Methods("GET")
    mux.HandleFunc("/static", LoadStatic).Methods("GET")

    http.Handle("/", mux)
    http.Handle("/css/", http.StripPrefix("/css/", cssHandle))

    log.Println("El servidor se encuantra en el puerto 8000 ...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
