package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func f(n, k int) int {
	m := 1<<(n-1) - 1
	if k == m {
		return n
	}
	if k > m {
		return f(n-1, k-m-1)
	}
	return f(n-1, k)
}

func run(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	Fprint(out, f(n, k-1))
}

func main() { run(os.Stdin, os.Stdout) }
