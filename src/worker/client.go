package worker

import (
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

type Client struct {
	httpClient *http.Client
	userAgent  string
}

func (c *Client) get(url string) (*http.Response, error) {
	return c.getConditional(url, "", "")
}

func (c *Client) getConditional(url, lastModified, etag string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	if lastModified != "" {
		req.Header.Set("If-Modified-Since", lastModified)
	}
	if etag != "" {
		req.Header.Set("If-None-Match", etag)
	}
	return c.httpClient.Do(req)
}

var client *Client

func init() {
	r := req.C().
		SetProxy(http.ProxyFromEnvironment).
		SetTimeout(time.Second * 30).
		SetTLSHandshakeTimeout(time.Second * 10).
		DisableKeepAlives()
	r.ImpersonateChrome()
	client = &Client{
		httpClient: r.GetClient(),
		userAgent:  "Feedport/1.0",
	}
}
