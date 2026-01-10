package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf98E(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)

	mx := max(n, m)
	dp := make([][]float64, mx+1)
	for i := range dp {
		dp[i] = make([]float64, mx+1)
	}
	var f func(int, int) float64
	f = func(n, m int) float64 {
		if n == 0 {
			return 1 / float64(m+1)
		}
		if m == 0 {
			return 1
		}
		p := &dp[n][m]
		if *p != 0 {
			return *p
		}
		a := (1 - f(m-1, n)) * float64(m) / float64(m+1)
		b := 1/float64(m+1) + a
		c := 1.
		d := 1 - f(m, n-1)
		fix := (d - c) / (a - b - c + d)
		*p = fix*a + (1-fix)*c
		return *p
	}
	ans := f(n, m)
	Fprintf(out, "%.9f %.9f", ans, 1-ans)
}

//func main() { cf98E(bufio.NewReader(os.Stdin), os.Stdout) }
