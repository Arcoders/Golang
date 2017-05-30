package main

import(
    "fmt"
    "strconv"
)

func main() {
    edad := "22"

    edad_int,_ := strconv.Atoi(edad)

    fmt.Println(10 + edad_int)
}
