package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Subhana rabiya l3adim")
    }

    i := 0

    for {
        if i == 2 {
            i++
            continue
        }

        fmt.Println(i)
        i++

        if i > 10 {
            break
        }
    }
}
