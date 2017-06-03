package main

import (
    "fmt"
)

func main() {
    nombres := []string{"Ismael", "Haytam", "Victor", "Marta", "Paula"}

    for index, nombre := range nombres {
        fmt.Printf("El nombre %v esta en el index %d \n", nombre, index)
    }

    fmt.Println("----------------------")

    for _, nombre := range nombres {
        fmt.Println(nombre)
    }

    fmt.Println("----------------------")

    cadena := "Arcoders lorem"

    for index, letra := range cadena {
        fmt.Printf("La letra %q esta en el index %d \n", letra, index)
    }
}
