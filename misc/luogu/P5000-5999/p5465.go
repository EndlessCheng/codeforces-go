package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p5465(in io.Reader, _w io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, r, x int
	Fscan(in, &n)
	p0 := make([]int, n+1)
	p0[1] = 1
	for i := 2; i <= n; i++ {
		Fscan(in, &p0[i])
	}

	const mx = 19
	type pair struct{ to, sum int }
	pa := make([][mx]pair, n+1)
	l := n
	for i := n; i > 0; i-- {
		l = min(l, p0[i])
		pa[i][0] = pair{l, i - l}
	}

	for i := 0; i < mx-1; i++ {
		for v := range pa {
			p := pa[v][i]
			q := pa[p.to][i]
			pa[v][i+1] = pair{q.to, p.sum + q.sum + (p.to-q.to)<<i}
		}
	}

	f := func(l int) int {
		if l >= p0[x] {
			return x - l
		}
		res := x - p0[x]
		step := 1
		cur := p0[x]
		for k := mx - 1; k >= 0; k-- {
			p := pa[cur][k]
			if p.to > l {
				res += p.sum + (cur-p.to)*step
				step += 1 << k
				cur = p.to
			}
		}
		step++
		res += (cur - l) * step
		return res
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r, &x)
		res := f(l) - f(r+1)
		d := r - l + 1
		g := gcd(res, d)
		Fprintf(out, "%d/%d\n", res/g, d/g)
	}
}

//func main() { p5465(bufio.NewReader(os.Stdin), os.Stdout) }
