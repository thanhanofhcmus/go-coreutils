package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "true - return true value\n")
	}
	flag.Parse()

	os.Exit(0)
}
