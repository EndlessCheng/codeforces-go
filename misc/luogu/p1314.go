package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func p1314(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, m, t int
	Fscan(in, &n, &m, &t)
	a := make([]struct{ w, v int }, n)
	for i := range a {
		Fscan(in, &a[i].w, &a[i].v)
	}
	ps := make([]struct{ l, r int }, m)
	for i := range ps {
		Fscan(in, &ps[i].l, &ps[i].r)
		ps[i].l--
	}

	ans := t
	s := make([]struct{ c, s int }, n+1)
	sort.Search(1e6+1, func(w int) bool {
		for i, p := range a {
			s[i+1] = s[i]
			if p.w >= w {
				s[i+1].c++
				s[i+1].s += p.v
			}
		}
		sum := 0
		for _, p := range ps {
			sum += (s[p.r].c - s[p.l].c) * (s[p.r].s - s[p.l].s)
			if sum >= t+ans {
				break
			}
		}
		ans = min(ans, abs(sum-t))
		return sum <= t
	})
	Fprint(out, ans)
}

//func main() { p1314(bufio.NewReader(os.Stdin), os.Stdout) }
