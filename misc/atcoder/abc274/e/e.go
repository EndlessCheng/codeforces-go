package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]struct{ x, y int }, n+m)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	dp := make([][]float64, 1<<(n+m))
	for i := range dp {
		dp[i] = make([]float64, n+m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dis := make([][]float64, n+m)
	for i := range dis {
		dis[i] = make([]float64, n+m)
		for j := range dis[i] {
			dis[i][j] = math.Sqrt(float64((a[i].x-a[j].x)*(a[i].x-a[j].x) + (a[i].y-a[j].y)*(a[i].y-a[j].y)))
		}
	}

	full := 1<<n - 1
	var f func(int, int) float64
	f = func(mask, i int) (res float64) {
		dv := &dp[mask][i]
		if *dv > -1 {
			return *dv
		}
		defer func() { *dv = res }()
		spd := float64(int(1) << bits.OnesCount(uint(mask>>n)))
		if mask&full == full {
			res = math.Sqrt(float64(a[i].x*a[i].x+a[i].y*a[i].y)) / spd
			for j := n; j < n+m; j++ {
				if mask>>j&1 == 0 {
					d2 := f(mask|1<<j, j) + dis[i][j]/spd
					res = min(res, d2)
				}
			}
			return
		}
		res = 1e18
		for j := 0; j < n+m; j++ {
			if mask>>j&1 == 0 {
				r := f(mask|1<<j, j) + dis[i][j]/spd
				res = min(res, r)
			}
		}
		return
	}
	ans := 1e18
	for i := range a {
		r := f(1<<i, i) + math.Sqrt(float64(a[i].x*a[i].x+a[i].y*a[i].y))
		ans = min(ans, r)
	}
	Fprintf(out, "%.10f\n", ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
