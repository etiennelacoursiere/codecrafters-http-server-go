package handlers

import (
	"fmt"
	"net"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

func HandleFile(conn net.Conn, request *http.Request) {
	switch request.Method {
	case "GET":
		getFile(conn, request)
	case "POST":
		postFile(conn, request)
	}

}

func getFile(conn net.Conn, request *http.Request) {
	filename, found := strings.CutPrefix(request.Path, "/files/")

	if !found {
		fmt.Println("Failed to parse request")
		os.Exit(1)
	}

	filepath := path.Join(Directory, filename)
	buffer, error := os.ReadFile(filepath)

	if error != nil {
		response := http.NotFound()
		conn.Write(response.Serialize())
		return
	}

	response := &http.Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  200,
		Status:      "OK",
		Headers: map[string]string{
			"Content-Type":   "application/octet-stream",
			"Content-Length": strconv.Itoa(len(buffer)),
		},
		Body: string(buffer),
	}

	conn.Write(response.Serialize())
}

func postFile(conn net.Conn, request *http.Request) {
	filename, found := strings.CutPrefix(request.Path, "/files/")

	if !found {
		fmt.Println("Failed to parse request")
		os.Exit(1)
	}

	filepath := path.Join(Directory, filename)
	err := os.WriteFile(filepath, []byte(request.Body), 0666)

	if err != nil {
		response := http.InternalServerError()
		conn.Write(response.Serialize())
		return
	}

	response := http.Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  201,
		Status:      "Created",
	}

	conn.Write(response.Serialize())
}
