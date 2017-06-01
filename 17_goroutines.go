package main

import (
    "fmt"
    "time"
    "strings"
)

func main() {
    go mi_nombre_lento("Ismael Haytam")
    fmt.Println("Que Aburrimiento...")

    var wait string
    fmt.Scanln(&wait)
}

func mi_nombre_lento(name string)  {
    letras := strings.Split(name, "")

    for _,letra := range(letras) {
        time.Sleep(500 * time.Millisecond)
        fmt.Println(letra)
    }
}
