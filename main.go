package main

import (
	"fmt"
	"log"

	"github.com/kautsarady/golang-grpc/bulubulu"
)

func main() {
	text := "sometimes i wonder why the sky is full of star"
	res, err := bulubulu.Translate(text)
	if err != nil {
		log.Fatalf("Error translating '%s': %v", text, err)
	}
	fmt.Println(res)
}
