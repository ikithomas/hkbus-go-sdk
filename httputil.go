package hkbus

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://data.etabus.gov.hk"
)

func getPath(ctx context.Context, client http.Client, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", baseURL, path)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
