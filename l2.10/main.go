package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func transformation(lines []string, column int, numeric, revers, unique bool) []string {
	sort.SliceStable(lines, func(i, j int) bool {
		fieldsI := strings.Split(lines[i], "\t")
		fieldsJ := strings.Split(lines[j], "\t")

		var a, b string
		if column < len(fieldsI) {
			a = fieldsI[column]
		}
		if column < len(fieldsJ) {
			b = fieldsJ[column]
		}

		if numeric {
			numA, errA := strconv.ParseFloat(a, 64)
			numB, errB := strconv.ParseFloat(b, 64)

			if errA == nil && errB == nil {
				if revers {
					return numA > numB
				}
				return numA < numB
			}
		}

		if revers {
			return a > b
		}
		return a < b
	})

	if unique {
		lines = uniqueSort(lines)
	}

	return lines
}

func uniqueSort(lines []string) []string {
	m := make(map[string]bool, cap(lines))
	var uniqueLines []string

	for _, val := range lines {
		if !m[val] {
			m[val] = true
			uniqueLines = append(uniqueLines, val)
		}
	}

	return uniqueLines
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Не указано имя файла! Укажите имя файла.")
	}

	fileName := os.Args[1]

	var column int
	numericSort := false
	reversLines := false
	uniqueLines := false

	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "-k" && i+1 < len(os.Args) {
			col, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				log.Fatal("Не верный номер столбца: ", os.Args[i+1])
			}
			column = col - 1
			i++
		} else if os.Args[i] == "-n" {
			numericSort = true
		} else if os.Args[i] == "-r" {
			reversLines = true
		} else if os.Args[i] == "-u" {
			uniqueLines = true
		}
	}

	//fmt.Println(len(os.Args), os.Args)
	str, err := readFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	str = transformation(str, column, numericSort, reversLines, uniqueLines)
	for _, val := range str {
		fmt.Println(val)
	}

}
