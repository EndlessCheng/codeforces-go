package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
type vec11 struct{ x, y int }

func (a vec11) sub(b vec11) vec11 { return vec11{a.x - b.x, a.y - b.y} }
func (a vec11) dot(b vec11) int   { return a.x*b.x + a.y*b.y }
func (a vec11) det(b vec11) int   { return a.x*b.y - a.y*b.x }

func cf311B(in io.Reader, out io.Writer) {
	var n, m, p, h, t int
	Fscan(in, &n, &m, &p)
	if p >= m {
		Fprint(out, 0)
		return
	}
	dis := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &dis[i])
		dis[i] += dis[i-1]
	}
	a := make([]int, m)
	for i := range a {
		Fscan(in, &h, &t)
		a[i] = t - dis[h-1]
	}
	slices.Sort(a)
	s := make([]int, m+1)
	for i, v := range a {
		s[i+1] = s[i] + v
	}

	f := make([]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = 1e18
	}
	for k := 1; k <= p; k++ {
		q := []vec11{{}}
		for i := 1; i <= m; i++ {
			pi := vec11{-a[i-1], 1}
			for len(q) > 1 && pi.dot(q[0]) >= pi.dot(q[1]) {
				q = q[1:]
			}
			v := vec11{i, s[i] + f[i]}
			f[i] = pi.dot(q[0]) + a[i-1]*i - s[i]
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(v.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
	}
	Fprint(out, f[m])
}

//func main() { cf311B(bufio.NewReader(os.Stdin), os.Stdout) }
