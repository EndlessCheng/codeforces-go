package main

import (
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
type vecF struct{ x, y float64 }

func (a vecF) sub(b vecF) vecF     { return vecF{a.x - b.x, a.y - b.y} }
func (a vecF) det(b vecF) float64  { return a.x*b.y - a.y*b.x }
func (a vecF) len2() float64       { return a.x*a.x + a.y*a.y }
func (a vecF) dis2(b vecF) float64 { return a.sub(b).len2() }

func CF1C(in io.Reader, out io.Writer) {
	const eps = 1e-2 // 由于题目保证正多边形边数不超过 100，故 gcdf 的结果不会小于 2*Pi/100，这里简单地写成 1e-2 即可
	gcdf := func(a, b float64) float64 {
		for a > eps {
			a, b = math.Mod(b, a), a
		}
		return b
	}

	var a, b, c vecF
	Fscan(in, &a.x, &a.y, &b.x, &b.y, &c.x, &c.y)
	ab, ac := b.sub(a), c.sub(a)
	ab2, ac2 := ab.len2(), ac.len2()
	r2 := 0.25 * ab2 / ab.det(ac) * ac2 / ab.det(ac) * b.dis2(c) // 外接圆半径 r = abc/4S△abc
	a1 := math.Acos(1 - ab2/2/r2) // 余弦定理
	a2 := math.Acos(1 - ac2/2/r2)
	a3 := 2*math.Pi - a1 - a2
	t := gcdf(gcdf(a1, a2), a3)
	Fprintf(out, "%.8f", math.Pi/t*r2*math.Sin(t))
}

//func main() { CF1C(os.Stdin, os.Stdout) }
