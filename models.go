package logserver

import (
	"context"

	pb "git.sgu.ru/ultramarine/logserver/pb"
)

// EncodeDHCPLogsRequest encodes DHCP logs request.
func EncodeDHCPLogsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(DHCPLogsRequest)
	return &pb.DHCPLogsRequest{
		MAC:  req.MAC,
		From: req.From,
		To:   req.To,
	}, nil
}

// DecodeDHCPLogsRequest decodes DHCP logs request.
func DecodeDHCPLogsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.DHCPLogsRequest)
	return DHCPLogsRequest{
		MAC:  req.MAC,
		From: req.From,
		To:   req.To,
	}, nil
}

// EncodeDHCPLogsResponse encodes DHCP logs response.
func EncodeDHCPLogsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(DHCPLogsResponse)
	var logs []*pb.DHCPLog
	for _, l := range res.Logs {
		log := &pb.DHCPLog{
			Ip:        l.IP,
			Timestamp: l.TimeStamp,
			Message:   l.Message,
		}
		logs = append(logs, log)
	}
	return &pb.DHCPLogsResponse{
		Log: logs,
		Err: res.Err,
	}, nil
}

// DecodeDHCPLogsResponse decodes DHCP logs response.
func DecodeDHCPLogsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.DHCPLogsResponse)
	var logs []dhcpLog
	for _, l := range res.Log {
		log := dhcpLog{
			IP:        l.Ip,
			TimeStamp: l.Timestamp,
			Message:   l.Message,
		}
		logs = append(logs, log)
	}
	return DHCPLogsResponse{
		Logs: logs,
		Err:  res.Err,
	}, nil
}

// EncodeSwitchLogsRequest encodes switch logs request.
func EncodeSwitchLogsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(SwitchLogsRequest)
	return &pb.SwitchLogsRequest{
		Name: req.Name,
		From: req.From,
		To:   req.To,
	}, nil
}

// DecodeSwitchLogsRequest decodes switch logs request.
func DecodeSwitchLogsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SwitchLogsRequest)
	return SwitchLogsRequest{
		Name: req.Name,
		From: req.From,
		To:   req.To,
	}, nil
}

// EncodeSwitchLogsResponse encodes switch logs response.
func EncodeSwitchLogsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SwitchLogsResponse)
	var logs []*pb.SwitchLog
	for _, l := range res.Logs {
		log := &pb.SwitchLog{
			Ip:      l.IP,
			Name:    l.Name,
			Ts:      l.TimeStamp,
			Message: l.Message,
		}
		logs = append(logs, log)
	}
	return &pb.SwitchLogsResponse{
		Log: logs,
		Err: res.Err,
	}, nil
}

// DecodeSwitchLogsResponse decodes switch logs response.
func DecodeSwitchLogsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.SwitchLogsResponse)
	var logs []switchLog
	for _, l := range res.Log {
		log := switchLog{
			IP:        l.Ip,
			Name:      l.Name,
			TimeStamp: l.Ts,
			Message:   l.Message,
		}
		logs = append(logs, log)
	}
	return SwitchLogsResponse{
		Logs: logs,
		Err:  res.Err,
	}, nil
}

// EncodeSimilarSwitchesRequest encodes similar switches request.
func EncodeSimilarSwitchesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(SimilarSwitchesRequest)
	return &pb.SimilarSwitchesRequest{
		Name: req.Name,
	}, nil
}

// DecodeSimilarSwitchesRequest decodes similar switches request.
func DecodeSimilarSwitchesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SimilarSwitchesRequest)
	return SimilarSwitchesRequest{
		Name: req.Name,
	}, nil
}

// EncodeSimilarSwitchesResponse encodes similar switches response.
func EncodeSimilarSwitchesResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SimilarSwitchesResponse)
	var sws []*pb.SimilarSwitch
	for _, l := range res.Sws {
		s := &pb.SimilarSwitch{
			IP:   l.IP,
			Name: l.Name,
		}
		sws = append(sws, s)
	}
	return &pb.SimilarSwitchesResponse{
		Switch: sws,
		Err:    res.Err,
	}, nil
}

// DecodeSimilarSwitchesResponse decodes similar switches response.
func DecodeSimilarSwitchesResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.SimilarSwitchesResponse)
	var sws []similarSwitch
	for _, l := range res.Switch {
		s := similarSwitch{
			IP:   l.IP,
			Name: l.Name,
		}
		sws = append(sws, s)
	}
	return SimilarSwitchesResponse{
		Sws: sws,
		Err: res.Err,
	}, nil
}
