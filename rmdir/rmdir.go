package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	USAGE = `rmdir - remove empty directories

FORM(S)
rmdir [-p] PATH...

OPTION(S)
`
)

func getAllPaths(path string) []string {
	parts := strings.Split(path, "/")
	v := parts[0]
	for i := 1; i < len(parts); i++ {
		v += "/" + parts[i]
		parts[i] = v
	}
	return parts
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	argParent := flag.Bool("p", false, "Remove parents if they are empty")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Must have aleast one operant")
		os.Exit(1)
	}

	for _, arg := range args {
		if *argParent {
			paths := getAllPaths(arg)
			for i := len(paths) - 1; i >= 0; i-- {
				if err := os.Remove(paths[i]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else {
			if err := os.Remove(arg); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
