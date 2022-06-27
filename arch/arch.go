package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("arch: ", runtime.GOARCH)
	fmt.Println("os: ", runtime.GOOS)
}
