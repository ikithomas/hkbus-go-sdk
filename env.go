package hkbus

import "net/http"

var (
	Client *http.Client
)

func init() {
	Client = &http.Client{}
}