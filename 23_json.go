package main

import(
    "net/http"
    "encoding/json"
)

type Usuario struct {
    Nombre string `json:"title"`
    Edad int
}

type Usuarios []Usuario

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    usuarios := Usuarios{
        Usuario{"Ismael Haytam", 21},
        Usuario{"Fatima Chadli", 45},
        Usuario{"Hassan Tanan", 45},
    }

    json.NewEncoder(w).Encode(usuarios)
})

    http.ListenAndServe(":8000", nil)
}
