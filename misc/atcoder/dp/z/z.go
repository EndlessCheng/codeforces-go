package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }

func run(in io.Reader, out io.Writer) {
	var n, c int
	Fscan(in, &n, &c)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	q := []vec{{a[0], a[0] * a[0]}}
	for i := 1; ; i++ {
		v := a[i]
		p := vec{-v * 2, 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}
		fi := p.dot(q[0]) + v*v + c
		if i == n-1 {
			Fprint(out, fi)
			return
		}

		p = vec{v, v*v + fi}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
