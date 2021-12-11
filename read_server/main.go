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
	defer func(conn net.Conn) {
		err := conn.Close()
		must(err)
	}(conn)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3333")
	must(err)

	defer func(l net.Listener) {
		err := l.Close()
		must(err)
	}(listener)

	for {
		conn, err := listener.Accept()
		must(err)

		handle(conn)
	}
}