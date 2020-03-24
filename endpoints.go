package logserver

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a endpoints wrapper.
type Endpoints struct {
	DHCPEndpoint    endpoint.Endpoint
	SwitchEndpoint  endpoint.Endpoint
	SimilarEndpoint endpoint.Endpoint
}

func (e Endpoints) GetDHCPLogs(ctx context.Context, mac, from, to string) (DHCPLogsResponse, error) {
	req := DHCPLogsRequest{
		MAC:  mac,
		From: from,
		To:   to,
	}
	resp, err := e.DHCPEndpoint(ctx, req)
	if err != nil {
		return DHCPLogsResponse{}, err
	}
	dhcpResp := resp.(DHCPLogsResponse)
	if dhcpResp.Err != "" {
		return DHCPLogsResponse{}, errors.New(dhcpResp.Err)
	}
	return dhcpResp, nil
}

func (e Endpoints) GetSwitchLogs(ctx context.Context, name, from, to string) (SwitchLogsResponse, error) {
	req := SwitchLogsRequest{
		Name: name,
		From: from,
		To:   to,
	}
	resp, err := e.SwitchEndpoint(ctx, req)
	if err != nil {
		return SwitchLogsResponse{}, err
	}
	switchResp := resp.(SwitchLogsResponse)
	if switchResp.Err != "" {
		return SwitchLogsResponse{}, errors.New(switchResp.Err)
	}
	return switchResp, nil
}

func (e Endpoints) GetSimilarSwitches(ctx context.Context, name string) (SimilarSwitchesResponse, error) {
	req := SimilarSwitchesRequest{
		Name: name,
	}
	resp, err := e.SimilarEndpoint(ctx, req)
	if err != nil {
		return SimilarSwitchesResponse{}, err
	}
	similarResp := resp.(SimilarSwitchesResponse)
	if similarResp.Err != "" {
		return SimilarSwitchesResponse{}, errors.New(similarResp.Err)
	}
	return similarResp, nil
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
		req := request.(SimilarSwitchesRequest)

		switches, err := svc.GetSimilarSwitches(ctx, req.Name)
		if err != nil {
			return nil, err
		}

		return switches, nil
	}
}

// DHCPLogsRequest is a request for DHCP logs.
type DHCPLogsRequest struct {
	MAC  string `json:"mac"`
	From string `json:"from"`
	To   string `json:"to"`
}

type dhcpLog struct {
	IP        string `json:"ip"`
	TimeStamp string `json:"ts"`
	Message   string `json:"message"`
}

// DHCPLogsResponse is a response with DHCP logs.
type DHCPLogsResponse struct {
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

// SwitchLogsResponse is a response with logs from switch.
type SwitchLogsResponse struct {
	Logs []switchLog `json:"logs"`
	Err  string      `json:"err,omitempty"`
}

// SimilarSwitchRequest is a request for similar available switches.
type SimilarSwitchesRequest struct {
	Name string `json:"name"`
}

type similarSwitch struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

// SimilarSwitchResponse is a response with array of similar available switches.
type SimilarSwitchesResponse struct {
	Sws []similarSwitch `json:"switches"`
	Err string          `json:"err,omitempty"`
}
