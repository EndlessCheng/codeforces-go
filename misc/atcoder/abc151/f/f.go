package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/rand"
	"os"
)

// https://space.bilibili.com/206214
const eps = 1e-12

type vec struct{ x, y float64 }

func (a vec) add(b vec) vec      { return vec{a.x + b.x, a.y + b.y} }
func (a vec) sub(b vec) vec      { return vec{a.x - b.x, a.y - b.y} }
func (a vec) len2() float64      { return a.x*a.x + a.y*a.y }
func (a vec) dis2(b vec) float64 { return a.sub(b).len2() }
func (a vec) div(k float64) vec  { return vec{a.x / k, a.y / k} }

func circumcenter(a, b, c vec) vec {
	a1, b1, a2, b2 := b.x-a.x, b.y-a.y, c.x-a.x, c.y-a.y
	c1, c2, d := a1*a1+b1*b1, a2*a2+b2*b2, 2*(a1*b2-a2*b1)
	return vec{a.x + (c1*b2-c2*b1)/d, a.y + (a1*c2-a2*c1)/d}
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	ps := make([]vec, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}

	rand.Shuffle(len(ps), func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })
	o := ps[0]
	r2 := 0.
	for i, p := range ps {
		if p.dis2(o) < r2+eps {
			continue
		}
		o, r2 = p, 0
		for j, q := range ps[:i] {
			if q.dis2(o) < r2+eps {
				continue
			}
			o = vec{(p.x + q.x) / 2, (p.y + q.y) / 2}
			r2 = p.dis2(o)
			for _, x := range ps[:j] {
				if x.dis2(o) > r2+eps {
					o = circumcenter(p, q, x)
					r2 = p.dis2(o)
				}
			}
		}
	}
	Fprintf(out, "%.10f", math.Sqrt(r2))
}

func main() { run(os.Stdin, os.Stdout) }
