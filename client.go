package hkbus

import (
	"context"
	"net/http"
)

type Client struct {
	Client http.Client
}

func (s Client) RouteList(ctx context.Context) ([]Route, error) {
	return RouteList(ctx, s.Client)
}

func (s Client) RouteGet(ctx context.Context, route string, direction string, serviceType string) (Route, error) {
	return RouteGet(ctx, s.Client, route, direction, serviceType)
}

func (s Client) StopList(ctx context.Context) ([]Stop, error) {
	return StopList(ctx, s.Client)
}

func (s Client) StopGet(ctx context.Context, stopId string) (Stop, error) {
	return StopGet(ctx, s.Client, stopId)
}

func (s Client) RouteStopList(ctx context.Context) ([]RouteStop, error) {
	return RouteStopList(ctx, s.Client)
}

func (s Client) RouteStopGet(ctx context.Context, route string, direction string, serviceType string) ([]RouteStop, error) {
	return RouteStopGet(ctx, s.Client, route, direction, serviceType)
}

func (s Client) EtaGet(ctx context.Context, route string, direction string, serviceType string) ([]Eta, error) {
	return EtaGet(ctx, s.Client, route, direction, serviceType)
}

func (s Client) StopEtaGet(ctx context.Context, stopId string) ([]Eta, error) {
	return StopEtaGet(ctx, s.Client, stopId)
}

func (s Client) RouteEtaGet(ctx context.Context, route string, serviceType string) ([]Eta, error) {
	return RouteEtaGet(ctx, s.Client, route, serviceType)
}
