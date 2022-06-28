package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	USAGE = `echo - write argument to the output

FORM(S)
echo [OPTION...] [STRING...]

OPTION(S)
`
	ESCAPES = `
If -e is in effect, the following escape is recognized:

	\\ backslash

	\a alert (BEL)

	\b backspace

	\c procedure no futher input

	\e escape

	\f form feed

	\n line feed

	\r carriage return

	\t horizontal tab

	\v vertical tab

	\xHH byte with hexadecimal value of HH (1 or 2 digits)

	\0NNN byte with octal value of NNN (1 to 3 digits)
`
)

func isOctalByte(r byte) bool {
	return '0' <= r && r <= '7'
}

func octToBin(r byte) byte {
	return r - '0'
}

func isHexaByte(r byte) bool {
	return ('0' <= r && r <= '9') || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
}

func hexToBin(r byte) byte {
	if '0' <= r && r <= '9' {
		return r - '0'
	} else if 'a' <= r && r <= 'f' {
		return r - 'a' + 10
	} else {
		return r - 'A' + 10
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE)
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), ESCAPES)
	}
	argNoNewline := flag.Bool("n", false, "Do not end the output line with a newline")
	argEscape := flag.Bool("e", false, "Enable backslash escape, overrides -E")
	_ = flag.Bool("E", false, "Disable backslash escape (default)")
	flag.Parse()

	concated := strings.Join(flag.Args(), " ")

	if *argEscape {
		idx := 0
		for idx < len(concated) {
			c := concated[idx]
			if c == '\\' && idx != len(concated)-1 {
				idx++
				switch concated[idx] {
				case '\\':
					c = '\\'
				case 'a':
					c = '\a'
				case 'b':
					c = '\b'
				case 'c':
					os.Exit(0)
				case 'e':
					c = '\x1B'
				case 'f':
					c = '\f'
				case 'n':
					c = '\n'
				case 'r':
					c = '\r'
				case 't':
					c = '\t'
				case 'v':
					c = '\v'
				case 'x':
					c = 0
					// check for the first hex char
					if idx < len(concated)-1 && isHexaByte(concated[idx+1]) {
						idx++
						c = hexToBin(concated[idx])
					} else {
						goto write
					}
					// check for the second hex char
					if idx < len(concated)-1 && isHexaByte(concated[idx+1]) {
						idx++
						c = 16*c + hexToBin(concated[idx])
					}
				case '0':
					c = 0
					// check for the first oct char
					if idx < len(concated)-1 && isOctalByte(concated[idx+1]) {
						idx++
						c = octToBin(concated[idx])
					} else {
						goto write
					}
					// check for the second oct char
					if idx < len(concated)-1 && isOctalByte(concated[idx+1]) {
						idx++
						c = 8*c + octToBin(concated[idx])
					} else {
						goto write
					}
					// check for the third oct char
					if idx < len(concated)-1 && isOctalByte(concated[idx+1]) {
						idx++
						c = 8*c + octToBin(concated[idx])
					}
				}
			}
		write:
			os.Stdout.Write([]byte{c})
			idx++
		}
	} else { // no escape, print args concatinated as they are
		fmt.Print(concated)
	}

	if !*argNoNewline {
		fmt.Println()
	}
}
