package main

import (
	. "fmt"
	"io"
	"os"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	for i, b := range s {
		if b > '1' {
			for ; i < len(s); i++ {
				s[i] = '1'
			}
			break
		}
	}
	v, _ := strconv.ParseInt(string(s), 2, 64)
	Fprint(out, v)
}

func main() { run(os.Stdin, os.Stdout) }
