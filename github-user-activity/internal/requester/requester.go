package requester

import "net/http"

type Requester interface {
	Get(url string, header map[string]string) (*http.Response, error)
}
