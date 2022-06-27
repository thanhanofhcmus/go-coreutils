package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	SEPERATOR = '/'
	USAGE     = `dirname - strip last componment of a file

FORM(S)
dirname [OPTION...] NAME...

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

func stripLast(s string) string {
	if pos := getTokenPositionReverse(s, SEPERATOR); pos >= 0 {
		return s[:pos]
	}
	return s
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	zeroArg := flag.Bool("z", false, "End each output line with a NULL instead of a newline")
	flag.Parse()

	names := flag.Args()

	if len(names) == 0 {
		fmt.Println("must have aleast one operant")
		os.Exit(1)
	}

	sep := "\n"
	if *zeroArg {
		sep = ""
	}

	var sb strings.Builder
	for _, name := range names {
		sb.WriteString(stripLast(name))
		sb.WriteString(sep)
	}

	res := sb.String()
	if sep == "\n" {
		res = res[:len(res)-1]  // remove trailing newline
	}
	
	fmt.Println(res)

}
