package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Iniciar Goroutines")

    go func() {
        for count := 100; count >= 0; count-- {
            fmt.Printf("[A:%d] \n", count)
        }
        wg.Done()
    }()

    go func() {
        for count := 0; count <= 100; count++ {
            fmt.Printf("[B:%d] \n", count)
        }
        wg.Done()
    }()

    fmt.Println("Esperando a que terminen")
    wg.Wait()

    fmt.Println("Finalizar el programa...")
}
