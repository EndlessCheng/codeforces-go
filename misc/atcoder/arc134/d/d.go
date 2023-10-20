package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	x := a[:0]
	y := b[:0]
	for i, v := range a {
		for len(x) > 0 && v < x[len(x)-1] {
			x = x[:len(x)-1]
			y = y[:len(y)-1]
		}
		x = append(x, v)
		y = append(y, b[i])
	}

	mn := y[0]
	for i, v := range x {
		if v > x[0] {
			break
		}
		mn = min(mn, y[i])
	}
	if mn <= x[0] {
		Fprint(out, x[0], mn)
		return
	}

	r := sort.SearchInts(x, y[0]+1)
	x = x[:r]
	l := sort.SearchInts(x, y[0])

	ans := append(x, y[:r]...)
	less := func() bool {
		for i, v := range y[:l] {
			if v != ans[l+i] {
				return v < ans[l+i]
			}
		}
		return true
	}
	if less() {
		ans = append(ans[:l], ans[r:r+l]...)
	}

	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
