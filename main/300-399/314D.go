package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf314D(in io.Reader, out io.Writer) {
	var n, x, y int
	Fscan(in, &n)
	type point struct{ x, y int }
	a := make([]point, n)
	for i := range a {
		Fscan(in, &x, &y)
		a[i] = point{x + y, y - x}
	}
	slices.SortFunc(a, func(a, b point) int { return a.x - b.x })

	type pair struct{ min, max int }
	pre := make([]pair, n+1)
	pre[0].min = 2e9
	pre[0].max = -2e9
	for i, p := range a {
		pre[i+1].min = min(pre[i].min, p.y)
		pre[i+1].max = max(pre[i].max, p.y)
	}
	suf := make([]pair, n+1)
	suf[n] = pre[0]
	for i := n - 1; i >= 0; i-- {
		y := a[i].y
		suf[i].min = min(suf[i+1].min, y)
		suf[i].max = max(suf[i+1].max, y)
	}

	ans := sort.Search(2e9, func(mx int) bool {
		l := 0
		for i, p := range a {
			if i+1 < n && p.x == a[i+1].x {
				continue
			}
			for p.x-a[l].x > mx {
				l++
			}
			if max(pre[l].max, suf[i+1].max)-min(pre[l].min, suf[i+1].min) <= mx {
				return true
			}
		}
		return false
	})
	Fprintf(out, "%.1f", float64(ans)/2)
}

//func main() { cf314D(bufio.NewReader(os.Stdin), os.Stdout) }
