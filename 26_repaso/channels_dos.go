package main

import(
    "fmt"
    "time"
)

func main() {
    numero := make(chan int)
    cuadrado := make(chan int)

    go func() {
        for x := 0; ; x++ {
            numero <- x
        }
    }()

    go func() {
        for {
            x := <- numero
            cuadrado <- x * x
        }
    }()

    for {
        fmt.Println(<-cuadrado)
        time.Sleep(1 * time.Second)
    }
}
