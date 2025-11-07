package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod7 = 12345

type matrix7 [][]int

func newMatrix7(n, m int) matrix7 {
	a := make(matrix7, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix7) mul(b matrix7) matrix7 {
	c := newMatrix7(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod7
			}
		}
	}
	return c
}

func (a matrix7) powMul(n int, f0 matrix7) matrix7 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf107D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }

	var n, k, v, ans int
	var s string
	Fscan(in, &n, &k)
	g := [26][]int{}
	lcms := [26]int{}
	for i := range lcms {
		lcms[i] = 1
	}
	for range k {
		Fscan(in, &s, &v)
		b := s[0] - 'A'
		g[b] = append(g[b], v)
		lcms[b] = lcm(lcms[b], v)
	}

	size := 1
	for _, l := range lcms {
		size *= l
	}

	m := newMatrix7(size, size)
	var dfs func(int, int, int, bool)
	dfs = func(i, new, old int, ch bool) {
		if i < 0 {
			if ch {
				m[new][old]++
			}
			return
		}
		if g[i] == nil {
			dfs(i-1, new, old, ch)
			return
		}
		l := lcms[i]
		for j := range l {
			dfs(i-1, new*l+j, old*l+j, ch)
			if !ch {
				dfs(i-1, new*l+(j+1)%l, old*l+j, true)
			}
		}
	}
	dfs(25, 0, 0, false)

	f0 := newMatrix7(size, 1)
	f0[0][0] = 1
	fn := m.powMul(n, f0)
nxt:
	for mask, row := range fn {
		for i, a := range g {
			if a == nil {
				continue
			}
			c := mask % lcms[i]
			for _, v := range a {
				if c%v == 0 {
					goto ok
				}
			}
			continue nxt
		ok:
			mask /= lcms[i]
		}
		ans += row[0]
	}
	Fprint(out, ans%mod7)
}

//func main() { cf107D(bufio.NewReader(os.Stdin), os.Stdout) }
