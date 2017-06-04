package main

import (
    "fmt"
    "strings"
)

func print(cadena string) {
    fmt.Println(cadena)
}

func printDos(cadena string) {
    fmt.Println(cadena)
}

func printTres(cadenaUno, cadenaDos string) {
    fmt.Println(cadenaUno + cadenaDos)
}

func printCuatro(fprint func(string)) {
    fprint("Hola desde printCuatro")
}

func incrementar() func() int {
    i := 0
    return func() (r int) {
        r = i
        i += 2
        return
    }
}

func multiplicar(numero int) (n1, n2, n3 int) {
    n1 = numero * 10
    n2 = numero * 20
    n3 = numero * 30
    return
}

func main() {
    nombre := "Ismael Haytam"
    imprimir := print

    imprimir(nombre)

    imprimirDos := func() {
        fmt.Println(nombre)
    }

    imprimirDos()

    imprimir = printDos
    imprimir("Ismael Haytam")

    fmt.Printf("Function print: %T \n", print)
    fmt.Printf("Function printDos: %T \n", printDos)
    fmt.Printf("Function printTres: %T \n", printTres)

    printCuatro(print)

    var nada func()
    if nada == nil {
        fmt.Println("La func Nada es igual a nil")
    }

    inc := incrementar()

    fmt.Println("Valor de i: ", inc())
    fmt.Println("Valor de i: ", inc())
    fmt.Println("Valor de i: ", inc())
    fmt.Println("Valor de i: ", inc())
    fmt.Println("Valor de i: ", inc())

    cadena := "abc"

    cadena = strings.Map(func(r rune) rune {
        return r + 3
    }, cadena)

    fmt.Println(cadena)

    fmt.Println(multiplicar(10))
}
