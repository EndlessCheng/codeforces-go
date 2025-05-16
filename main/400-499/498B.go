package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf498B(in io.Reader, out io.Writer) {
	var n, T, t int
	var p, ans float64
	Fscan(in, &n, &T)
	f := make([]float64, T+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		Fscan(in, &p, &t)
		p /= 100
		np := math.Pow(1-p, float64(t-1))
		np2 := np * (1 - p)
		nf := make([]float64, T+1)
		for j := i; j <= T; j++ {
			d := 0.
			if j > t {
				d = f[j-t-1] * np
			}
			nf[j] = (nf[j-1]-d)*(1-p) + f[j-1]*p
			if j >= t {
				nf[j] += f[j-t] * np2
			}
			ans += nf[j]
		}
		f = nf
	}
	Fprintf(out, "%.9f", ans)
}

//func main() { cf498B(bufio.NewReader(os.Stdin), os.Stdout) }
