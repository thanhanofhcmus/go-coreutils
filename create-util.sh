#!/bin/sh

echo $1

mkdir $1
cd $1
go mod init "an/go-coreutils-$1"
echo "package main

import (
	\"flag\"
	\"fmt\"
)

const (
	USAGE = \`$1 - \`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()
}
" >> "$1.go"
cd ..
go work use "./$1"
