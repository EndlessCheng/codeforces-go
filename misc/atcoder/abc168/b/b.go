package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	var n int
	var s string
	Fscan(_r, &n, &s)
	if len(s) <= n {
		Fprint(_w, s)
	} else {
		Fprint(_w, s[:n]+"...")
	}
}

func main() { run(os.Stdin, os.Stdout) }
