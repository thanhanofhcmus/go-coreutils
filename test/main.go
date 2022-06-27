package main

import (
	"fmt"
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

func printCString(b []byte) {
	fmt.Print(string(b[:clen(b)]))
}

func main() {
	u := unix.Utsname{}
	unix.Uname(&u)
	printCString(u.Sysname[:])
}
