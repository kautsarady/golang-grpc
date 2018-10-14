package bulubulu

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	if res, _ := Translate(""); res != "Helo" {
		t.Errorf("not expected output\n")
	}
}
