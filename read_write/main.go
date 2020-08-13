package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			fmt.Println(connection, "Hello from Ravi's Server")
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	err := connection.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection Timeout!!!")
	}
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(connection, "I heard what u say: %s\n", line)
	}
	defer connection.Close()
	fmt.Println("Connection closed!!!")
}
