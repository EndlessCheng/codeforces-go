package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1824C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	
	var n, v, w, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var f func(int, int) map[int]int
	f = func(v, fa int) map[int]int {
		cnt := map[int]int{}
		mx := 1
		for _, w := range g[v] {
			if w != fa {
				a[w] ^= a[v]
				c := f(w, v)
				if len(cnt) < len(c) {
					cnt, c = c, cnt
				}
				for v, c := range c {
					cnt[v] += c
					mx = max(mx, cnt[v])
				}
			}
		}
		c := len(g[v])
		if v > 0 {
			c--
		}
		if c > 0 {
			ans += c - mx
		} else {
			cnt[a[v]] = 1
		}
		if mx > 1 {
			for v, c := range cnt {
				if c < mx {
					delete(cnt, v)
				} else {
					cnt[v] = 1
				}
			}
		}
		return cnt
	}
	cnt := f(0, -1)
	if cnt[0] == 0 {
		ans++
	}
	Fprint(out, ans)
}

//func main() { CF1824C(os.Stdin, os.Stdout) }
