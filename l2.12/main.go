package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	context := flag.Int("C", 0, "Выводит N строк до и после найденной строки")
	count := flag.Bool("c", false, "Считает количество найденных строк")
	ignoreCase := flag.Bool("i", false, "Игнорирует регистр")
	invert := flag.Bool("v", false, "выводит строки, не содержащие шаблон")
	fixed := flag.Bool("F", false, "Фиксированная строка (не шаблон) - точный поиск")
	lineNum := flag.Bool("n", false, "Выводит номер строки перед каждой найденной строкой")
	flag.Parse()

	pattern := flag.Arg(0)

	var regex *regexp.Regexp
	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if *ignoreCase {
		regex = regexp.MustCompile("(?i)" + pattern)
	} else {
		regex = regexp.MustCompile(pattern)
	}

	var file *os.File
	if flag.NArg() > 1 {
		var err error
		file, err = os.Open(flag.Arg(1))
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	scanner := bufio.NewScanner(file)
	var matchingLines int
	var outputLines []string
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		matched := regex.MatchString(line)
		if (*invert && !matched) || (!*invert && matched) {
			if *count {
				matchingLines++
			} else {
				if *lineNum {
					line = fmt.Sprintf("%d:%s", lineNumber, line)
				}
				outputLines = append(outputLines, line)
				if *context > 0 {
					outputLines = append(outputLines, getContextLines(scanner, *context)...)
				}
			}
		}
	}

	if *count {
		fmt.Println(matchingLines)
	} else {
		for _, line := range outputLines {
			fmt.Println(line)
		}
	}
}

func getContextLines(scanner *bufio.Scanner, context int) []string {
	var lines []string
	for i := 0; i < context && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}
	return lines
}
