package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg83 []struct {
	l, r int
	val  int64
}

func (t seg83) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg83) update(o, i int, val int64) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t[o].val = max83(t[o<<1].val, t[o<<1|1].val)
}

func (t seg83) query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return max83(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF1483C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	posL := make([]int, n)
	type pair struct{ v, i int }
	s := []pair{{0, -1}}
	for i := range posL {
		Fscan(in, &v)
		for s[len(s)-1].v > v {
			s = s[:len(s)-1]
		}
		posL[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	dp := make([]int64, n)
	t := make(seg83, 4*n)
	t.build(1, 1, n)
	for i, l := range posL {
		if Fscan(in, &v); i == 0 {
			dp[i] = int64(v)
		} else if l < 0 {
			dp[i] = max83(0, t.query(1, 1, i)) + int64(v)
		} else {
			dp[i] = max83(dp[l], t.query(1, l+1, i)+int64(v))
		}
		t.update(1, i+1, dp[i])
	}
	Fprint(out, dp[n-1])
}

func max83(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

//func main() { CF1483C(os.Stdin, os.Stdout) }
