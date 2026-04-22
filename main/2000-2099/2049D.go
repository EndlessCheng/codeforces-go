package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2049D(in io.Reader, out io.Writer) {
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		f := make([]int, m)
		for i := 1; i < m; i++ {
			f[i] = 1e18
		}
		a := make([]int, m)
		for range n {
			for i := range a {
				Fscan(in, &a[i])
			}

			nf := make([]int, m)
			for i := range nf {
				nf[i] = 1e18
			}
			for shift := range m {
				s := int(1e18)
				for j := range m {
					s = min(s, f[j]) + a[(j+shift)%m]
					nf[j] = min(nf[j], s+shift*k)
				}
			}
			f = nf
		}
		Fprintln(out, f[m-1])
	}
}

//func main() { cf2049D(bufio.NewReader(os.Stdin), os.Stdout) }
