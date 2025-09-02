package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func main() {
	str, err := ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
}
