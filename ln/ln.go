package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	USAGE = `ln - create links

FORM(S)
ln [-s] SOURCE TARGET_FILE
ln [-s] SOURCE... TARGET_DIR

OPTION(S)
`
)

var (
	linkFunc func(oldname, newname string) error
)

func reportError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	optSymLink := flag.Bool("s", false, "Create symbolic link instead of hard link")
	flag.Parse()

	linkFunc = os.Link
	if *optSymLink {
		linkFunc = os.Symlink
	}

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Must have aleast two operant")
		os.Exit(1)
	}

	dest := args[len(args)-1]
	destStat, err := os.Stat(dest)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) && len(args) == 2 {
			if err := linkFunc(args[0], args[1]); err != nil {
				reportError(err)
			}
			return
		}
		reportError(err)
	}

	if !destStat.IsDir() {
		fmt.Printf("%s is not a directory\n", dest)
		os.Exit(1)
	}

	for i := 0; i < len(args)-1; i++ {
		file := args[i]
		path := filepath.Join(dest, file)
		if err := linkFunc(file, path); err != nil {
			reportError(err)
		}
	}
}
