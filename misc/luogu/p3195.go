package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type vec3195 struct{ x, y int }

func (a vec3195) sub(b vec3195) vec3195 { return vec3195{a.x - b.x, a.y - b.y} }
func (a vec3195) dot(b vec3195) int     { return a.x*b.x + a.y*b.y }
func (a vec3195) det(b vec3195) int     { return a.x*b.y - a.y*b.x }

func p3195(in io.Reader, out io.Writer) {
	var n, L int
	Fscan(in, &n, &L)
	L++
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1] + 1
	}

	q := []vec3195{{}}
	for i := 1; ; i++ {
		v := s[i]
		p := vec3195{-2 * (v - L), 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}
		f := p.dot(q[0]) + (v-L)*(v-L)
		if i == n {
			Fprint(out, f)
			return
		}
		vi := vec3195{v, v*v + f}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(vi.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, vi)
	}
}

//func main() { p3195(bufio.NewReader(os.Stdin), os.Stdout) }
