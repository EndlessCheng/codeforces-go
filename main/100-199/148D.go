package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF148D(_r io.Reader, _w io.Writer) {
	var W, B int
	Fscan(_r, &W, &B)
	dp := make([][][2]float64, W+1)
	for i := range dp {
		dp[i] = make([][2]float64, B+1)
		for j := range dp[i] {
			dp[i][j] = [2]float64{-1, -1}
		}
	}
	var f func(w, b, who int) float64
	f = func(w, b, who int) (res float64) {
		if w <= 0 {
			if who == 0 {
				return 0
			}
			return 1
		}
		if b <= 0 {
			return 1
		}
		dv := &dp[w][b][who]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if who == 0 {
			return (float64(w) + float64(b)*(1-f(w, b-1, 1))) / float64(w+b)
		}
		p := (float64(w)*f(w-1, b-1, 0) + float64(b-1)*f(w, b-2, 0)) / float64(w+b-1)
		return (float64(w) + float64(b)*(1-p)) / float64(w+b)
	}
	Fprintf(_w, "%.10f", f(W, B, 0))
}

//func main() { CF148D(os.Stdin, os.Stdout) }
