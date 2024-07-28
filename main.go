package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/html\r\nConnection: close\r\n\r\n%s",
		len(data),
		string(data),
	)
	_, err = c.Write([]byte(response))
	if err != nil {
		log.Println(err)
	}
}
