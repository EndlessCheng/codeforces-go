package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
type vec29 struct{ x, y int64 }

func merge29(a, b []vec29) []vec29 {
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]vec29, 0, n+m)
	for {
		if i == n {
			return append(res, b[j:]...)
		}
		if j == m {
			return append(res, a[i:]...)
		}
		if a[i].y < b[j].y {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
}

func closestPair(ps []vec29) int64 {
	n := len(ps)
	if n <= 1 {
		return math.MaxInt64
	}
	m := n >> 1
	x := ps[m].x
	d2 := min29(closestPair(ps[:m]), closestPair(ps[m:]))
	copy(ps, merge29(ps[:m], ps[m:]))
	checkPs := []vec29{}
	for _, pi := range ps {
		if (pi.x-x)*(pi.x-x) > d2 {
			continue
		}
		for j := len(checkPs) - 1; j >= 0; j-- {
			pj := checkPs[j]
			dy := pi.y - pj.y
			if dy*dy >= d2 {
				break
			}
			dx := pi.x - pj.x
			d2 = min29(d2, dx*dx+dy*dy)
		}
		checkPs = append(checkPs, pi)
	}
	return d2
}

func CF429D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, s int
	Fscan(in, &n)
	ps := make([]vec29, n)
	for i := range ps {
		Fscan(in, &v)
		s += v
		ps[i] = vec29{int64(i), int64(s)}
	}
	Fprint(out, closestPair(ps))
}

func min29(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

//func main() { CF429D(os.Stdin, os.Stdout) }
