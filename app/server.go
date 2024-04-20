package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type HTTPRequest struct {
	Method      string
	Path        string
	HTTPVersion string
}

type HTTPResponse struct {
	HTTPVersion string
	StatusCode  int
	Status      string
	Headers     map[string]string
	Body        string
}

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

	requestParts := bytes.Split(buffer, []byte(CRLF))
	headerParts := bytes.Split(requestParts[0], []byte(WhiteSpace))

	request := &HTTPRequest{
		Method:      string(headerParts[0]),
		Path:        string(headerParts[1]),
		HTTPVersion: string(headerParts[2]),
	}

	switch {
	case request.Path == "/":
		conn.Write([]byte(OK + CRLF + CRLF))
	case strings.HasPrefix(request.Path, "/echo"):
		handle_echo(conn, request)
	default:
		conn.Write([]byte(NOT_FOUND + CRLF + CRLF))
	}
}

func handle_echo(conn net.Conn, request *HTTPRequest) {
	body, found := strings.CutPrefix(request.Path, "/echo/")

	if !found {
		fmt.Println("Failed to parse request")
		os.Exit(1)
	}

	content_type := "Content-Type: text/plain"
	content_length := "Content-Length: " + strconv.Itoa(len(body))
	conn.Write([]byte(OK + CRLF + content_type + CRLF + content_length + CRLF + CRLF + body))
}
