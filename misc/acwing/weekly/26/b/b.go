package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	var maxK, pos [26]int
	for i := range pos {
		pos[i] = -1
	}
	for i, b := range s {
		b -= 'a'
		maxK[b] = max(maxK[b], i-pos[b])
		pos[b] = i
	}
	ans := int(1e9)
	for i, p := range pos {
		ans = min(ans, max(maxK[i], len(s)-p))
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
