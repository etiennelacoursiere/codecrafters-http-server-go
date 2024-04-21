package handlers

import (
	"net"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

func HandleUserAgent(conn net.Conn, request *http.Request) {
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
