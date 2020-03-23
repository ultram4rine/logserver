package client

import (
	"git.sgu.ru/ultramarine/logserver"
	pb "git.sgu.ru/ultramarine/logserver/logpb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// New creates new gPRC client for logserver.
func New(conn *grpc.ClientConn) logserver.Service {
	var dhcpEndpoint = grpctransport.NewClient(
		conn, "logpb.Log", "GetDHCPLogs",
		logserver.EncodeDHCPLogsRequest,
		logserver.DecodeDHCPLogsResponse,
		pb.DHCPLogsResponse{},
	).Endpoint()

	var switchEndpoint = grpctransport.NewClient(
		conn, "logpb.Log", "GetSwitchLogs",
		logserver.EncodeSwitchLogsRequest,
		logserver.DecodeSwitchLogsResponse,
		pb.SwitchLogsResponse{},
	).Endpoint()

	var similarEndpoint = grpctransport.NewClient(
		conn, "logpb.Log", "GetSimilarSwitches",
		logserver.EncodeSimilarSwitchesRequest,
		logserver.DecodeSimilarSwitchesResponse,
		pb.SimilarSwitchesResponse{},
	).Endpoint()

	return logserver.Endpoints{
		DHCPEndpoint:    dhcpEndpoint,
		SwitchEndpoint:  switchEndpoint,
		SimilarEndpoint: similarEndpoint,
	}
}
