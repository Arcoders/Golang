package main

import (
    "fmt"
)

func imprimir() {
    fmt.Println("Ismael Haytam")

    defer func() {
        cadena := recover()
        fmt.Println(cadena)
    }()
    panic("Error detected... line 14")
}

func main() {
    imprimir()
    fmt.Println("Arcoders")
}
