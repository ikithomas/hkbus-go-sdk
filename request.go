package hkbus

import (
	"context"
	"io"
)

func HttpGet(ctx context.Context, url string) ([]byte, error) {
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
