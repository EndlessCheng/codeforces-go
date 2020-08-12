package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func Sol1059D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	search := func(l, r float64, f func(x float64) bool) float64 {
		for i := 0; i < 70; i++ {
			mid := (l + r) / 2
			if f(mid) {
				r = mid
			} else {
				l = mid
			}
		}
		return (l + r) / 2
	}

	var n int
	Fscan(in, &n)
	type pair struct{ x, y float64 }
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
		if ps[i].y*ps[0].y < 0 {
			Fprint(out, -1)
			return
		}
	}
	if ps[0].y < 0 {
		for i := range ps {
			ps[i].y = -ps[i].y
		}
	}

	ans := search(0, 1e14, func(r float64) bool {
		minOx := math.MaxFloat64
		for _, p := range ps {
			if 2*r < p.y {
				return false
			}
			if x := p.x + math.Sqrt(p.y*(2*r-p.y)); x < minOx {
				minOx = x
			}
		}
		for _, p := range ps {
			if (p.x-minOx)*(p.x-minOx) > (1+1e-8)*p.y*(2*r-p.y) {
				return false
			}
		}
		return true
	})
	Fprintf(out, "%.16f", ans)
}

//func main() {
//	Sol1059D(os.Stdin, os.Stdout)
//}
