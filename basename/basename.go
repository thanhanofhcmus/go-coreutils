package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	SEPARATOR = '/'
	USAGE     = `basename - strip directory and suffix from a file name

FORM(S)
basename NAME [SUFFIX]
basename [OPTION...] NAME...

OPTION(S)
`
)

func getTokenPositionReverse(s string, suffix rune) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(suffix) {
			return i
		}
	}
	return -1
}

func stripBase(s string) string {
	if pos := getTokenPositionReverse(s, SEPARATOR); pos >= 0 {
		return s[pos+1:]
	}
	return s
}

func main() {
	flag.Usage = func() {
		if _, err := fmt.Fprintf(flag.CommandLine.Output(), USAGE); err != nil {
			log.Fatal(err)
		}
		flag.PrintDefaults()
	}
	argMultiple := flag.Bool("a", false, "Support multiple input as NAME")
	argSuffix := flag.String("s", "", "Take SUFFIX as an argument, implies -a")
	argZero := flag.Bool("z", false, "End each output line with a NULL instead of a newline")
	flag.Parse()

	names := flag.Args()
	multiple := *argMultiple
	suffix := ""

	if len(names) == 0 {
		fmt.Println("Must have at least one parameter")
		os.Exit(1)
	}

	if len(*argSuffix) != 0 {
		multiple = true
		suffix = *argSuffix
	}

	if !multiple && len(names) == 2 {
		suffix = names[1]
		names = names[:len(names)-1]
	}

	var sb strings.Builder
	sep := "\n"
	if *argZero {
		sep = ""
	}

	for _, name := range names {
		sb.WriteString(strings.TrimSuffix(stripBase(name), suffix))
		sb.WriteString(sep)
	}

	res := sb.String()
	if sep == "\n" {
		res = res[:len(res)-1] // remove the trailing newline
	}
	fmt.Println(res)
}
