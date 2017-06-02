package main

import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    fmt.Println(readFile())
}

func readFile() bool {
    archivo, error := os.Open("./error.txt")

    if error != nil {
        panic(error)
    }

    scanner := bufio.NewScanner(archivo)

    var i int

    for scanner.Scan() {
        i++
        linea := scanner.Text()
        fmt.Println(i)
        fmt.Println(linea)
    }

    defer func() {
        archivo.Close()
        fmt.Println("Cerrar")
        recover()
    }()

    return true
}
