package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1029C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ l, r int }
	merge := func(p, q pair) pair { return pair{max(p.l, q.l), min(p.r, q.r)} }

	var n, ans int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}

	suf := make([]pair, n+1)
	suf[n].r = 1e9
	for i := n - 1; i > 0; i-- {
		suf[i] = merge(suf[i+1], a[i])
	}

	pre := pair{0, 1e9}
	for i, p := range a {
		m := merge(pre, suf[i+1])
		ans = max(ans, m.r-m.l)
		pre = merge(pre, p)
	}
	Fprint(out, ans)
}

//func main() { cf1029C(os.Stdin, os.Stdout) }
