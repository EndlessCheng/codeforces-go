package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	type vec struct{ x, y float64 }
	merge := func(a, b []vec) []vec {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]vec, 0, n+m)
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

	// 最近点对
	// 调用前 ps 必须按照 x 坐标排序
	var closestPair func([]vec) float64
	closestPair = func(ps []vec) float64 {
		n := len(ps)
		if n <= 1 {
			return math.MaxFloat64
		}
		m := n >> 1
		x := ps[m].x
		d := math.Min(closestPair(ps[:m]), closestPair(ps[m:]))
		for i, p := range merge(ps[:m], ps[m:]) {
			ps[i] = p
		}
		checkPs := []vec{}
		for _, pi := range ps {
			if math.Abs(pi.x-x) > d+1e-8 {
				continue
			}
			for j := len(checkPs) - 1; j >= 0; j-- {
				pj := checkPs[j]
				dy := pi.y - pj.y
				if dy >= d {
					break
				}
				dx := pi.x - pj.x
				d = math.Min(d, math.Hypot(dx, dy))
			}
			checkPs = append(checkPs, pi)
		}
		return d
	}

	n := read()
	ps := make([]vec, n)
	for i := range ps {
		ps[i] = vec{float64(read()), float64(read())}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].x < ps[j].x })
	Fprintf(out, "%.4f", closestPair(ps))
}

func main() {
	solve(os.Stdin, os.Stdout)
}
