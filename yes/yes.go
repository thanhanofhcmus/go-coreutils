package main

import (
	"flag"
	"fmt"
)

const (
	DEFAULT = "yes"
	USAGE   = `yes - print a string until interrupted

FORM(S)
yes [REP]

Repeatedly print REP or '%s' by default
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE, DEFAULT)
		flag.PrintDefaults()
	}
	flag.Parse()

	rep := DEFAULT

	if args := flag.Args(); len(args) != 0 {
		rep = args[0]
	}

	for {
		fmt.Println(rep)
	}
}
