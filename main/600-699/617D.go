package main

import (
	. "fmt"
	"io"
)

type vec617 struct{ x, y int64 }

func (a vec617) sub(b vec617) vec617 { return vec617{a.x - b.x, a.y - b.y} }
func (a vec617) dot(b vec617) int64  { return a.x*b.x + a.y*b.y }
func (a vec617) det(b vec617) int64  { return a.x*b.y - a.y*b.x }

// github.com/EndlessCheng/codeforces-go
func CF617D(in io.Reader, out io.Writer) {
	f := func(a, b, c vec617) bool {
		a, b = b.sub(a), c.sub(a)
		return (a.x == 0 || a.y == 0 || b.x == 0 || b.y == 0) && a.dot(b) <= 0
	}
	var a, b, c vec617
	Fscan(in, &a.x, &a.y, &b.x, &b.y, &c.x, &c.y)
	if d := b.sub(a); (d.x == 0 || d.y == 0) && d.det(c.sub(b)) == 0 {
		Fprint(out, 1)
	} else if f(a, b, c) || f(b, a, c) || f(c, a, b) {
		Fprint(out, 2)
	} else {
		Fprint(out, 3)
	}
}

//func main() { CF617D(os.Stdin, os.Stdout) }
