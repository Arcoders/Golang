package main

import "fmt"

/*
    == Igual a
    != Diferente de
    < Menor que
    > Mayor que
    >= Mayor o igual que
    <= Menor o igual que
    && AND
    || OR
*/

func main() {
    x := 10
    y := 10

    if x > y {
        fmt.Printf("El numero %d es mayor que %d", x, y)
    }else if x < y {
        fmt.Printf("El numero %d es mayor que %d", y, x)
    }else {
        fmt.Println("Son iguales")
    }
}
