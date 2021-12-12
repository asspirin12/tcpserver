package main

import (
	"bufio"
	"fmt"
	"net"
)

func must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func handle(conn net.Conn) {
	for {
		messageFromClient, err := bufio.NewReader(conn).ReadString('\n')
		must(err)

		fmt.Print(string(messageFromClient))

		fmt.Fprintf(conn, "Server received a message: %s", messageFromClient)
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

		handle(conn)
	}
}
