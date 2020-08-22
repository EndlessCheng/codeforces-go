package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	var a, b, c int
	Fscan(bufio.NewReader(_r), &a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a+c < b {
		a += c
	} else {
		c -= b - a
		a = b + c/2
	}
	Fprint(out, sort.Search(1e7, func(h int) bool { return h*(h+1)/2 > a })-1)
}

func main() { run(os.Stdin, os.Stdout) }
