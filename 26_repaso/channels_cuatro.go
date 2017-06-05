package main

import(
    "fmt"
    "time"
)

func main() {
    numero := make(chan int)
    cuadrado := make(chan int)

    go func() {
        for x := 0; x < 5; x++ {
            numero <- x
        }
        close(numero)
    }()

    go func() {
        for x := range numero {
            cuadrado <- x * x
        }
        close(cuadrado)
    }()

    for x := range cuadrado {
        fmt.Println(x)
        time.Sleep(1 * time.Second)
    }
}
