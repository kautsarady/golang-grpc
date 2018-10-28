package bulubulu

import (
	"fmt"
	"log"
)

func ExampleTranslate() {
	text := "once upon of a time live a king without an heir"
	s, err := Translate(text)
	if err != nil {
		log.Fatalf("Error translating '%s': %v", text, err)
	}
	fmt.Println(s)
	// Output:
	// uncu upun uf u tumu buvu u kung wuthuut un huur
}
