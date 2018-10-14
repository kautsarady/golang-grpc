package bulubulu_test

import (
	"fmt"
	"log"

	"github.com/kautsarady/golang-grpc/bulubulu"
)

func ExampleTranslate() {
	text := "once upon of a time live a king without an heir"
	s, err := bulubulu.Translate(text)
	if err != nil {
		log.Fatalf("Error translating '%s': %v", text, err)
	}
	fmt.Println(s)
	// Output:
	// uncu upun uf u tumu buvu u kung wuthuut un huur
}
