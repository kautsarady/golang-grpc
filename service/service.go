package server

import (
	"fmt"

	pb "github.com/kautsarady/golang-grpc/api"
	"github.com/kautsarady/golang-grpc/bulubulu"
	"golang.org/x/net/context"
)

// Service handler
type Service struct{}

// Translate call grpc service
func (Service) Translate(ctx context.Context, txt *pb.Text) (*pb.Text, error) {
	res, err := bulubulu.Translate(txt.GetText())
	if err != nil {
		return nil, fmt.Errorf("Error translating %s : %v", txt.GetText(), err)
	}

	return &pb.Text{Text: res}, nil
}
