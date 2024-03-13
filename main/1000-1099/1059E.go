package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1059E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, limL, limS, ans int
	Fscan(in, &n, &limL, &limS)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] > limS {
			Fprint(out, -1)
			return
		}
	}
	const mx = 17
	type pair struct{ p, s int }
	pa := make([][mx]pair, n)
	pa[0][0].p = -1
	for i := 1; i < n; i++ {
		Fscan(in, &pa[i][0].p)
		pa[i][0].p--
		p := pa[i][0].p
		if p >= 0 {
			pa[i][0].s = a[p]
		}
	}

	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1] = pair{pp.p, p.s + pp.s}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}

	left := make([]int, n)
	for i := n - 1; i > 0; i-- {
		p := pa[i][0].p
		if left[i] > 0 {
			left[p] = max(left[p], left[i]-1)
			continue
		}
		ans++
		v, l, s := i, 0, limS-a[i]
		for j := mx - 1; j >= 0; j-- {
			if l|1<<j >= limL {
				continue
			}
			if p := pa[v][j]; p.p != -1 && p.s <= s {
				l |= 1 << j
				s -= p.s
				v = p.p
			}
		}
		left[p] = max(left[p], l)
	}
	if left[0] == 0 {
		ans++
	}
	Fprint(out, ans)
}

//func main() { cf1059E(os.Stdin, os.Stdout) }
