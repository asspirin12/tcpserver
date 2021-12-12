package main

import (
	"io"
	"os"
	"fmt"
	"net"
)

func must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func handle(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	must(err)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3333")
	must(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		must(err)

		defer conn.Close()

		handle(conn)
	}
}