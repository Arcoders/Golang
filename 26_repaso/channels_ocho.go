package main

import(
    "fmt"
    "sync"
)

const (
    goroutines = 20
    capacidad = 4
)

var wg sync.WaitGroup

var recursos = make(chan string, capacidad)

func main() {
    wg.Add(goroutines)

    for gr := 1; gr <= goroutines; gr++ {
        go func(gr int) {
            worker(gr)
            wg.Done()
        }(gr)
    }

    for s := 'A'; s < 'A'+capacidad; s++ {
        recursos <- string(s)
    }

    wg.Wait()
}

func worker(worker int) {
    valor := <-recursos
    fmt.Printf("worker: %d : %s\n", worker, valor)
    recursos <- valor
}
