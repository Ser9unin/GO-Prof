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
	runeStrLen := len(runeStr) - 1

	var (
		resultString      string
		n, backslashCount int
	)

	for i, symbol := range runeStr {
		switch {
		case symbol == '\\' && i == runeStrLen:
			return "", ErrInvalidString
		case symbol == '\\':
			backslashCount++
		}

		if unicode.IsDigit(symbol) {
			n = int(symbol - '0')
			switch {
			case i == 0 || (unicode.IsDigit(runeStr[i-1]) && runeStr[i-2] != '\\'):
				return "", ErrInvalidString
			case backslashCount%2 == 0 && backslashCount > 0:
				switch {
				case unicode.IsDigit(runeStr[i-1]):
					return "", ErrInvalidString
				default:
					resultString += strings.Repeat(`\`, n)
					continue
				}
			case n == 0:
				symbolLen := len(string(runeStr[i-1]))
				myStrLen := len(resultString)
				resultString = resultString[:myStrLen-symbolLen]
				continue
			case backslashCount%2 != 0:
				if backslashCount > 1 {
					resultString += `\`
				}
				resultString += string(symbol)
				backslashCount = 0
				continue
			default:
				resultString += strings.Repeat(string(runeStr[i-1]), n-1)
				continue
			}
		}

		if i == 0 || i > 0 && symbol != '\\' {
			resultString += string(symbol)
		}
	}
	return resultString, nil
}
