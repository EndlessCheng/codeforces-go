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
	var n, mx, r int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
	}
	for _, v := range a {
		if v < mx && abs(v*2-mx) < abs(r*2-mx) {
			r = v
		}
	}
	Fprint(out, mx, r)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
func abs(x int) int { if x < 0 { return -x }; return x }
