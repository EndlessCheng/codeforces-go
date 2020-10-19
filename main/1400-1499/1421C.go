package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1421C(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	n := len(s)
	Fprintln(out, 3)
	Fprintln(out, "R", n-1)
	Fprintln(out, "L", n)
	Fprintln(out, "L", 2)
}

//func main() { CF1421C(os.Stdin, os.Stdout) }
