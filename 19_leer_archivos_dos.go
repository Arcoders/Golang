package main

import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    archivo, error := os.Open("./19_exm.txt")

    if error != nil {
        fmt.Println("Hubo un error")
    }

    scanner := bufio.NewScanner(archivo)

    var i int

    for scanner.Scan() {
        i++
        linea := scanner.Text()
        fmt.Println(i)
        fmt.Println(linea)
    }
}
