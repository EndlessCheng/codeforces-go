package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2145D(in io.Reader, out io.Writer) {
	const mx = 30
	f := [mx + 1][mx*(mx+1)/2 + 1]bool{0: {true}}
	for i := 1; i <= mx; i++ {
		for sz := 1; sz <= i; sz++ {
			j := i - sz
			for v, ok := range f[j][:j*(j+1)/2+1] {
				if ok {
					f[i][v+sz*(sz+1)/2] = true
				}
			}
		}
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		k = n*(n+1)/2 - k
		if !f[n][k] {
			Fprintln(out, 0)
			continue
		}
		for n > 0 {
			for sz := 1; sz <= n; sz++ {
				if f[n-sz][k-sz*(sz+1)/2] {
					for i := n - sz + 1; i <= n; i++ {
						Fprint(out, i, " ")
					}
					n -= sz
					k -= sz * (sz + 1) / 2
					break
				}
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2145D(os.Stdin, os.Stdout) }
