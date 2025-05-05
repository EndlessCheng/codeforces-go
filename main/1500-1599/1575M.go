package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type vec75 struct{ x, y int }
func (a vec75) sub(b vec75) vec75 { return vec75{a.x - b.x, a.y - b.y} }
func (a vec75) dot(b vec75) int   { return a.x*b.x + a.y*b.y }
func (a vec75) det(b vec75) int   { return a.x*b.y - a.y*b.x }

func cf1575M(in io.Reader, out io.Writer) {
	var n, m, ans int
	var s []byte
	Fscan(in, &n, &m)
	n++
	m++
	pos := make([][]int, n)
	for i := range n {
		Fscan(in, &s)
		for j, b := range s {
			if b == '1' {
				pos[i] = append(pos[i], j)
			}
		}
	}

	for y := range m {
		q := []vec75{}
		for x, ys := range pos {
			if ys == nil {
				continue
			}
			if len(ys) > 1 && ys[1]-y < y-ys[0] {
				ys = ys[1:]
				pos[x] = ys
			}
			p := vec75{x, x*x - y*ys[0]*2 + ys[0]*ys[0]}
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}
		for x := range n {
			p := vec75{-x * 2, 1}
			for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
				q = q[1:]
			}
			ans += p.dot(q[0]) + x*x + y*y
		}
	}
	Fprint(out, ans)
}

//func main() { cf1575M(bufio.NewReader(os.Stdin), os.Stdout) }
