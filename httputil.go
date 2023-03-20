package hkbus

import (
	"context"
	"fmt"
	"io"
)

var (
	baseURL = "https://data.etabus.gov.hk"
)

func getPath(ctx context.Context, path string) ([]byte, error) {
	return httpGet(ctx, fmt.Sprintf("%s/%s", baseURL, path))
}

func httpGet(ctx context.Context, url string) ([]byte, error) {
	res, err := Client.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
