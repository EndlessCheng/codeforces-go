package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
type pair85 struct{ x, y int }

func f85(v, c, lim int, a []pair85, cs []int) bool {
	cs[v] = c
	for w, p := range a {
		if abs85(a[v].x-p.x)+abs85(a[v].y-p.y) > lim && (cs[w] == c || cs[w] == 0 && !f85(w, -c, lim, a, cs)) {
			return false
		}
	}
	return true
}

func CF85E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	var n, ansCC int
	Fscan(in, &n)
	a := make([]pair85, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	ans := sort.Search(1e4+1, func(lim int) bool {
		cs := make([]int, n)
		cc := 0 // 想象用两个圆去覆盖这 n 个点，某些点在两圆交集中，这些点去哪个集合都可以
		for i, c := range cs {
			if c == 0 {
				cc++
				if !f85(i, 1, lim, a, cs) {
					return false
				}
			}
		}
		ansCC = cc
		return true
	})
	p2 := 1
	for ; ansCC > 0; ansCC-- {
		p2 = p2 * 2 % mod
	}
	Fprintln(out, ans)
	Fprintln(out, p2)
}

//func main() { CF85E(os.Stdin, os.Stdout) }
func abs85(x int) int { if x < 0 { return -x }; return x }
