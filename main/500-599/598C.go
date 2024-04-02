package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/big"
	"sort"
)

// https://space.bilibili.com/206214
func cf598C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y, minI int
	Fscan(in, &n)
	type point struct {
		a       float64
		x, y, i int
	}
	a := make([]point, n, n+1)
	for i := range a {
		Fscan(in, &x, &y)
		a[i] = point{math.Atan2(float64(y), float64(x)), x, y, i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].a < a[j].a })
	a = append(a, a[0])

	mx := big.NewRat(-1e18, 1)
	for i := 0; i < n; i++ {
		p, q := a[i], a[i+1]
		d := p.x*q.x + p.y*q.y
		neg := d < 0
		d *= d
		if neg {
			d = -d
		}
		cos2 := big.NewRat(int64(d), int64((p.x*p.x+p.y*p.y)*(q.x*q.x+q.y*q.y)))
		if cos2.Cmp(mx) > 0 {
			mx = cos2
			minI = i
		}
	}
	Fprint(out, a[minI].i+1, a[minI+1].i+1)
}

//func main() { cf598C(os.Stdin, os.Stdout) }
