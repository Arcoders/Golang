package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8009")
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan struct{})

    go func() {
        io.Copy(os.Stdout, conn)
        log.Println("Terminamos")
        done <- struct{}{}
    }()

    mustCopy(Conn, os.Stdin)
    conn.Close()
    <-done
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err!= nil {
        log.Fatal(err)
    }
}
