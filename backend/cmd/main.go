package main

import (
	"log"
	"net"

	"github.com/jestradaramos/group-lift/backend/pkg/domain"
	"github.com/jestradaramos/group-lift/backend/pkg/repo/bun"
	"github.com/jestradaramos/group-lift/backend/pkg/user"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	db := bun.InitBunDB("string")

	domain.RegisterUserServiceServer(grpcServer, user.NewUserServiceServer(db))
	grpcServer.Serve(lis)
}
