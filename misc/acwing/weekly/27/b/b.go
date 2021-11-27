package main

import (
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	Fprint(out, sort.Search(n*m, func(v int) bool {
		cnt := 0
		for i := 1; i <= n; i++ {
			cnt += min(v/i, m)
		}
		return cnt >= k
	}))
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
