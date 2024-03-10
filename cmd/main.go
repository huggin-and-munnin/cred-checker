package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/hugin-and-munin/cred-checker/internal/app"
	"github.com/hugin-and-munin/cred-checker/internal/config"
	pb "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/cred-checker"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	ServeHealthProbe()

	port := fmt.Sprintf(":%s", config.GetValue(config.Port))

	lis, err := net.Listen("tcp", port)
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

func ServeHealthProbe() {
	port := fmt.Sprintf(":%s", config.GetValue(config.HealthPort).String())
	go func() {
		log.Printf("Starting health check server on port %s", port)
		log.Printf("Health check path: %s", config.GetValue(config.HealthPath).String())
		http.ListenAndServe(port, nil)
	}()

	log.Printf("health probe listening at port %s", port)
}
