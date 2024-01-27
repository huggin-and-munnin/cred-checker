package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/hugin-and-munin/cred-checker/internal/app"
	pb "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/cred-checker"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := NewCredCheckerServer()

	pb.RegisterCredCheckerServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewCredCheckerServer() pb.CredCheckerServer {
	httpClient := &http.Client{}

	return app.NewCredChecker(httpClient)
}
