package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	USAGE = `printenv - print all ore same environment variables

FORM(S)
printenv [NAME]

if NAME is specified, print the variables of that name, or else print all
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		envs := os.Environ()
		for _, e := range envs {
			fmt.Println(e)
		}
	} else {
		if e, present := os.LookupEnv(args[0]); !present {
			os.Exit(1)
		} else {
			fmt.Println(e)
		}
	}
}
