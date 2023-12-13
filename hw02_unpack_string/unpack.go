package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	// Place your code here.

	runeStr := []rune(s)

	var (
		resultString string

		n int
	)

	for i, simbol := range runeStr {

		// первый символ цифра

		if unicode.IsDigit(simbol) && i == 0 {

			return "", ErrInvalidString

		}

		// две цифры подряд и нет обратного слэша

		if unicode.IsDigit(simbol) && unicode.IsDigit(runeStr[i-1]) && runeStr[i-2] != '\\' {

			return "", ErrInvalidString

		}

		// заполняем итоговую строку

		if unicode.IsDigit(simbol) {

			n = int(simbol - '0')

			// если n == 0 то убираем последний символ из строки

			switch {

			case n == 0:

				resultString = resultString[:len(resultString)-1]

				continue

			case runeStr[i-1] == '\\' && runeStr[i-2] == '\\':

				for n > 0 {

					resultString += string(runeStr[i-1])

					n--

				}

				continue

			case runeStr[i-1] == '\\' && runeStr[i-2] != '\\':

				resultString += string(simbol)

				continue

			default:

				for n > 1 {

					resultString += string(runeStr[i-1])

					n--

				}

				continue

			}

		}

		if i == 0 || i > 0 && simbol != '\\' {

			resultString += string(simbol)

		}

	}

	return resultString, nil

}

import (
	"errors"
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
		// первый символ цифра
		if unicode.IsDigit(simbol) && i == 0 {
			return "", ErrInvalidString
		}

		// две цифры подряд и нет обратного слэша
		if unicode.IsDigit(simbol) && unicode.IsDigit(runeStr[i-1]) && runeStr[i-2] != '\\' {
			return "", ErrInvalidString
		}

		// заполняем итоговую строку
		if unicode.IsDigit(simbol) {
			n = int(simbol - '0')
			// если n == 0 то убираем последний символ из строки
			switch {
			case n == 0:
				resultString = resultString[:len(resultString)-1]
				continue
			case runeStr[i-1] == '\\' && runeStr[i-2] == '\\':
				for n > 0 {
					resultString += string(runeStr[i-1])
					n--
				}
				continue
			case runeStr[i-1] == '\\' && runeStr[i-2] != '\\':
				resultString += string(simbol)
				continue
			default:
				for n > 1 {
					resultString += string(runeStr[i-1])
					n--
				}
				continue
			}
		}

		if i == 0 || i > 0 && simbol != '\\' {
			resultString += string(simbol)
		}
	}
	return resultString, nil
}
