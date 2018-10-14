package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kautsarady/golang-grpc/bulubulu"

	"golang.org/x/net/context"

	pb "github.com/kautsarady/golang-grpc/server/api"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "Port to listen to")
	flag.Parse()

	log.Printf("Listening to port %d\n", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Error listening port %d : %v\n", *port, err)
	}

	srv := grpc.NewServer()
	pb.RegisterTranslatorServer(srv, service{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Error serving : %v\n", err)
	}
}

type service struct{}

func (service) Translate(ctx context.Context, txt *pb.Text) (*pb.Text, error) {
	res, err := bulubulu.Translate(txt.GetText())
	if err != nil {
		return nil, fmt.Errorf("Error translating %s : %v", txt.GetText(), err)
	}

	return &pb.Text{Text: res}, nil
}
