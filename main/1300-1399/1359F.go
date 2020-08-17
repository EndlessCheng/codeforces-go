package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

type vec1359 struct{ x, y int }

func (a vec1359) sub(b vec1359) vec1359 { return vec1359{a.x - b.x, a.y - b.y} }
func (a vec1359) dot(b vec1359) int     { return a.x*b.x + a.y*b.y }
func (a vec1359) det(b vec1359) int     { return a.x*b.y - a.y*b.x }
func (a vec1359) len() float64          { return math.Hypot(float64(a.x), float64(a.y)) }

// github.com/EndlessCheng/codeforces-go
func CF1359F(_r io.Reader, out io.Writer) {
	// faster than math.Min and math.Max
	min := func(a, b float64) float64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
	type car struct {
		p, d    vec1359
		s, norm float64
	}

	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	cs := make([]car, n)
	for i := range cs {
		Fscan(in, &cs[i].p.x, &cs[i].p.y, &cs[i].d.x, &cs[i].d.y, &cs[i].s)
		cs[i].norm = cs[i].d.len()
	}

	ans := 1e99
	for i, p := range cs {
		for j := i + 1; j < n; j++ {
			q := cs[j]
			u := p.p.sub(q.p)
			if d := p.d.det(q.d); d != 0 {
				d1, d2 := q.d.det(u), p.d.det(u)
				if d > 0 && d1 >= 0 && d2 >= 0 || d < 0 && d1 <= 0 && d2 <= 0 {
					ans = min(ans, max(p.norm*float64(d1)/(p.s*float64(d)), q.norm*float64(d2)/(q.s*float64(d))))
				}
				continue
			}
			if u.det(p.d) != 0 {
				continue
			}
			// u.len() could be avoided https://codeforces.com/contest/1359/submission/81868329
			if l := u.len(); p.d.dot(q.d) > 0 {
				if u.dot(q.d) >= 0 {
					ans = min(ans, l/q.s)
				} else {
					ans = min(ans, l/p.s)
				}
			} else {
				if u.dot(q.d) > 0 {
					ans = min(ans, l/(p.s+q.s))
				}
			}
		}
	}
	if ans == 1e99 {
		Fprint(out, "No show :(")
	} else {
		Fprintf(out, "%.8f", ans)
	}
}

//func main() { CF1359F(os.Stdin, os.Stdout) }
