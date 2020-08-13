package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(connection, "Hello from Server")
		go handle(connection)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()

	fmt.Println("Connection Closed!")
}
