package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main() {
    wg.Add(2)
    fmt.Println("Iniciamos las gorutinas...")

    go imprimirCantidad("A")
    go imprimirCantidad("B")

    fmt.Println("Esperando que finalicen...")
    wg.Wait()
    fmt.Println("Terminando el programa")
}

func imprimirCantidad(etiqueta string) {
    defer wg.Done()

    for cantidad := 1; cantidad <= 10; cantidad++ {
        time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
        fmt.Printf("Cantidad: %d de %s \n", cantidad, etiqueta)
    }
}
