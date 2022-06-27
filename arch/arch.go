package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func clen(b []byte) int {
	for i, v := range b {
		if v == 0 {
			return i
		}
	}
	return len(b)
}

func cString(b []byte) string {
	return string(b[:clen(b)])
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "arch - print machine hardware name (same as uname -m)\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if args := flag.Args(); len(args) != 0 {
		fmt.Println("too many arguments")
		os.Exit(1)
	}

	utsname := &unix.Utsname{}
	err := unix.Uname(utsname)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(cString(utsname.Machine[:]))

}
