package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

type vecF struct{ x, y float64 }

func (a vecF) add(b vecF) vecF     { return vecF{a.x + b.x, a.y + b.y} }
func (a vecF) sub(b vecF) vecF     { return vecF{a.x - b.x, a.y - b.y} }
func (a vecF) dot(b vecF) float64  { return a.x*b.x + a.y*b.y }
func (a vecF) det(b vecF) float64  { return a.x*b.y - a.y*b.x }
func (a vecF) len2() float64       { return a.x*a.x + a.y*a.y }
func (a vecF) dis2(b vecF) float64 { return a.sub(b).len2() }
func (a vecF) len() float64        { return math.Hypot(a.x, a.y) }
func (a vecF) dis(b vecF) float64  { return a.sub(b).len2() }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, h, m float64
	Fscan(in, &a, &b, &h, &m)

	calcP := func(d, r float64) vecF {
		d = -d
		d *= 2
		if d <= 0.5 {
			d *= math.Pi
			return vecF{r * math.Sin(d), r * math.Cos(d)}
		}
		d -= 0.5
		if d <= 0.5 {
			d *= math.Pi
			return vecF{r * math.Cos(d), -r * math.Sin(d)}
		}
		d -= 0.5
		if d <= 0.5 {
			d *= math.Pi
			return vecF{-r * math.Sin(d), -r * math.Cos(d)}
		}
		d -= 0.5
		d *= math.Pi
		return vecF{-r * math.Cos(d), r * math.Sin(d)}
	}
	m /= 60
	h += m
	h /= 12
	x, y := calcP(h, a), calcP(m, b)
	Fprintf(out, "%.13f", x.dis(y))
}

func main() { run(os.Stdin, os.Stdout) }
