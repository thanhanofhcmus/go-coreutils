package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	_ = flag.Bool("L", false, "Print logical path")
	physical := flag.Bool("P", false, "Print physical path")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `pwd - print current working directory
if there are no arguments given print logical path
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	if opts := flag.Args(); len(opts) != 0 {
		fmt.Println("too many arguments")
		os.Exit(1)
	}

	pwd, err := os.Getwd();
	if err != nil {
		os.Exit(1)
	}

	if *physical {
		if realPath, err := filepath.EvalSymlinks(pwd); err != nil {
			os.Exit(1)
		} else {
			pwd = realPath
		}
	}

	fmt.Println(pwd)
	
}
