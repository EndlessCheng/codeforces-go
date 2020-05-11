package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type vec struct{ x, y int }
type vs []vec

func (a *vec) add(b vec)          { a.x += b.x; a.y += b.y }
func (a vec) len2() int           { return a.x*a.x + a.y*a.y }
func (a vec) polarAngle() float64 { return math.Atan2(float64(a.y), float64(a.x)) }
func (v vs) Len() int             { return len(v) }
func (v vs) Less(i, j int) bool   { return v[i].polarAngle() < v[j].polarAngle() }
func (v vs) Swap(i, j int)        { v[i], v[j] = v[j], v[i] }

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	ps := make(vs, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	sort.Sort(ps)
	ps = append(ps, ps...)
	for i := 0; i < n; i++ {
		p := vec{}
		for j := i; j < n+i; j++ {
			p.add(ps[j])
			if l := p.len2(); l > ans {
				ans = l
			}
		}
	}
	Fprintf(_w, "%.12f", math.Sqrt(float64(ans)))
}

func main() { run(os.Stdin, os.Stdout) }
