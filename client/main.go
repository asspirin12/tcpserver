package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func writeRead(conn net.Conn) {
	for {
		messageToServer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		must(err)

		fmt.Fprint(conn, messageToServer)

		replyFromServer, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(replyFromServer)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:3333")
	must(err)

	defer conn.Close()

	writeRead(conn)
}