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
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3333")
	must(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		must(err)

		handle(conn)
	}
}