package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	USAGE = `mkdir - create directories

FORM(S)
mkdir [-p] PATH...

OPTION(S)
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	argParent := flag.Bool("p", false, "Create parent directory as needed")
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		fmt.Println("Must have aleast one operant")
		os.Exit(1)
	}

	for _, path := range paths {
		if *argParent {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			if err := os.Mkdir(path, os.ModePerm); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

