package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprintln(out, 6*n*(n-1)+1)
}

func main() { run(os.Stdin, os.Stdout) }
