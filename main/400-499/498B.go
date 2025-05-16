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
		nf := make([]float64, T+1)
		for j := i; j < t; j++ {
			nf[j] = nf[j-1]*(1-p) + f[j-1]*p
		}
		p1 := math.Pow(1-p, float64(t-1))
		p2 := math.Pow(1-p, float64(t))
		nf[t] = nf[t-1]*(1-p) + f[t-1]*p + f[0]*p2
		for j := t + 1; j <= T; j++ {
			nf[j] = (nf[j-1]-f[j-t-1]*p1*p-f[j-t-1]*p2)*(1-p) + f[j-1]*p + f[j-t]*p2
		}
		f = nf
		for _, v := range f {
			ans += v
		}
	}
	Fprintf(out, "%.9f", ans)
}

//func main() { cf498B(bufio.NewReader(os.Stdin), os.Stdout) }
