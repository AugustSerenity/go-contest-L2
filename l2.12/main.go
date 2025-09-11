package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type config struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
	filenames  []string
}

func main() {
	cfg := parseFlags()

	var regex *regexp.Regexp
	var err error
	pattern := cfg.pattern

	if cfg.fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	if cfg.ignoreCase {
		pattern = "(?i)" + pattern
	}

	regex, err = regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling regex: %v\n", err)
		os.Exit(1)
	}

	if cfg.context > 0 {
		cfg.after = cfg.context
		cfg.before = cfg.context
	}

	if len(cfg.filenames) == 0 {
		processInput(os.Stdin, "", regex, cfg)
	} else {
		for _, filename := range cfg.filenames {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
				continue
			}
			processInput(file, filename, regex, cfg)
			file.Close()
		}
	}
}

func parseFlags() config {
	after := flag.Int("A", 0, "print N lines after match")
	before := flag.Int("B", 0, "print N lines before match")
	context := flag.Int("C", 0, "print N lines of context around match")
	count := flag.Bool("c", false, "print only count of matching lines")
	ignoreCase := flag.Bool("i", false, "ignore case")
	invert := flag.Bool("v", false, "select non-matching lines")
	fixed := flag.Bool("F", false, "fixed string match (not regex)")
	lineNum := flag.Bool("n", false, "print line number with output")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: grep [OPTIONS] PATTERN [FILE...]")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	filenames := flag.Args()[1:]

	return config{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
		pattern:    pattern,
		filenames:  filenames,
	}
}

func processInput(file *os.File, filename string, regex *regexp.Regexp, cfg config) {
	scanner := bufio.NewScanner(file)
	var lines []string
	var lineNumbers []int

	lineNumber := 1
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineNumbers = append(lineNumbers, lineNumber)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	var matches []int
	for i, line := range lines {
		matched := regex.MatchString(line)
		if (cfg.invert && !matched) || (!cfg.invert && matched) {
			matches = append(matches, i)
		}
	}

	if cfg.count {
		if len(cfg.filenames) > 1 {
			fmt.Printf("%s:%d\n", filename, len(matches))
		} else {
			fmt.Printf("%d\n", len(matches))
		}
		return
	}

	printed := make(map[int]bool)
	for _, matchIdx := range matches {
		printMatchWithContext(matchIdx, lines, lineNumbers, cfg, printed, filename)
	}
}

func printMatchWithContext(matchIdx int, lines []string, lineNumbers []int, cfg config, printed map[int]bool, filename string) {
	start := max(0, matchIdx-cfg.before)
	end := min(len(lines)-1, matchIdx+cfg.after)

	if len(printed) > 0 && (cfg.before > 0 || cfg.after > 0) {
		lastPrinted := -1
		for i := range printed {
			if i > lastPrinted {
				lastPrinted = i
			}
		}
		if matchIdx-cfg.before > lastPrinted+1 {
			fmt.Println("--")
		}
	}

	for i := start; i <= end; i++ {
		if printed[i] {
			continue
		}
		printed[i] = true

		var output strings.Builder

		if len(cfg.filenames) > 1 {
			output.WriteString(filename)
			output.WriteString(":")
		}

		if cfg.lineNum {
			output.WriteString(fmt.Sprintf("%d:", lineNumbers[i]))
		}

		output.WriteString(lines[i])

		fmt.Println(output.String())
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
