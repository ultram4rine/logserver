package logserver

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a endpoints wrapper.
type Endpoints struct {
	LogEndpoint endpoint.Endpoint
}

// MakeDHCPEndpoint creates endpoint for DHCP logs.
func MakeDHCPEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DHCPLogsRequest)

		logs, err := svc.GetDHCPLogs(ctx, req.MAC, req.From, req.To)
		if err != nil {
			return nil, err
		}

		return logs, nil
	}
}

// MakeSwitchEndpoint creates endpoint for logs from switch.
func MakeSwitchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SwitchLogsRequest)

		logs, err := svc.GetSwitchLogs(ctx, req.Name, req.From, req.To)
		if err != nil {
			return nil, err
		}

		return logs, nil
	}
}

// MakeSimilarEndpoint creates endpoint for similar available switches.
func MakeSimilarEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SwName)

		switches, err := svc.GetSimilarSwitches(ctx, req.Name)
		if err != nil {
			return nil, err
		}

		return switches, nil
	}
}

// DHCPLogsRequest is a request for DHCP logs.
type DHCPLogsRequest struct {
	MAC  uint64 `json:"mac"`
	From string `json:"from"`
	To   string `json:"to"`
}

type dhcpLog struct {
	IP        string `json:"ip"`
	TimeStamp string `json:"ts"`
	Message   string `json:"message"`
}

// DHCPLogs is a response with DHCP logs.
type DHCPLogs struct {
	Logs []dhcpLog `json:"logs"`
	Err  string    `json:"err,omitempty"`
}

// SwitchLogsRequest is a request for logs from switch.
type SwitchLogsRequest struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

type switchLog struct {
	IP        string `json:"ip"`
	Name      string `json:"name"`
	TimeStamp string `json:"ts"`
	Message   string `json:"message"`
}

// SwitchLogs is a response with logs from switch.
type SwitchLogs struct {
	Logs []SwitchLogs `json:"logs"`
	Err  string       `json:"err,omitempty"`
}

// SwName is a request for similar available switches.
type SwName struct {
	Name string `json:"name"`
}

type sw struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

// Switches is a response with array of similar available switches.
type Switches struct {
	Sws []sw   `json:"switches"`
	Err string `json:"err,omitempty"`
}
