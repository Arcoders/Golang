package main

import (
    "fmt"
)

func print(cadena string) {
    fmt.Println(cadena)
}

func main() {
    nombre := "Ismael Haytam"
    imprimir := print

    imprimir(nombre)

    print := func() {
        fmt.Println(nombre)
    }

    print()
}
