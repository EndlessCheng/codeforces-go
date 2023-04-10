package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1473D(_r io.Reader, _w io.Writer) {
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

	var T, n, m, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s)
		pre := make([]struct{ cur, max, min int }, n+1)
		for i, c := range s {
			pre[i+1].cur = pre[i].cur + int(c&2) - 1
			pre[i+1].max = max(pre[i].max, pre[i+1].cur)
			pre[i+1].min = min(pre[i].min, pre[i+1].cur)
		}
		type pair struct{ up, down int }
		suf := make([]pair, n+1)
		var cur, mx, mn int
		for i := n - 1; i >= 0; i-- {
			cur -= int(s[i]&2) - 1
			mx = max(mx, cur)
			mn = min(mn, cur)
			suf[i] = pair{mx - cur, mn - cur}
		}
		for ; m > 0; m-- {
			Fscan(in, &l, &r)
			p, s := pre[l-1], suf[r]
			Fprintln(out, max(p.max, p.cur+s.up)-min(p.min, p.cur+s.down)+1)
		}
	}
}

//func main() { CF1473D(os.Stdin, os.Stdout) }
