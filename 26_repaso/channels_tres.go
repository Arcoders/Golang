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
        for {
            x, ok := <- numero; if !ok {
                break
            }
            cuadrado <- x * x
        }
        close(cuadrado)
    }()

    for {
        x, ok := <- cuadrado; if !ok {
            fmt.Println("Fin...")
            break
        }
        fmt.Println(x)
        time.Sleep(1 * time.Second)
    }
}
