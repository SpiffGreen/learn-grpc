package main

import (
	"log"
	"net"

	pb "github.com/spiffgreen/learn-grpc/learn"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	addr string
}

func NewGrpcServer(addr string) *GrpcServer {
	return &GrpcServer{addr: addr}
}

func (g *GrpcServer) Run() error {
	lis, err := net.Listen("tcp", g.addr)

	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterHelloAppServer(s, &HelloApp{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	grpcServer := NewGrpcServer(":9000")
	grpcServer.Run()
}
