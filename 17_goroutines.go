package main

import (
    "fmt"
    "strings"
)

func main() {
    mi_nombre_lento("Ismael Haytam")
}

func mi_nombre_lento(name string)  {
    letras := strings.Split(name, "")

    for _,letra := range(letras) {
        fmt.Println(letra)
    }
}
