package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	c1 := strings.Count(s, "1")
	p1 := strings.LastIndexByte(s, '1')
	if p1 == 0 {
		Fprint(out, len(s)-1)
	} else {
		Fprint(out, len(s)+p1-c1+2)
	}
}

func main() { run(os.Stdin, os.Stdout) }
