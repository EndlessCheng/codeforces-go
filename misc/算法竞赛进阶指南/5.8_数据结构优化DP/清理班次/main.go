package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
const inf int = 1e18

type seg []struct{ l, r, val int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].val = l, r, inf
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) update(o, i, val int) {
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
	t[o].val = min(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, st, end int
	Fscan(in, &n, &st, &end)
	type pair struct{ l, r, c int }
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &ps[i].l, &ps[i].r, &ps[i].c)
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].r < ps[j].r })

	dp := make([]int, end+1)
	for i := st; i <= end; i++ {
		dp[i] = inf
	}
	t := make(seg, 4*(end+1))
	t.build(1, st-1, end)
	t.update(1, st-1, 0)
	for _, p := range ps {
		dp[p.r] = min(dp[p.r], t.query(1, p.l-1, p.r)+p.c)
		t.update(1, p.r, dp[p.r])
	}
	ans := dp[end]
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
