package server

import (
	"fmt"
	"github/eugene-krivtsov/idler-email/internal/transport/grpc/api"
	"google.golang.org/grpc"
	"net"
)

const (
	protocol = "tcp"
)

type GrpcServer struct {
	grpcServer  *grpc.Server
	MailHandler api.MailSenderServiceServer
}

func NewGrpcServer(mailHandler api.MailSenderServiceServer) *GrpcServer {
	return &GrpcServer{
		grpcServer:  grpc.NewServer(),
		MailHandler: mailHandler,
	}
}

func (s *GrpcServer) Run(port int) error {
	lis, err := net.Listen(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	api.RegisterMailSenderServiceServer(s.grpcServer, s.MailHandler)

	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *GrpcServer) Shutdown() {
	s.grpcServer.GracefulStop()
}
