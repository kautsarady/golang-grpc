package service_test

import (
	"context"
	"testing"

	pb "github.com/kautsarady/golang-grpc/api"
	"github.com/kautsarady/golang-grpc/service"
)

func TestTranslateService(t *testing.T) {
	s := service.Service{}
	req := &pb.Text{Text: "testing bulubulu"}

	resp, err := s.Translate(context.Background(), req)
	if err != nil {
		t.Errorf("%v", err)
	}

	expect := "tustung lubulubu"
	if resp.Text != expect {
		t.Errorf("Expecting \"%s\" got \"%s\"", expect, resp.Text)
	}
}
