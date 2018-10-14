package bulubulu

import (
	"errors"
	"strings"
)

//Translate translate bulubulu language to human language
func Translate(txt string) (string, error) {
	if len(txt) == 0 {
		return "", errors.New("Can't convert empty text")
	}
	buf := ""
	for i := 0; i < len(txt); i++ {
		ch := strings.ToLower(txt[i : i+1])
		switch ch {
		case "l":
			buf += "b"
			break
		case "b":
			buf += "l"
			break
		default:
			buf += repVowel(ch)
		}
	}
	return buf, nil
}

func repVowel(s string) string {
	for _, v := range []string{"a", "i", "u", "e", "o"} {
		if s == v {
			return "u"
		}
	}
	return s
}
