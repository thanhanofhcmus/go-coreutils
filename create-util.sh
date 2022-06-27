#! /bin/sh

echo $1

mkdir $1
cd $1
go mod init "an/go-coreutils-$1"
touch "$1.go"
cd ..
go work use "./$1"
