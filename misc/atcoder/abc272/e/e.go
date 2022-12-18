package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v int
	Fscan(in, &n, &m)
	b := make([][]int, m+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		l := 1
		if v < 0 {
			l = (-v + i - 1) / i
		}
		r := min(m, (n-v)/i)
		for j := l; j <= r; j++ {
			b[j] = append(b[j], v+j*i)
		}
	}
	time := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for _, v := range b[i] {
			time[v] = i
		}
		mex := 0
		for time[mex] == i {
			mex++
		}
		Fprintln(out, mex)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
