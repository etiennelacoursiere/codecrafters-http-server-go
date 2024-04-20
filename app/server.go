package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

type HTTPRequest struct {
	Method      string
	Path        string
	HTTPVersion string
}

func NewHTTPRequest(method, path, version string) *HTTPRequest {
	return &HTTPRequest{
		Method:      method,
		Path:        path,
		HTTPVersion: version,
	}
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

	// _, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
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

	request := NewHTTPRequest(string(headerParts[0]), string(headerParts[1]), string(headerParts[2]))

	if request.Path == "/" {
		conn.Write([]byte(OK + CRLF + CRLF))
	} else {
		conn.Write([]byte(NOT_FOUND + CRLF + CRLF))
	}

}
