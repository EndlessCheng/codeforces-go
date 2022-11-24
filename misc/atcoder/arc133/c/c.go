package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, sa, sb int
	Fscan(in, &n, &m, &k)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sa += ((k-1)*m - v) % k
	}
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		sb += ((k-1)*n - v) % k
	}
	if sa%k != sb%k {
		Fprint(out, -1)
	} else {
		Fprint(out, n*m*(k-1)-max(sa, sb))
	}
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
