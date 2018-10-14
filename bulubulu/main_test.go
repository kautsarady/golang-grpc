package bulubulu

import (
	"errors"
	"testing"
)

func TestTranslate(t *testing.T) {
	tt := []struct {
		name, in, want string
		err            error
	}{
		{"simple1", "lblbl", "blblb", nil},
		{"simple2", "aiueo", "uuuuu", nil},
		{
			"complex",
			"once upon of a time live a king without an heir",
			"uncu upun uf u tumu buvu u kung wuthuut un huur",
			nil,
		},
		{"empty string", "", "", errors.New("Can't convert empty text")},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if res, _ := Translate(tc.in); res != tc.want {
				t.Errorf("Expected: %s; Got: %s", tc.want, res)
			}
		})
	}
}
