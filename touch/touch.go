package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	USAGE = `touch - change file timestamp

FORM(S)
touch [OPTION...] FILE...

OPTION(S)
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	argNoCreate := flag.Bool("c", false, "Do not create file if not exists")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Must have at least one parameter")
		os.Exit(1)
	}

	for _, file := range files {
		_, err := os.Stat(file)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			if *argNoCreate {
				continue
			}
			f, err := os.Create(file)
			if err != nil {
				fmt.Println("Cannot create ", file)
			}
			f.Close()
		} else {
			now := time.Now()
			os.Chtimes(file, now, now)
		}
	}
}
