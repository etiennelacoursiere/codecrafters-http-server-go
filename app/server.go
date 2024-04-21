package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/handlers"
	"github.com/codecrafters-io/http-server-starter-go/http"
)

var Directory string

func main() {
	flag.StringVar(&Directory, "directory", "", "Directory")
	flag.Parse()
	handlers.Directory = Directory

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handle_connection(conn)
	}
}

func handle_connection(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read data")
		os.Exit(1)
	}

	// fmt.Println(string(buffer))

	request := http.ParseRequest(buffer)

	switch {
	case request.Path == "/":
		response := http.Ok()
		conn.Write(response.Serialize())
	case strings.HasPrefix(request.Path, "/echo"):
		handlers.HandleEcho(conn, request)
	case strings.HasPrefix(request.Path, "/user-agent"):
		handlers.HandleUserAgent(conn, request)
	case strings.HasPrefix(request.Path, "/files"):
		handlers.HandleFile(conn, request)
	default:
		response := http.NotFound()
		conn.Write(response.Serialize())
	}

	conn.Close()
}
