package hkbus

import (
	"context"
	"encoding/json"
	"fmt"
)

func RouteList(ctx context.Context) ([]Route, error) {
	path := "/v1/transport/kmb/route/"
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r routeListResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}

func RouteGet(ctx context.Context, route string, direction string, serviceType string) (Route, error) {
	path := fmt.Sprintf("/v1/transport/kmb/route/%s/%s/%s", route, direction, serviceType)
	b, err := getPath(ctx, path)
	if err != nil {
		return Route{}, err
	}

	var r routeGetResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return Route{}, err
	}

	if err := r.Validate(); err != nil {
		return Route{}, err
	}

	return r.Data, nil
}

func StopList(ctx context.Context) ([]Stop, error) {
	path := "/v1/transport/kmb/stop"
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r stopListResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}

func StopGet(ctx context.Context, stopId string) (Stop, error) {
	path := fmt.Sprintf("/v1/transport/kmb/stop/%s", stopId)
	b, err := getPath(ctx, path)
	if err != nil {
		return Stop{}, err
	}

	var r stopGetResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return Stop{}, err
	}

	if err := r.Validate(); err != nil {
		return Stop{}, err
	}

	return r.Data, nil
}

func RouteStopList(ctx context.Context) ([]RouteStop, error) {
	path := "/v1/transport/kmb/route-stop"
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r routeStopListResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}

func RouteStopGet(ctx context.Context, route string, direction string, serviceType string) ([]RouteStop, error) {
	path := fmt.Sprintf("/v1/transport/kmb/route-stop/%s/%s/%s", route, direction, serviceType)
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r routeStopGetResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}

func EtaGet(ctx context.Context, route string, direction string, serviceType string) ([]Eta, error) {
	path := fmt.Sprintf("/v1/transport/kmb/eta/%s/%s/%s", route, direction, serviceType)
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r etaGetResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}

func StopEtaGet(ctx context.Context, stopId string) ([]Eta, error) {
	path := fmt.Sprintf("/v1/transport/kmb/stop-eta/%s", stopId)
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r stopEtaResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}

func RouteEtaGet(ctx context.Context, route string, serviceType string) ([]Eta, error) {
	path := fmt.Sprintf("/v1/transport/kmb/route-eta/%s/%s", route, serviceType)
	b, err := getPath(ctx, path)
	if err != nil {
		return nil, err
	}

	var r routeEtaResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r.Data, nil
}
