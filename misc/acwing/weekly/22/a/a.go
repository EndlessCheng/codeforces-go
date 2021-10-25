package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, n, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b)
		Fprintln(out, min(n-a, b+1))
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
