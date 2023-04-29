package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1608C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y, i int }, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
			a[i].i = i
		}

		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
		suf := make([]int, n)
		suf[n-1] = a[n-1].y
		mn := a[n-1].y
		for i := n - 2; i >= 0; i-- {
			mn = min(mn, a[i].y)
			if a[i].y > suf[i+1] {
				suf[i] = mn
			} else {
				suf[i] = suf[i+1]
			}
		}

		ans := bytes.Repeat([]byte{'1'}, n)
		pre := 0
		for i, p := range a[:n-1] {
			pre = max(pre, p.y)
			if pre < suf[i+1] {
				ans[p.i] = '0'
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { CF1608C(os.Stdin, os.Stdout) }
