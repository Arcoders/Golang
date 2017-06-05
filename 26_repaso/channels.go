package main

import (
    "fmt"
    "time"
)

func imprimirPing(c chan string) {
    var contador int
    for {
        contador++
        fmt.Println(<-c, "", contador)
        time.Sleep(time.Second * 1)
    }
}

func enviarPing(c chan string) {
    for {
        c <- "Ping"
    }
}

func main() {
    ch := make(chan string)

    go enviarPing(ch)
    go imprimirPing(ch)

    var input string
    fmt.Scanln(&input)
    fmt.Println("Fin...")
}
