package common_service

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/common_service/pb"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"syscall"
)

var (
	_ healthpb.HealthServer  = (*CommonServiceServer)(nil)
	_ pb.CommonServiceServer = (*CommonServiceServer)(nil)
)

type CommonServiceServer struct {
	pb.UnimplementedCommonServiceServer
	Logger *logrus.Logger
}

func NewCommonServiceServer(logger *logrus.Logger) *CommonServiceServer {
	return &CommonServiceServer{
		Logger: logger,
	}
}

func (c *CommonServiceServer) Echo(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *CommonServiceServer) Check(_ context.Context, request *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (c *CommonServiceServer) Watch(request *healthpb.HealthCheckRequest, server healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}

func (c *CommonServiceServer) Kill(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	if err != nil {
		return nil, err
	}
	err = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
