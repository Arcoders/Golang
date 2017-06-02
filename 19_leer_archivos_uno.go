package main

import(
    "io/ioutil"
    "fmt"
)

func main() {
    file_data, err := ioutil.ReadFile("./19_exm.txt")

    if err != nil {
        fmt.Println("Hubo un error")
    }

    fmt.Println(string(file_data))
}
