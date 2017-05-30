package main

import "fmt"

func main()  {
    arreglo := [10] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    for i := 0; i < len(arreglo); i++ {
        fmt.Println(arreglo[i])
    }

    var matriz [3][2] int
    matriz[0][1] = 1

    fmt.Println(matriz)
}
