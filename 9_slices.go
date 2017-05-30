package main

import "fmt"

func main() {
    matriz := []int{1, 2, 4, 99, -7}

    if matriz == nil {
        fmt.Println("Está vació")
    }else{
        fmt.Println(len(matriz))
    }

    slice := matriz[:3]
    fmt.Println(slice)

    // Puntero al arreglo
    // Longitud del arreglo al que apunta
    // Capacidad
}
