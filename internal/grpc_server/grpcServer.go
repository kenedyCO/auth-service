package grpcServer

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kenedyCO/auth-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50052

type server struct {
	auth_v1.UnimplementedAuthV1Server
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	auth_v1.RegisterAuthV1Server(s, &server{})

	log.Printf("GRPC server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *auth_v1.CreateRequest) (*auth_v1.CreateResponse, error) {
	log.Printf("req: %v", req)

	return &auth_v1.CreateResponse{
		Id: 123,
	}, nil
}
func (s *server) Get(ctx context.Context, req *auth_v1.GetRequest) (*auth_v1.GetResponse, error) {
	log.Printf("req: %v", req)

	return &auth_v1.GetResponse{}, nil
}
func (s *server) Update(ctx context.Context, req *auth_v1.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("req: %v", req)

	return nil, nil
}
func (s *server) Delete(ctx context.Context, req *auth_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("req: %v", req)

	return nil, nil
}
