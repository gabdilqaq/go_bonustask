package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	text := bufio.NewReader(conn).ReadString('\n')
	msg := fmt.Sprintf("Hello, %s", text)
}
