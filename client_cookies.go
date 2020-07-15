package request

import (
	"net/http"
	"net/url"
)

func (rq *Client) GetCookie(host string, name string) string {
	a := rq.client.Jar.Cookies(&url.URL{Scheme: "http", Host: host, Path: "/"})
	for _, c := range a {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}

func (rq *Client) SetCookie(host string, name string, value string) {
	rq.client.Jar.SetCookies(&url.URL{Scheme: "http", Host: host, Path: "/"}, []*http.Cookie{{Name: name, Value: value}})
}
