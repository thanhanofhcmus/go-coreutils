package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	USAGE = `link - make a hard link to a file

FORMS
link FILE LINK

create a link named LINK to an existing file named FILE
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Must have exactly two parameters")
		os.Exit(1)
	}

	if err := os.Link(args[0], args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
