package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("некорректная строка, т.к. в строке только цифры")

func Unpacking(s string) (string, error) {
	var res string
	var char byte
	var chekDigit int

	count := -1

	if s == "" {
		return "", nil
	}

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			chekDigit++
		}
	}

	if chekDigit == len(s) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			countStr := ""
			for unicode.IsDigit(rune(s[i])) {
				countStr += string(s[i])
				i++
				if i == len(s) {
					break
				}
			}
			i--
			count, _ = strconv.Atoi(countStr)
			res += strings.Repeat(string(char), count-1)
		} else {
			res += string(s[i])
			char = s[i]
		}
	}
	return res, nil
}

func main() {
	var s string
	fmt.Scan(&s)

	unpacked, err := Unpacking(s)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println(unpacked)

}
