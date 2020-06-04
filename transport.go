package logserver

import (
	"context"

	pb "git.sgu.ru/ultramarine/logserver/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	DHCP    grpctransport.Handler
	Switch  grpctransport.Handler
	Similar grpctransport.Handler
}

func (s *grpcServer) GetDHCPLogs(ctx context.Context, req *pb.DHCPLogsRequest) (*pb.DHCPLogsResponse, error) {
	_, resp, err := s.DHCP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DHCPLogsResponse), nil
}

func (s *grpcServer) GetSwitchLogs(ctx context.Context, req *pb.SwitchLogsRequest) (*pb.SwitchLogsResponse, error) {
	_, resp, err := s.Switch.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SwitchLogsResponse), nil
}

func (s *grpcServer) GetSimilarSwitches(ctx context.Context, req *pb.SimilarSwitchesRequest) (*pb.SimilarSwitchesResponse, error) {
	_, resp, err := s.Similar.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SimilarSwitchesResponse), nil
}

// NewGRPCServer creates new gRPC server with endpoints.
func NewGRPCServer(_ context.Context, endpoint Endpoints) pb.LogServiceServer {
	return &grpcServer{
		DHCP: grpctransport.NewServer(
			endpoint.DHCPEndpoint,
			DecodeDHCPLogsRequest,
			EncodeDHCPLogsResponse,
		),
		Switch: grpctransport.NewServer(
			endpoint.SwitchEndpoint,
			DecodeSwitchLogsRequest,
			EncodeSwitchLogsResponse,
		),
		Similar: grpctransport.NewServer(
			endpoint.SimilarEndpoint,
			DecodeSimilarSwitchesRequest,
			EncodeSimilarSwitchesResponse,
		),
	}
}
