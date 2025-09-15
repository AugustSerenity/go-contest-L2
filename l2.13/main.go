package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "fields to select")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "only lines with delimiter")
	flag.Parse()

	fieldSet := make(map[int]bool)
	if *fields != "" {
		for _, part := range strings.Split(*fields, ",") {
			if strings.Contains(part, "-") {
				rangeParts := strings.Split(part, "-")
				if len(rangeParts) == 2 {
					start := atoi(rangeParts[0])
					end := atoi(rangeParts[1])
					if start > 0 && end >= start {
						for i := start; i <= end; i++ {
							fieldSet[i] = true
						}
					}
				}
			} else {
				fieldNum := atoi(part)
				if fieldNum > 0 {
					fieldSet[fieldNum] = true
				}
			}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		parts := strings.Split(line, *delimiter)

		if len(fieldSet) == 0 {
			fmt.Println(line)
			continue
		}

		var selected []string
		for i := 0; i < len(parts); i++ {
			if fieldSet[i+1] {
				selected = append(selected, parts[i])
			}
		}

		if len(selected) > 0 {
			fmt.Println(strings.Join(selected, *delimiter))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "erorr read:", err)
		os.Exit(1)
	}
}

func atoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
