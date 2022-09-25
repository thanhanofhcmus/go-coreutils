package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	USAGE = `wc - print newline, word, byte counts for files

FORM(S)
wc [OPTION...] FILE...

print in order of line - words - characters - bytes

if no OPTION, same as -clw

OPTION(S)
`
)

type stat struct {
	bytes, words, lines, chars int64
}

// insufficient but easy to write
func (s *stat) countFromLineBuffer(buffer []byte) {
	s.bytes += int64(len(buffer))
	s.chars += int64(utf8.RuneCount(buffer))
	s.lines += 1
	s.words += int64(len(strings.Fields(string(buffer))))
}

func newStatFromFile(f *os.File) stat {
	br := bufio.NewScanner(f)
	s := stat{}

	for br.Scan() {
		s.countFromLineBuffer(br.Bytes())
	}

	return s
}

var (
	totalStat stat
	argBytes  = flag.Bool("c", false, "Print the byte counts")
	argLines  = flag.Bool("l", false, "Print the line counts")
	argWords  = flag.Bool("w", false, "Print the word counts")
	argChars  = flag.Bool("m", false, "Print the character counts")
)

func addToTotal(s stat) {
	totalStat.bytes += s.bytes
	totalStat.chars += s.chars
	totalStat.words += s.words
	totalStat.lines += s.lines
}

func displayStat(s stat, filename string) {
	if !(*argBytes || *argWords || *argChars || *argLines) {
		fmt.Printf(" %6d %6d %6d %s", s.lines, s.words, s.bytes, filename)
	}

	var sb strings.Builder
	if *argLines {
		sb.WriteRune(' ')
		sb.WriteString(fmt.Sprintf("%6d", s.lines))
	}
	if *argWords {
		sb.WriteRune(' ')
		sb.WriteString(fmt.Sprintf("%6d", s.words))
	}
	if *argChars {
		sb.WriteRune(' ')
		sb.WriteString(fmt.Sprintf("%6d", s.chars))
	}
	if *argBytes {
		sb.WriteRune(' ')
		sb.WriteString(fmt.Sprintf("%6d", s.bytes))
	}

	fmt.Println(sb.String())
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
	}
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Must have at least one parameter")
	}
	counter := 0 // to check how many file checked

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error open file: %s\n", file)
			continue
		}
		s := newStatFromFile(f)
		counter += 1
		displayStat(s, file)
		addToTotal(s)
		f.Close()
	}

	if counter > 1 { // more than one file checked, print total
		displayStat(totalStat, "total")
	}
}
