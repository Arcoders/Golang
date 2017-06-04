package main

import (
    "fmt"
    "os"
)

func primera() {
    fmt.Println("Primera")
}

func segunda() {
    fmt.Println("Segunda")
}

func main() {
    defer primera()
    segunda()

    f, err := os.Open("file.txt")

    if err != nil {
        panic(err)
    }

    defer f.Close()

    data := make([]byte, 15)
    c, err := f.Read(data)

    if err != nil {
        panic(err)
    }

    fmt.Printf("Cantidad de byte leidos : %d \n Texto leido: %s Error: %v \n", c, data, err)

    for i := 0; i <= 5; i++ {
        defer fmt.Println(i)
    }
}
