package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/hugin-and-munin/cred-checker/internal/app/cred_checker"
	"github.com/hugin-and-munin/cred-checker/internal/app/health"
	"github.com/hugin-and-munin/cred-checker/internal/config"
	cred_checker_pb "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/cred-checker"
	health_pb "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/health"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	port := fmt.Sprintf(":%s", config.GetValue(config.Port))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	credCheckerServer := NewCredCheckerServer()
	healthProbe := NewHealthProbe()

	health_pb.RegisterHealthServer(s, healthProbe)
	cred_checker_pb.RegisterCredCheckerServer(s, credCheckerServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewCredCheckerServer() cred_checker_pb.CredCheckerServer {
	httpClient := &http.Client{}

	return cred_checker.NewCredChecker(httpClient)
}

func NewHealthProbe() health_pb.HealthServer {
	return health.NewHealthProbe()
}
