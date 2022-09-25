package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	USAGE = `factor - factor numbers

FORM(S)
factor NUMBER...
`
)

func getFactors(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}

	var factors []int

	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for f := 3; f <= n; f += 2 {
		if n%f == 0 {
			factors = append(factors, f)
			n /= f
		} else if f*f == n {
			factors = append(factors, f)
			factors = append(factors, f)
			break
		} else if f*f > n {
			factors = append(factors, n)
			break
		}
	}

	return factors
}

func printFactors(n int) {
	factors := getFactors(n)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d: ", n))
	for _, f := range factors {
		sb.WriteString(fmt.Sprintf("%d ", f))
	}
	fmt.Println(sb.String())
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Must have at least one parameter")
		os.Exit(1)
	}

	for _, arg := range args {
		if n, err := strconv.Atoi(arg); err == nil {
			printFactors(n)
		} else {
			fmt.Printf("%s is not a non negative number", arg)
		}
	}
}
