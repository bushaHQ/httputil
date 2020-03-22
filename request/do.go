package request

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Req the request struct
type Req struct {
	BaseURL *url.URL
	Path    string
	Body    io.Reader
	Header  http.Header
	Method  string
}

// Res holds the needed contents of the response
type Res struct {
	Status   int
	Body     []byte
	Header   http.Header
	Response *http.Response
}

// Do makes the request
func Do(r Req) (Res, error) {
	u, err := r.BaseURL.Parse(r.Path)
	if err != nil {
		return Res{}, err
	}

	req, err := http.NewRequest(r.Method, u.String(), r.Body)
	if err != nil {
		return Res{}, err
	}

	req.Header = r.Header
	c := http.Client{Timeout: time.Second * 30}

	res, err := c.Do(req)
	if err != nil {
		return Res{}, err
	}

	b, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		return Res{}, err
	}

	return Res{Status: res.StatusCode, Body: b, Header: res.Header, Response: res}, err
}
