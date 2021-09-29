package common_service

import (
	"context"
	"os"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/common_service/pb"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	_ healthpb.HealthServer  = (*CommonServiceServer)(nil)
	_ pb.CommonServiceServer = (*CommonServiceServer)(nil)
)

type CommonServiceServer struct {
	pb.UnimplementedCommonServiceServer
	Logger    *logrus.Logger
	AllowKill bool
}

func NewCommonServiceServer(logger *logrus.Logger, allowKill bool) *CommonServiceServer {
	return &CommonServiceServer{
		Logger:    logger,
		AllowKill: allowKill,
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
	c.Logger.Warnf("Recieved shutting down request at %s", time.Now())
	if c.AllowKill {
		c.Logger.Warnf("Accepted shutting down request at %s", time.Now())
		// err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		// if err != nil {
		// 	return nil, err
		// }
		// err = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		// if err != nil {
		// 	return nil, err
		// }
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			return nil, err
		}
		err = process.Signal(syscall.SIGTERM)
		if err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	}
	c.Logger.Warnf("Rejected shutting down request at %s", time.Now())
	return nil, status.Errorf(codes.PermissionDenied, "Do you like school?")
}
