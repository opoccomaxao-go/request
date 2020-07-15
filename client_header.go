package request

func (rq *Client) GetHeader(key string) string {
	return rq.defaultHeaders[key]
}

func (rq *Client) SetHeader(key string, value string) {
	rq.defaultHeaders[key] = value
}
