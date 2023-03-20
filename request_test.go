package hkbus_test

import (
	"context"
	"testing"

	"github.com/ikithomas/hkbus-go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	b, err := hkbus.HttpGet(context.Background(), "https://captive.apple.com")
	assert.NoError(t, err)
	assert.Equal(t, "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>\n", string(b))
}
