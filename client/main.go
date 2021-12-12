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

func main() {
	conn, err := net.Dial("tcp", "localhost:3333")
	must(err)

	for {
		messageToServer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		must(err)

		fmt.Fprint(conn, messageToServer)

		replyFromServer, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(replyFromServer)
	}
}
