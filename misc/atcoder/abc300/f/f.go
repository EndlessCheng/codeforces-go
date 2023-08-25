package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, m, k int
	var s string
	Fscan(bufio.NewReader(in), &n, &m, &k, &s)
	pos := []int{}
	for i, c := range s {
		if c == 'x' {
			pos = append(pos, i)
		}
	}
	cntX := len(pos)
	ans := min(k/cntX*n+pos[k%cntX], n*m)
	for _, p := range pos {
		k++
		res := min(k/cntX*n+pos[k%cntX], n*m) - p - 1
		ans = max(ans, res)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
