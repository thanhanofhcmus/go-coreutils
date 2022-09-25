package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

const USAGE = `uname - print system information

FORM(S)
uname [OPTION...]

if no OPTION, same as -s

OPTION(S)
`

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
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}

	argOsName := flag.Bool("s", false, "Print the operating system name")
	argAll := flag.Bool("a", false, "Print all information (as if running with mnrsv), nullify other options")
	argOsRelease := flag.Bool("r", false, "Print the operating system release")
	argOsVersion := flag.Bool("v", false, "Print the operating system version")
	argNodeName := flag.Bool("n", false, "Print the network nodename")
	argMachineName := flag.Bool("m", false, "Print the machine name")

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

	osName := cString(utsname.Sysname[:])
	osRelease := cString(utsname.Release[:])
	osVersion := cString(utsname.Version[:])
	nodeName := cString(utsname.Nodename[:])
	machineName := cString(utsname.Machine[:])

	if *argAll {
		fmt.Printf("%s %s %s %s %s\n", machineName, nodeName, osRelease, osName, osVersion)
		os.Exit(0)
	}

	var sb strings.Builder

	if *argOsName {
		sb.WriteString(osName)
		sb.WriteRune(' ')
	}
	if *argNodeName {
		sb.WriteString(nodeName)
		sb.WriteRune(' ')
	}
	if *argOsRelease {
		sb.WriteString(osRelease)
		sb.WriteRune(' ')
	}
	if *argOsVersion {
		sb.WriteString(osVersion)
		sb.WriteRune(' ')
	}
	if *argMachineName {
		sb.WriteString(machineName)
		sb.WriteRune(' ')
	}

	line := sb.String()

	if len(line) == 0 {
		line = osName
	} else {
		line = line[:len(line)-1] // remove trailing whitespace
	}

	fmt.Println(line)
}
