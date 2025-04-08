package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
type vec31 struct{ x, y int }

func (a vec31) sub(b vec31) vec31 { return vec31{a.x - b.x, a.y - b.y} }
func (a vec31) dot(b vec31) int   { return a.x*b.x + a.y*b.y }
func (a vec31) det(b vec31) int   { return a.x*b.y - a.y*b.x }

func cf631E(in io.Reader, out io.Writer) {
	var n, tot, mx, s int
	Fscan(in, &n)
	a := make([]int, n)
	q := []vec31{}
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i] * (i + 1)
		s += a[i]
		v := vec31{i, -s}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(v.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, v)
	}

	s = 0
	for i, x := range a {
		s += x
		p := vec31{x, 1}
		j := sort.Search(len(q)-1, func(j int) bool { return p.dot(q[j]) > p.dot(q[j+1]) })
		mx = max(mx, p.dot(q[j])-x*i+s)
	}
	Fprint(out, tot+mx)
}

//func main() { cf631E(bufio.NewReader(os.Stdin), os.Stdout) }
