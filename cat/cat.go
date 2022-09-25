package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	USAGE = `cat - concatenate and print files

FORM(S)
cat [OPTION...] FILE...

OPTION(S)
`
)

var (
	argNonBlank = flag.Bool("b", false, "Number non-empty output line, suppress -n")
	argNumber   = flag.Bool("n", false, "Number all output line")
	argSqueeze  = flag.Bool("s", false, "Squeeze multiple adjacent empty line to a single one")

	simpleCat = io.Copy
)

func complexCat(w io.Writer, r io.Reader) (n int64, err error) {
	var line, lastLine string
	lineNum := 0
	br := bufio.NewReader(r)

	for {
		line, err = br.ReadString('\n')
		if err != nil {
			return
		}

		if *argSqueeze && lastLine == "\n" && line == "\n" {
			continue
		}
		if *argNonBlank {
			lineNum++
			fmt.Fprintf(w, "%5d\t%s", lineNum, line)
		} else if *argNumber && line != "\n" {
			lineNum++
			fmt.Fprintf(w, "%5d\t%s", lineNum, line)
		} else {
			fmt.Fprint(w, line)
		}
		lastLine = line
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	files := flag.Args()

	cat := simpleCat
	if *argNumber || *argNonBlank || *argSqueeze {
		cat = complexCat
	}

	if len(files) == 0 {
		cat(os.Stdout, os.Stdin)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("Error open file: %s", file)
				continue
			}
			cat(os.Stdout, f)
			f.Close()
		}
	}
}
