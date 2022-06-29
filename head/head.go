package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	DEFAUTL_COUNT int = 10

	USAGE = `head - output the first part of files

FORM(S)
head [-n NUMBER | -c NUMBER]  [FILE...]

if there are no OPTIONS, same as -n %d

if there are no FILES, read from STDIN

if there are more than one FILES, each file will have a header in a form of ==> FILE <==

OPTION(S)
`
)

var (
	files []string

	optBytes      *int = flag.Int("c", DEFAUTL_COUNT, "Print the first NUMBER bytes of each file, suppress -n")
	optLines      *int = flag.Int("n", DEFAUTL_COUNT, "Print the first NUMBER lines of each file")
	isOptBytesSet bool = false
)

func headBytes(r io.Reader, n int) {
	if n < 0 {
		fmt.Println("Illegal number, non negative number only")
	}
	lr := io.LimitReader(r, int64(n))
	io.Copy(os.Stdout, lr)
}

func headLines(r io.Reader, n int) {
	count := 0
	scanner := bufio.NewScanner(r)
	for count < n && scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))
		count++
	}
}

func head(r io.Reader, i int) {
	if isOptBytesSet {
		headBytes(r, *optBytes)
		// do not print extra endline after the last file
		if i != len(files)-1 {
			fmt.Println()
		}
	} else {
		headLines(r, *optLines)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE, DEFAUTL_COUNT)
		flag.PrintDefaults()
	}
	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "c" {
			isOptBytesSet = true
		}
	})

	files = flag.Args()

	if len(files) == 0 {
		// passing number != -1 as the second argument so the program
		// does not print extra endline after it ends
		head(os.Stdin, 0)
	} else {
		for i := 0; i < len(files); i++ {
			file := files[i]
			if f, err := os.Open(file); err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Printf("==> %s <==\n", file)
				head(f, i)
			}
		}
	}
}
