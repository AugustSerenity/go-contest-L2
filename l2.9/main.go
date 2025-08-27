package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func unpacking(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	r := []rune(s)

	var res string

	for i := 1; i < len(r); i++ {

		if unicode.IsNumber(r[i]) && unicode.IsNumber(r[i-1]) {
			continue
		}

		if unicode.IsNumber(r[i]) && unicode.IsLetter(r[i-1]) {
			count, _ := strconv.Atoi(string(r[i]))

			for ; count != 0; count-- {
				res += string(r[i-1])
			}

		}

		if !unicode.IsNumber(r[i]) {
			res += string(r[i])
		}

	}

	return res, nil
}

func main() {
	var s string
	fmt.Scan(&s)

	unpacked, err := unpacking(s)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println(unpacked)

}
