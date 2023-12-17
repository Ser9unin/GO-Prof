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
		resultString, prevSymbolString string
		n, backslashCount              int
		prevSymbolDigit                bool
	)

	for i, symbol := range runeStr {
		switch {
		case symbol == '\\' && i == runeStrLen:
			return "", ErrInvalidString
		case symbol == '\\':
			backslashCount++
			fallthrough
		case i > 0:
			prevSymbolString = string(runeStr[i-1])
			prevSymbolDigit = unicode.IsDigit(runeStr[i-1])
		}

		if unicode.IsDigit(symbol) {
			n = int(symbol - '0')
			switch {
			case i == 0:
				return "", ErrInvalidString
			case prevSymbolDigit && runeStr[i-2] != '\\':
				return "", ErrInvalidString
			case prevSymbolDigit && backslashCount%2 == 0 && backslashCount > 0:
				return "", ErrInvalidString
			case backslashCount%2 == 0 && backslashCount > 0:
				resultString += strings.Repeat(`\`, n)
				continue
			case n == 0:
				resultString = strings.TrimSuffix(resultString, prevSymbolString)
				continue
			case backslashCount%2 != 0:
				if backslashCount > 1 {
					resultString += `\`
				}
				resultString += string(symbol)
				backslashCount = 0
				continue
			default:
				resultString += strings.Repeat(prevSymbolString, n-1)
				continue
			}
		}

		if i == 0 || i > 0 && symbol != '\\' {
			resultString += string(symbol)
		}
	}
	return resultString, nil
}
