package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/kautsarady/golang-grpc/api"
	service "github.com/kautsarady/golang-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "Port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Error listening port %d : %v\n", *port, err)
	}

	srv := grpc.NewServer()
	pb.RegisterTranslatorServer(srv, service.Service{})

	log.Printf("Listening to port %d\n", *port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Error serving : %v\n", err)
	}
}
