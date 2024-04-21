package http

import "strconv"

const (
	CRLF       = "\r\n"
	WhiteSpace = " "
)

type Response struct {
	HTTPVersion string
	StatusCode  int
	Status      string
	Headers     map[string]string
	Body        string
}

func Ok() *Response {
	return &Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  200,
		Status:      "OK",
	}
}

func NotFound() *Response {
	return &Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  404,
		Status:      "Not Found",
	}
}

func InternalServerError() *Response {
	return &Response{
		HTTPVersion: "HTTP/1.1",
		StatusCode:  500,
		Status:      "Internal Server Error",
	}
}

func (r *Response) serialize_headers() string {
	headers := ""

	if r.Headers == nil {
		return headers
	}

	for key, value := range r.Headers {
		headers += key + ": " + value + CRLF
	}

	return headers
}

func (r *Response) Serialize() []byte {
	start_line := r.HTTPVersion + WhiteSpace + strconv.Itoa(r.StatusCode) + WhiteSpace + r.Status + CRLF
	headers := r.serialize_headers() + CRLF
	return []byte(start_line + headers + r.Body)
}
