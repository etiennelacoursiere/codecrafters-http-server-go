package http

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	Bytes_WhiteSpace = []byte(" ")
	Bytes_CRLF       = []byte("\r\n")
)

type Request struct {
	Method      string
	Path        string
	HTTPVersion string
	Headers     map[string]string
}

func ParseRequest(data []byte) *Request {
	lines := bytes.Split(data, Bytes_CRLF)
	start_line_parts := bytes.Split(lines[0], Bytes_WhiteSpace)

	request := &Request{}
	request.Method = string(start_line_parts[0])
	request.Path = string(start_line_parts[1])
	request.HTTPVersion = string(start_line_parts[2])

	headers := make(map[string]string)

	for i := 1; string(lines[i]) != ""; i++ {
		line := string(lines[i])
		parts := strings.Split(line, ": ")
		headers[parts[0]] = parts[1]
	}

	request.Headers = headers

	fmt.Printf("%#v\n", request)

	return request
}
