package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	port := flag.Uint64("port", 6969, "http port")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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

	buffer := make([]byte, 512)
	_, err := c.Read(buffer)
	if err != nil {
		log.Println(err)
	}

	var file string
	requestData := strings.Split(string(buffer), "\r\n")

	// for better experience, create a Parser to parse the request (method, path, relative path, headers, etc)
	if requestData[0] == "GET / HTTP/1.1" {
		file = "index.html"
	} else {
		file = "404.html"
	}

	data, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
		return
	}
	if len(data) <= 1 {
		data = []byte("something went wronG")
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
