package main

import (
    "fmt"
    "net"
)

func must(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    listener, err := net.Listen("tcp", "localhost:3333")
    must(err)

    defer listener.Close()

    for {
        conn, err := listener.Accept()
        must(err)

        defer conn.Close()

        _, err = conn.Write([]byte("Hello world"))
        must(err)
    }
}