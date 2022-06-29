#!/bin/sh

if [ -d $1 ]; then
	echo "$1 already exists"
	exit 1
fi

mkdir $1
cd $1
go mod init "an/go-coreutils-$1"
echo "package main

import (
	\"flag\"
	\"fmt\"
	\"os\"
)

const (
	USAGE = \`$1 - 

FORM(S)
$1

OPTION(S)
\`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(\"Must have aleast one operant\")
		os.Exit(1)
	}
}
" >> "$1.go"
cd ..
go work use "./$1"
