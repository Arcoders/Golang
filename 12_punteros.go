package main

import "fmt"

func main() {

    /*
        1. Es una direccion de memoria
        2. En lugar del valor, tenemos la dirección
        3. Valor zero => nil
    */

    var x, y *int
    entero := 5

    x = &entero
    y = &entero

    *x = 6

    fmt.Println(x)
    fmt.Println(y)

    fmt.Println(*x)
    fmt.Println(*y)

}
