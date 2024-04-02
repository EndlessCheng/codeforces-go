package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

type vec71 struct {
	x, y int
}

func (a vec71) sub(b vec71) vec71 { return vec71{a.x - b.x, a.y - b.y} }
func (a vec71) len() float64      { return math.Hypot(float64(a.x), float64(a.y)) }

// github.com/EndlessCheng/codeforces-go
func cf671A(reader io.Reader, writer io.Writer) {
	type pair struct {
		d   float64
		idx int
	}
	read := func(in io.Reader) vec71 {
		var x, y int
		Fscan(in, &x, &y)
		return vec71{x, y}
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	pa, pb, bin := read(in), read(in), read(in)
	pa = pa.sub(bin)
	pb = pb.sub(bin)
	var n int
	Fscan(in, &n)
	das := make([]pair, n)
	dbs := make([]pair, n)
	ans := 0.0
	for i := range das {
		p := read(in).sub(bin)
		lenP := p.len()
		ans += lenP
		das[i] = pair{lenP - p.sub(pa).len(), i}
		dbs[i] = pair{lenP - p.sub(pb).len(), i}
	}
	ans *= 2

	sort.Slice(das, func(i, j int) bool { return das[i].d > das[j].d })
	sort.Slice(dbs, func(i, j int) bool { return dbs[i].d > dbs[j].d })
	if n == 1 || das[0].d <= 0 || dbs[0].d <= 0 {
		ans -= max(das[0].d, dbs[0].d)
	} else if das[0].idx == dbs[0].idx {
		sum1 := das[0].d
		sum2 := dbs[0].d
		if dbs[1].d > 0 {
			sum1 += dbs[1].d
		}
		if das[1].d > 0 {
			sum2 += das[1].d
		}
		ans -= max(sum1, sum2)
	} else {
		ans -= das[0].d + dbs[0].d
	}
	Fprintf(out, "%.12f", ans)
}

//func main() { cf671A(os.Stdin, os.Stdout) }
