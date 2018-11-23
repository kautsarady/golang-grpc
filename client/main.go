package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/kautsarady/golang-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	server := flag.String("b", "localhost:8080", "Server address")
	text := flag.String("t", "", "Text to translate")
	flag.Parse()

	if *text == "" {
		log.Println("Text not specified, sending empty string to translate; use -t ... flag instead")
	}

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewTranslatorClient(conn)

	txt := &pb.Text{Text: *text}
	res, err := client.Translate(context.Background(), txt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Text)
}
