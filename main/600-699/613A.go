package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

const eps613 = 1e-6

type vec613 struct {
	x, y float64
}

func (a vec613) add(b vec613) vec613  { return vec613{a.x + b.x, a.y + b.y} }
func (a vec613) sub(b vec613) vec613  { return vec613{a.x - b.x, a.y - b.y} }
func (a vec613) mul(k float64) vec613 { return vec613{a.x * k, a.y * k} }
func (a vec613) len2() float64        { return a.x*a.x + a.y*a.y }
func (a vec613) dot(b vec613) float64 { return a.x*b.x + a.y*b.y }
func (a vec613) det(b vec613) float64 { return a.x*b.y - a.y*b.x }

type line613 struct {
	p1, p2 vec613
}

func (a vec613) perpendicular(l line613) line613 {
	return line613{a, a.add(vec613{l.p1.y - l.p2.y, l.p2.x - l.p1.x})}
}

func (a line613) intersection(b line613) vec613 {
	va, vb := a.p2.sub(a.p1), b.p2.sub(b.p1)
	k := vb.det(b.p1.sub(a.p1)) / vb.det(a.p2.sub(a.p1))
	return a.p1.add(va.mul(k))
}

func (a vec613) onSeg(l line613) bool {
	p1 := l.p1.sub(a)
	p2 := l.p2.sub(a)
	return math.Abs(p1.det(p2)) < eps613 && p1.dot(p2) < eps613
}

func (a vec613) disToSeg(l line613) float64 {
	p := l.intersection(a.perpendicular(l))
	if !p.onSeg(l) {
		if l.p2.sub(l.p1).dot(p.sub(l.p1)) < -eps613 {
			p = l.p1
		} else {
			p = l.p2
		}
	}
	return a.sub(p).len2()
}

// github.com/EndlessCheng/codeforces-go
func Sol613A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	var o vec613
	Fscan(in, &n, &o.x, &o.y)
	minD2, maxD2 := 1e18, 0.0
	ps := make([]vec613, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
		ps[i] = ps[i].sub(o)
		if d2 := ps[i].len2(); d2 > maxD2 {
			maxD2 = d2
		}
	}
	ls := make([]line613, n)
	for i := 0; i < n-1; i++ {
		ls[i] = line613{ps[i], ps[i+1]}
	}
	ls[n-1] = line613{ps[n-1], ps[0]}
	for _, l := range ls {
		if d2 := (vec613{0, 0}).disToSeg(l); d2 < minD2 {
			minD2 = d2
		}
	}
	Fprintf(out, "%.18f", (maxD2-minD2)*math.Pi)
}

//func main() {
//	Sol613A(os.Stdin, os.Stdout)
//}
