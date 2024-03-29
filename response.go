package hkbus

import (
	"fmt"
	"strconv"
	"time"
)

type kmbResponse interface {
	Validate() error
}

var (
	_ kmbResponse = &routeListResponse{}
	_ kmbResponse = &routeGetResponse{}
	_ kmbResponse = &stopListResponse{}
	_ kmbResponse = &stopGetResponse{}
	_ kmbResponse = &routeStopListResponse{}
	_ kmbResponse = &routeStopGetResponse{}
	_ kmbResponse = &etaGetResponse{}
	_ kmbResponse = &stopEtaResponse{}
	_ kmbResponse = &routeEtaResponse{}

	DirectionInbound  = "inbound"
	DirectionOutbound = "outbound"

	version = "1.0"
)

/*
	  GET /v1/transport/kmb/route/

		{
		   "type":"RouteList",
		   "version":"1.0",
		   "generated_timestamp":"2020-11-29T11:40:48+08:00",
		   "data":[
		      {
		         "co":"KMB",
		         "route":"74B",
		         "bound":"O",
		         "service_type":"1",
		         "orig_en":"KOWLOON BAY",
		         "orig_tc":"九龍灣",
		         "orig_sc":"九龙湾",
		         "dest_en":"TAI PO CENTRAL",
		         "dest_tc":"大埔中心",
		         "dest_sc":"大埔中心",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"74B",
		         "bound":"I",
		         "service_type":"1",
		         "orig_en":"TAI PO CENTRAL",
		         "orig_tc":"大埔中心",
		         "orig_sc":"大埔中心",
		         "dest_en":"KWUN TONG FERRY",
		         "dest_tc":"觀塘碼頭",
		         "dest_sc":"观塘码头",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      }
		   ]
		}
*/
type routeListResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      []Route   `json:"data"`
}

// Validate implements kmbResponse
func (r *routeListResponse) Validate() error {
	if r.Type != "RouteList" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

/*
	  GET /v1/transport/kmb/route/{route}/{direction}/{service_type}

		{
	    "type": "Route",
	    "version": "1.0",
	    "generated_timestamp": "2020-11-29T11:40:48+08:00",
	    "data": {
	    "co": "KMB",
	    "route": "74B",
	    "bound": "O",
	    "service_type": "1",
	    "orig_en": "KOWLOON BAY",
	    "orig_tc": "九龍灣",
	    "orig_sc": "九龙湾",
	    "dest_en": "TAI PO CENTRAL",
	    "dest_tc": "大埔中心",
	    "dest_sc": "大埔中心",
	    "generated_timestamp": "2020-11-29T11:40:00+08:00"
	    }
	  }
*/
type routeGetResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      Route     `json:"data"`
}

// Validate implements kmbResponse
func (r *routeGetResponse) Validate() error {
	if r.Type != "Route" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

type Route struct {
	Route       string `json:"route"`
	Bound       string `json:"bound"`
	ServiceType string `json:"service_type"`
	OrigEn      string `json:"orig_en"`
	OrigTc      string `json:"orig_tc"`
	OrigSc      string `json:"orig_sc"`
	DestEn      string `json:"dest_en"`
	DestTc      string `json:"dest_tc"`
	DestSc      string `json:"dest_sc"`
}

/*
	  GET /v1/transport/kmb/stop

		{
		   "type":"StopList",
		   "version":"1.0",
		   "generated_timestamp":"2020-11-29T11:40:48+08:00",
		   "data":[
		      {
		         "stop":" A3ADFCDF8487ADB9",
		         "name_tc":"中秀茂坪",
		         "name_en":"SAU MAU PING (CENTRAL)",
		         "name_sc":"中秀茂坪",
		         "lat":22.318856,
		         "long":114.231353,
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      },
		      {
		         "stop":"6F106FD26B684372",
		         "name_en":"SAU ON HOUSE",
		         "name_tc":"秀安樓",
		         "name_sc":"秀安楼",
		         "lat":"22.316738",
		         "long":"114.233354",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      }
		   ]
		}
*/
type stopListResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      []Stop    `json:"data"`
}

// Validate implements kmbResponse
func (r *stopListResponse) Validate() error {
	if r.Type != "StopList" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

/*
	  GET /v1/transport/kmb/stop/{stop_id}

		{
		  "type": "Stop",
		  "version": "1.0",
		  "generated_timestamp": "2020-11-29T11:40:48+08:00",
		  "data": {
		    "stop": " A3ADFCDF8487ADB9",
		    "name_tc": "中秀茂坪",
		    "name_en": " SAU MAU PING (CENTRAL)",
		    "name_sc": "中秀茂坪",
		    "lat": 22.318856,
		    "long": 114.231353,
		    "generated_timestamp": "2020-11-29T11:40:00+08:00"
		  }
		}
*/
type stopGetResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      Stop      `json:"data"`
}

// Validate implements kmbResponse
func (r *stopGetResponse) Validate() error {
	if r.Type != "Stop" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

type Stop struct {
	Stop       string `json:"stop"`
	NameTc     string `json:"name_tc"`
	NameEn     string `json:"name_en"`
	NameSc     string `json:"name_sc"`
	LatString  string `json:"lat"`
	LongString string `json:"long"`
}

func (s Stop) Lat() (float64, error) {
	return strconv.ParseFloat(s.LatString, 64)
}

func (s Stop) Long() (float64, error) {
	return strconv.ParseFloat(s.LongString, 64)
}

/*
	  GET /v1/transport/kmb/route-stop

		{
		   "type":"RouteStopList",
		   "version":"1.0",
		   "generated_timestamp":"2020-11-29T11:40:48+08:00",
		   "data":[
		      {
		         "co":"KMB",
		         "route":"1A",
		         "bound":"O",
		         "service_type":"1",
		         "seq":1,
		         "stop":"A3ADFCDF8487ADB9",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"1A",
		         "bound":"O",
		         "service_type":"1",
		         "seq":34,
		         "stop":"DCFF4041D0C0ACF8",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"2",
		         "bound":"O",
		         "service_type":"1",
		         "seq":1,
		         "stop":"25BD7B50919AA221",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      }
		   ]
		}
*/
type routeStopListResponse struct {
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Timestamp time.Time   `json:"generated_timestamp"`
	Data      []RouteStop `json:"data"`
}

// Validate implements kmbResponse
func (r *routeStopListResponse) Validate() error {
	if r.Type != "RouteStopList" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

/*
	  GET /v1/transport/kmb/routestop/{route}/{direction}/{service_type}

		{
		   "type":"RouteStop",
		   "version":"1.0",
		   "generated_timestamp":"2020-11-29T11:40:48+08:00",
		   "data":[
		      {
		         "co":"KMB",
		         "route":"1A",
		         "bound":"O",
		         "service_type":"1",
		         "seq":1,
		         "stop":"A3ADFCDF8487ADB9",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"1A",
		         "bound":"O",
		         "service_type":"1",
		         "seq":34,
		         "stop":"DCFF4041D0C0ACF8",
		         "generated_timestamp":"2020-11-29T11:40:00+08:00"
		      }
		   ]
		}
*/
type routeStopGetResponse struct {
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Timestamp time.Time   `json:"generated_timestamp"`
	Data      []RouteStop `json:"data"`
}

// Validate implements kmbResponse
func (r *routeStopGetResponse) Validate() error {
	if r.Type != "RouteStop" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

type RouteStop struct {
	Route       string `json:"route"`
	Bound       string `json:"bound"`
	ServiceType string `json:"service_type"`
	SeqString   string `json:"seq"`
	Stop        string `json:"stop"`
}

func (r RouteStop) Seq() (int, error) {
	return strconv.Atoi(r.SeqString)
}

/*
	  GET /v1/transport/kmb/eta/{stop_id}/{route}/{service_type}

		{
		   "type":"ETA",
		   "version":"1.0",
		   "generated_timestamp":"2021-03-04T17:22:35+08:00",
		   "data":[
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":1,
		         "seq":7,
		         "stop":"A60AE774B09A5E44",
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":1,
		         "eta":"2021-03-04T17:25:51+08:00",
		         "rmk_tc":"",
		         "rmk_sc":"",
		         "rmk_en":"",
		         "generated_timestamp":"2021-03-04T17:22:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":1,
		         "seq":7,
		         "stop":"A60AE774B09A5E44",
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":2,
		         "eta":"2021-03-04T17:41:44+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-04T17:22:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":1,
		         "seq":7,
		         "stop":"A60AE774B09A5E44",
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":3,
		         "eta":"2021-03-04T17:54:08+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-04T17:22:15+08:00"
		      }
		   ]
		}
*/
type etaGetResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      []Eta     `json:"data"`
}

// Validate implements kmbResponse
func (r *etaGetResponse) Validate() error {
	if r.Type != "ETA" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

/*
	  GET /v1/transport/kmb/stop-eta/{stop_id}

		{
		   "type":"StopETA",
		   "version":"1.0",
		   "generated_timestamp":"2021-03-18T12:21:38+08:00",
		   "data":[
		      {
		         "co":"KMB",
		         "route":"234X",
		         "dir":"I",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"尖沙咀東(麼地道)",
		         "dest_sc":"尖沙咀东(么地道)",
		         "dest_en":"TSIM SHA TSUI EAST (MODY ROAD)",
		         "eta_seq":1,
		         "eta":"2021-03-18T12:31:59+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"234X",
		         "dir":"I",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"尖沙咀東(麼地道)",
		         "dest_sc":"尖沙咀东(么地道)",
		         "dest_en":"TSIM SHA TSUI EAST (MODY ROAD)",
		         "eta_seq":2,
		         "eta":"2021-03-18T12:46:59+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"234X",
		         "dir":"I",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"尖沙咀東(麼地道)",
		         "dest_sc":"尖沙咀东(么地道)",
		         "dest_en":"TSIM SHA TSUI EAST (MODY ROAD)",
		         "eta_seq":3,
		         "eta":"2021-03-18T13:02:00+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":1,
		         "seq":3,
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":1,
		         "eta":"2021-03-18T12:33:19+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":1,
		         "seq":3,
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":2,
		         "eta":"2021-03-18T12:53:19+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":1,
		         "seq":3,
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":3,
		         "eta":"2021-03-18T12:53:19+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":2,
		         "seq":3,
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":1,
		         "eta":"2021-03-18T12:33:19+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":2,
		         "seq":3,
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":2,
		         "eta":"2021-03-18T12:53:19+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"40",
		         "dir":"O",
		         "service_type":2,
		         "seq":3,
		         "dest_tc":"麗港城",
		         "dest_sc":"丽港城",
		         "dest_en":"LAGUNA CITY",
		         "eta_seq":3,
		         "eta":"2021-03-18T12:53:19+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"N48",
		         "dir":"I",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"愉翠苑",
		         "dest_sc":"愉翠苑",
		         "dest_en":"YU CHUI COURT",
		         "eta_seq":1,
		         "eta":null,
		         "rmk_tc":"服務只限於星期日及公眾假期",
		         "rmk_sc":"服务只限于星期日及公众假期",
		         "rmk_en":"This route only operates on Sundays and Public Holidays",
		         "generated_timestamp":"2021-03-18T12:21:15+08:00"
		      }
		   ]
		}
*/
type stopEtaResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      []Eta     `json:"data"`
}

// Validate implements kmbResponse
func (r *stopEtaResponse) Validate() error {
	if r.Type != "StopETA" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

/*
	  GET /v1/transport/kmb/route-eta/{route}/{service_type}

		{
		   "type":"RouteETA",
		   "version":"1.0",
		   "generated_timestamp":"2021-03-18T13:39:50+08:00",
		   "data":[
		      {
		         "co":"KMB",
		         "route":"3M",
		         "dir":"O",
		         "service_type":1,
		         "seq":1,
		         "dest_tc":"彩雲",
		         "dest_sc":"彩云",
		         "dest_en":"CHOI WAN",
		         "eta_seq":1,
		         "eta":"2021-03-18T13:50:00+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T13:39:46+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"3M",
		         "dir":"O",
		         "service_type":1,
		         "seq":1,
		         "dest_tc":"彩雲",
		         "dest_sc":"彩云",
		         "dest_en":"CHOI WAN",
		         "eta_seq":2,
		         "eta":"2021-03-18T14:10:00+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T13:39:46+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"3M",
		         "dir":"O",
		         "service_type":1,
		         "seq":1,
		         "dest_tc":"彩雲",
		         "dest_sc":"彩云",
		         "dest_en":"CHOI WAN",
		         "eta_seq":3,
		         "eta":"2021-03-18T14:30:00+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T13:39:46+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"3M",
		         "dir":"O",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"彩雲",
		         "dest_sc":"彩云",
		         "dest_en":"CHOI WAN",
		         "eta_seq":1,
		         "eta":"2021-03-18T13:51:29+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T13:39:46+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"3M",
		         "dir":"O",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"彩雲",
		         "dest_sc":"彩云",
		         "dest_en":"CHOI WAN",
		         "eta_seq":2,
		         "eta":"2021-03-18T14:11:28+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T13:39:46+08:00"
		      },
		      {
		         "co":"KMB",
		         "route":"3M",
		         "dir":"O",
		         "service_type":1,
		         "seq":2,
		         "dest_tc":"彩雲",
		         "dest_sc":"彩云",
		         "dest_en":"CHOI WAN",
		         "eta_seq":3,
		         "eta":"2021-03-18T14:31:28+08:00",
		         "rmk_tc":"原定班次",
		         "rmk_sc":"原定班次",
		         "rmk_en":"Scheduled Bus",
		         "generated_timestamp":"2021-03-18T13:39:46+08:00"
		      }
		   ]
		}
*/
type routeEtaResponse struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"generated_timestamp"`
	Data      []Eta     `json:"data"`
}

// Validate implements kmbResponse
func (r *routeEtaResponse) Validate() error {
	if r.Type != "RouteETA" {
		return fmt.Errorf("%w type %s", errBadResponse, r.Type)
	}
	if r.Version != version {
		return fmt.Errorf("%w version %s", errBadResponse, r.Version)
	}
	return nil
}

type Eta struct {
	Route       string    `json:"route"`
	Dir         string    `json:"dir"`
	ServiceType int       `json:"service_type"`
	Seq         int       `json:"seq"`
	Stop        string    `json:"stop"`
	DestTc      string    `json:"dest_tc"`
	DestSc      string    `json:"dest_sc"`
	DestEn      string    `json:"dest_en"`
	EtaSeq      int       `json:"eta_seq"`
	Eta         time.Time `json:"eta"`
	RmkTc       string    `json:"rmk_tc"`
	RmkSc       string    `json:"rmk_sc"`
	RmkEn       string    `json:"rmk_en"`
	Timestamp   time.Time `json:"generated_timestamp"`
}

/*
	{
	  "code": "422",
	  "message": " Invalid/Missing parameter(s)"
	}
*/
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
