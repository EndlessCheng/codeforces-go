package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n int
	var p, ans float64
	Fscan(in, &n)
	f := make([]float64, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		Fscan(in, &p)
		for j := i; j > 0; j-- {
			f[j] = f[j-1]*p + f[j]*(1-p)
		}
		f[0] *= 1 - p
	}
	for _, v := range f[n/2+1:] {
		ans += v
	}
	Fprintf(out, "%.9f", ans)
}

func main() { run(os.Stdin, os.Stdout) }
