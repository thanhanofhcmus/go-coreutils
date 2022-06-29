package main

import (
	"flag"
	"fmt"
	"os/user"
)

const USAGE = `whoami - print effective user id`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	if user, err := user.Current(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Username)
	}
}
