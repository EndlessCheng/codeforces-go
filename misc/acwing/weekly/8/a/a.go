package main

import (
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, n, a, b, c int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b, &c, &s)
		c0 := strings.Count(s, "0")
		if a > b {
			a, b, c0 = b, a, n-c0
		}
		Fprintln(out, c0*a+(n-c0)*min(b, a+c))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
