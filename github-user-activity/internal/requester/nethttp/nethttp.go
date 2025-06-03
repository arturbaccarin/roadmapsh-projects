package nethttp

import (
	"log"
	"net/http"
)

type NetHTTP struct {
	client *http.Client
}

func New() *NetHTTP {
	return &NetHTTP{
		client: &http.Client{},
	}
}

func (n *NetHTTP) Get(url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := n.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
