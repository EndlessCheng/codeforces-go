package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF797D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
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

	var n int
	Fscan(in, &n)
	a := make([]struct{ v, l, r int }, n+1)
	cnt := map[int]int{}
	hasPa := make([]bool, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].v, &a[i].l, &a[i].r)
		cnt[a[i].v]++
		hasPa[a[i].l+1] = true
		hasPa[a[i].r+1] = true
	}
	rt := 1
	for hasPa[rt+1] {
		rt++
	}

	ans := n
	var f func(i, l, r int)
	f = func(i, l, r int) {
		if i < 0 {
			return
		}
		o := a[i]
		if l <= o.v && o.v <= r {
			ans -= cnt[o.v]
		}
		if l < o.v {
			f(o.l, l, min(r, o.v-1))
		}
		if r > o.v {
			f(o.r, max(l, o.v+1), r)
		}
	}
	f(rt, 0, 1e9)
	Fprint(out, ans)
}

//func main() { CF797D(os.Stdin, os.Stdout) }
