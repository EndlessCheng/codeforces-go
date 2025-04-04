package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type vec5017 struct{ x, y int }

func (a vec5017) sub(b vec5017) vec5017 { return vec5017{a.x - b.x, a.y - b.y} }
func (a vec5017) dot(b vec5017) int     { return a.x*b.x + a.y*b.y }
func (a vec5017) det(b vec5017) int     { return a.x*b.y - a.y*b.x }

func p5017(in io.Reader, out io.Writer) {
	var n, m, mx int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
	}
	cnt := make([]int, mx+m)
	for _, v := range a {
		cnt[v]++
	}

	f := make([]int, mx+m)
	q := []vec5017{}
	c, s := 0, 0
	c2, s2 := 0, 0
	for i, v := range cnt {
		c += v
		s += v * i
		if i < m {
			f[i] = i*c - s
			continue
		}

		c2 += cnt[i-m]
		s2 += cnt[i-m] * (i - m)
		p := vec5017{c2, s2 + f[i-m]}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)

		p = vec5017{-i, 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}
		f[i] = p.dot(q[0]) + i*c - s
	}
	Fprint(out, slices.Min(f[mx:]))
}

//func main() { p5017(bufio.NewReader(os.Stdin), os.Stdout) }
