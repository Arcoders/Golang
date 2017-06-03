package main

import (
    "fmt"
)

func main() {
    x := make(map[string]string, 2)
    fmt.Println(x)

    fmt.Println("--------------------------------")

    x["nombre"] = "Ismael"
    x["edad"] = "21"

    fmt.Println(x)

    fmt.Println("Mi nombre: " + x["nombre"])
    fmt.Println("Tengo: " + x["edad"])

    fmt.Println("--------------------------------")

    alfnum := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 6,
    }

    fmt.Println(alfnum["d"])
    alfnum["e"]--
    fmt.Println(alfnum["e"])

    fmt.Println("--------------------------------")

    dia := map[int]string{
        1: "Lunes",
        2: "Martes",
        3: "Miercoles",
        4: "Jueves",
        5: "Abril",
    }

    fmt.Println(dia[2])

    fmt.Println("--------------------------------")

    delete(dia, 5)
    fmt.Println(dia)

    fmt.Println("--------------------------------")

    if dia, ok := dia[3]; ok {
        if dia == "Miercoles" {
            fmt.Println("Correcto el dia 3 es el " + dia)
        } else {
            fmt.Println("Día incorrecto")
        }
    } else {
        fmt.Println("El valor no existe")
    }

    fmt.Println("--------------------------------")

    for numero, dia := range dia {
        fmt.Printf("El día %d es el %s \n", numero, dia)
    }
}
