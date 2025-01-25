package main

import (
	. "fmt"
	"io"
	"os"
	"strings"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(in, &n, &s)
	c := strings.Count(s, "T")
	if c*2 > n || c*2 == n && s[n-1] != 'T' {
		Fprint(out, "T")
	} else {
		Fprint(out, "A")
	}
}

func main() { run(os.Stdin, os.Stdout) }
