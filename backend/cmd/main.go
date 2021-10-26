package main

import (
	"flag"
	"log"
	"net"

	"github.com/jestradaramos/group-lift/backend/pkg/domain"
	"github.com/jestradaramos/group-lift/backend/pkg/user"
	"google.golang.org/grpc"
)

func newUserServiceServer() *user.UserService {
	return &user.UserService{}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	domain.RegisterUserServiceServer(grpcServer, newUserServiceServer())
	grpcServer.Serve(lis)
}
