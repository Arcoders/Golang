package main

import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Ingresa tu nombre: ")
    nombre, err := reader.ReadString('\n')

    if (err != nil) {
        fmt.Println(err)
    }else{
        fmt.Println("Hola " + nombre)
    }
}
