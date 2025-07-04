package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
type vec struct{ x, y int }

func (a vec) sub(b vec) vec  { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int  { return a.x*b.y - a.y*b.x }
func (a vec) len2() int      { return a.x*a.x + a.y*a.y }
func (a vec) dis2(b vec) int { return b.sub(a).len2() }

func cf1017E(in io.Reader, out io.Writer) {
	smallestRepresentation := func(s []int) []int {
		n := len(s)
		s = append(s, s...)
		i := 0
		for j := 1; j < n; {
			k := 0
			for k < n && s[i+k] == s[j+k] {
				k++
			}
			if k >= n {
				break
			}
			if s[i+k] < s[j+k] {
				j += k + 1
			} else {
				i, j = j, max(j, i+k)+1
			}
		}
		return s[i : i+n]
	}
	convexHull := func(n int) []int {
		ps := make([]vec, n)
		for i := range ps {
			Fscan(in, &ps[i].x, &ps[i].y)
		}

		q := []vec{}
		slices.SortFunc(ps, func(a, b vec) int { return cmp.Or(a.x-b.x, a.y-b.y) })
		for _, p := range ps {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}
		m := len(q)
		for i := len(ps) - 2; i >= 0; i-- {
			p := ps[i]
			for len(q) > m && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}

		s := make([]int, 1, len(q)*4-4)
		s[0] = q[0].dis2(q[1])
		for i := 2; i < len(q); i++ {
			s = append(s, q[i-2].dis2(q[i])+1e18, q[i-1].dis2(q[i]))
		}
		s = append(s, q[len(q)-2].dis2(q[1])+1e18)

		return smallestRepresentation(s)
	}

	var n, m int
	Fscan(in, &n, &m)
	a := convexHull(n)
	b := convexHull(m)
	if slices.Equal(a, b) {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

//func main() { cf1017E(bufio.NewReader(os.Stdin), os.Stdout) }
