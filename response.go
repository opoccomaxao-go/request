package request

import "net/http"

type Response struct {
	Status   int
	Error    error
	Response []byte
	Headers  http.Header
}

func (r *Response) Text() string {
	return string(r.Response)
}
