package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "encoding/json"
    "sync"
)

type Response struct {
    Message string `json:"message"`
    Status bool `json:"status"`
}

var User = struct {
    m map[string] User
}

type User Struct {
    User_Name string
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

function UserExist(user_name string) bool {

}

func validate(w http.ResponseWriter, r *http.Request) {
    r.ParserForm()
    user_name := r.FormValue("user_name")
}

func main()  {

    cssHandle := http.FileServer(http.Dir("./public/css/"))
    jsHandle := http.FileServer(http.Dir("./public/js/"))

    mux := mux.NewRouter()

    mux.HandleFunc("/test", Test).Methods("GET")
    mux.HandleFunc("/testJson", TestJson).Methods("GET")
    mux.HandleFunc("/static", LoadStatic).Methods("GET")
    mux.HandleFunc("/validate", validate).Methods("POST")

    http.Handle("/", mux)
    http.Handle("/css/", http.StripPrefix("/css/", cssHandle))
    http.Handle("/js/", http.StripPrefix("/js/", jsHandle))

    log.Println("El servidor se encuantra en el puerto 8000 ...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
