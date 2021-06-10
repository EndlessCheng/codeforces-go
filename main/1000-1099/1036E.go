package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type vec36 struct{ x, y int64 }
type line36 struct{ p1, p2 vec36 }

func (a vec36) sub(b vec36) vec36 { return vec36{a.x - b.x, a.y - b.y} }
func (a vec36) det(b vec36) int64 { return a.x*b.y - a.y*b.x }
func (a line36) vec() vec36       { return a.p2.sub(a.p1) }
func (a vec36) inRect(l line36) bool {
	return (a.x >= l.p1.x || a.x >= l.p2.x) &&
		(a.x <= l.p1.x || a.x <= l.p2.x) &&
		(a.y >= l.p1.y || a.y >= l.p2.y) &&
		(a.y <= l.p1.y || a.y <= l.p2.y)
}

func CF1036E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, ans int
	Fscan(in, &n)
	l := make([]line36, n)
	for i := range l {
		Fscan(in, &l[i].p1.x, &l[i].p1.y, &l[i].p2.x, &l[i].p2.y)
		ans += 1 + gcd(abs(int(l[i].p2.x-l[i].p1.x)), abs(int(l[i].p2.y-l[i].p1.y)))
	}
	for i, a := range l {
		dup := map[vec36]struct{}{}
		for _, b := range l[:i] {
			va := a.vec()
			d1 := va.det(b.p1.sub(a.p1))
			d2 := va.det(b.p2.sub(a.p1))
			x, y, d := b.p1.x*d2-b.p2.x*d1, b.p1.y*d2-b.p2.y*d1, d2-d1
			if d == 0 || x%d != 0 || y%d != 0 {
				continue
			}
			if p := (vec36{x / d, y / d}); p.inRect(a) && p.inRect(b) {
				dup[p] = struct{}{}
			}
		}
		ans -= len(dup)
	}
	Fprint(out, ans)
}

//func main() { CF1036E(os.Stdin, os.Stdout) }
