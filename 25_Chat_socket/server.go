package main

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/WebSocket"
    "log"
    "net/http"
    "encoding/json"
    "sync"
)

type Response struct {
    Message string `json:"message"`
    Status int `json:"status"`
    IsValid bool `json:"isvalid"`
}

var Users = struct {
    m map[string] User
    sync.RWMutex
}{m: make(map[string] User)}

type User struct {
    User_Name string
    WebSocket *websocket.Conn
}

func Test(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte ("<h1>Test Go Server</h1>"))
}

func TestJson(w http.ResponseWriter, r *http.Request) {
    response := CreateResponse("Test del formato JSON",200, true)
    json.NewEncoder(w).Encode(response)
}

func CreateResponse(message string, status int, valid bool) Response {
    return Response{message, status, valid}
}

func LoadStatic(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/index.html")
}

func UserExist(user_name string) bool {
    Users.RLock()
    defer Users.RUnlock()

    if _, ok := Users.m[user_name]; ok {
        return true
    }

    return false
}

func Validate(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    user_name := r.FormValue("user_name")

    response := Response{}

    if UserExist(user_name) {
        response.IsValid = false
    }else{
        response.IsValid = true
    }

    json.NewEncoder(w).Encode(response)
}

func AddUser(user User) {
    Users.Lock()
    defer Users.Unlock()

    Users.m[user.User_Name] = user
}

func CreateUser(user_name string, ws *websocket.Conn) User {
    return User{user_name, ws}
}

func RemoveUser(user_name string) {
    Users.Lock()
    defer Users.Unlock()
    delete(Users.m, user_name)
    log.Println("El usuario " + user_name + " se desconectó")
}

func SendMessage(type_message int, message []byte) {
    Users.RLock()
    defer Users.RUnlock()

    for _, user := range Users.m{
        if err := user.WebSocket.WriteMessage(type_message, message); err != nil {
            return
        }
    }
}

func ToArrayByte(value string) []byte {
    return []byte(value)
}

func ConcatMessage(user_name string, array []byte) string {
    return user_name + " : " + string(array[:])
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user_name := vars["user_name"]

    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

    if err != nil {
        log.Println(err)
        return
    }

    current_user := CreateUser(user_name, ws);
    AddUser(current_user)
    log.Println("Nuevo usuario agregado con el nombre: " + user_name)

    for {
        type_message, message, err := ws.ReadMessage();
        if err != nil {
            RemoveUser(user_name)
            return
        }
        final_message := ConcatMessage(user_name, message)
        SendMessage(type_message, ToArrayByte(final_message))
    }
}

func main()  {
    cssHandle := http.FileServer(http.Dir("./public/css/"))
    jsHandle := http.FileServer(http.Dir("./public/js/"))

    mux := mux.NewRouter()

    mux.HandleFunc("/test", Test).Methods("GET")
    mux.HandleFunc("/testJson", TestJson).Methods("GET")
    mux.HandleFunc("/static", LoadStatic).Methods("GET")
    mux.HandleFunc("/validate", Validate).Methods("POST")
    mux.HandleFunc("/chat/{user_name}", WebSocket).Methods("GET")

    http.Handle("/", mux)
    http.Handle("/css/", http.StripPrefix("/css/", cssHandle))
    http.Handle("/js/", http.StripPrefix("/js/", jsHandle))

    log.Println("El servidor se encuantra en el puerto 8000 ...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
