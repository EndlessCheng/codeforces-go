package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

type vec struct {
	x, y float64
}

func (a vec) sub(b vec) vec     { return vec{a.x - b.x, a.y - b.y} }
func (a vec) len() float64      { return math.Hypot(a.x, a.y) }
func (a vec) det(b vec) float64 { return a.x*b.y - a.y*b.x }

type line struct {
	p1, p2 vec
}

func (a vec) disToLine(l line) float64 {
	v, u := l.p2.sub(l.p1), a.sub(l.p1)
	return math.Abs(v.det(u)) / v.len()
}

// github.com/EndlessCheng/codeforces-go
func Sol772B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ps := make([]vec, n, n+2)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	ps = append(ps, ps[0], ps[1])
	ans := math.MaxFloat64
	for i := 1; i <= n; i++ {
		if d := ps[i].disToLine(line{ps[i-1], ps[i+1]}); d < ans {
			ans = d
		}
	}
	Fprintf(out, "%.10f", ans/2)
}

//func main() {
//	Sol772B(os.Stdin, os.Stdout)
//}
