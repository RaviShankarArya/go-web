package main

import (
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
			log.Println(err)
		}

		io.WriteString(connection, "\n Hello From TCP server \n")
		fmt.Fprintln(connection, "How is your day?")
		fmt.Fprintf(connection, "%v", "Well, I hope")

		connection.Close()

	}
}
