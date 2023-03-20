package hkbus_test

import (
	"context"
	"testing"

	"github.com/ikithomas/hkbus-go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestRouteList(t *testing.T) {
	got, err := hkbus.RouteList(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1501, len(got))
}

func TestRouteGet(t *testing.T) {
	got, err := hkbus.RouteGet(context.Background(), "74B", "outbound", "1")
	assert.NoError(t, err)
	assert.Equal(t, "O", got.Bound)
}

func TestStopList(t *testing.T) {
	got, err := hkbus.StopList(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 6389, len(got))
}

func TestStopGet(t *testing.T) {
	got, err := hkbus.StopGet(context.Background(), "A3ADFCDF8487ADB9")
	assert.NoError(t, err)
	assert.Equal(t, "SAU MAU PING (CENTRAL)", got.NameEn)
}

func TestRouteStopList(t *testing.T) {
	got, err := hkbus.RouteStopList(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 33755, len(got))
}

func TestRouteStopGet(t *testing.T) {
	got, err := hkbus.RouteStopGet(context.Background(), "1A", hkbus.DirectionOutbound, "1")
	assert.NoError(t, err)
	assert.Equal(t, 34, len(got))
}

func TestEtaGet(t *testing.T) {
	got, err := hkbus.EtaGet(context.Background(), "A60AE774B09A5E44", "40", "1")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(got))
}

func TestStopEtaGet(t *testing.T) {
	got, err := hkbus.StopEtaGet(context.Background(), "B8B04CD1E568B8F6")
	assert.NoError(t, err)
	assert.Equal(t, 9, len(got))
}

func TestRouteEtaGet(t *testing.T) {
	got, err := hkbus.RouteEtaGet(context.Background(), "3M", "1")
	assert.NoError(t, err)
	assert.Equal(t, 25, len(got))
}
