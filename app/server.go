package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

const (
	CRLF      = "\r\n"
	OK        = "HTTP/1.1 200 OK"
	NOT_FOUND = "HTTP/1.1 404 Not Found"
)

var (
	WhiteSpace = []byte(" ")
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handle_connection(conn)
}

func handle_connection(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read data")
		os.Exit(1)
	}

	fmt.Println(string(buffer))

	request := http.ParseRequest(buffer)

	switch {
	case request.Path == "/":
		response := http.Ok()
		conn.Write(response.Serialize())
	case strings.HasPrefix(request.Path, "/echo"):
		handle_echo(conn, request)
	case strings.HasPrefix(request.Path, "/user-agent"):
		handle_user_agent(conn, request)
	default:
		response := http.NotFound()
		conn.Write(response.Serialize())
	}
}

func handle_echo(conn net.Conn, request *http.Request) {
	body, found := strings.CutPrefix(request.Path, "/echo/")

	if !found {
		fmt.Println("Failed to parse request")
		os.Exit(1)
	}

	response := &http.Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  200,
		Status:      "OK",
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}

	conn.Write(response.Serialize())
}

func handle_user_agent(conn net.Conn, request *http.Request) {
	response := &http.Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  200,
		Status:      "OK",
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(request.Headers["User-Agent"])),
		},
		Body: request.Headers["User-Agent"],
	}

	conn.Write(response.Serialize())
}
