package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jestradaramos/group-lift/backend/pkg/domain"
	"github.com/jestradaramos/group-lift/backend/pkg/repo/bun"
	"github.com/jestradaramos/group-lift/backend/pkg/session"
	"github.com/jestradaramos/group-lift/backend/pkg/user"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:81")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	// Default postgres creds for now
	db := bun.InitBunDB("postgresql://postgres:postgres@localhost:5432/test?sslmode=disable")
	// db.CreateLiftTables(context.Background())

	domain.RegisterUserServiceServer(grpcServer, user.NewUserServiceServer(db))
	domain.RegisterLiftSessionServiceServer(grpcServer, session.NewSessionServiceServer(db))
	fmt.Print(grpcServer.GetServiceInfo())

	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Print("It Brokey no Workey")
		grpcServer.Stop()
	}
}
