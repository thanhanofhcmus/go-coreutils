package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "false - return false value\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	
	os.Exit(1)
}
