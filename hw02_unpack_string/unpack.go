package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here.
	runeStr := []rune(s)
	var (
		resultString string
		n            int
	)

	for i, simbol := range runeStr {
		if unicode.IsDigit(simbol) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(simbol) && unicode.IsDigit(runeStr[i-1]) && runeStr[i-2] != '\\' {
			return "", ErrInvalidString
		}

		if i == (len(runeStr)-1) && simbol == '\\' {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(simbol) {
			n = int(simbol - '0')
			switch {
			case n == 0:
				strsimbol := string(runeStr[i-1])
				resultString = resultString[:len(resultString)-len(strsimbol)]
				continue
			case runeStr[i-1] == '\\' && runeStr[i-2] == '\\':
				resultString += strings.Repeat(string(runeStr[i-1]), n)
				continue
			case runeStr[i-1] == '\\' && runeStr[i-2] != '\\':
				resultString += string(simbol)
				continue
			default:
				resultString += strings.Repeat(string(runeStr[i-1]), n-1)
				continue
			}
		}

		if i == 0 || i > 0 && simbol != '\\' {
			resultString += string(simbol)
		}
	}
	return resultString, nil
}
