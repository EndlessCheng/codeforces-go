package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2091F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, m, d int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &d)
		f0 := make([]int, m+1)
		f1 := make([]int, m+1)
		for i := range n {
			Fscan(in, &s)
			nf0 := make([]int, m+1)
			for j, b := range s {
				if b == '#' {
					nf0[j+1] = nf0[j]
				} else if i == 0 {
					nf0[j+1] = nf0[j] + 1
				} else {
					nf0[j+1] = (nf0[j] + f0[min(j+d, m)] - f0[max(j-d+1, 0)] + f1[min(j+d, m)] - f1[max(j-d+1, 0)]) % mod
				}
			}
			f0 = nf0
			for j, b := range s {
				if b == '#' {
					f1[j+1] = f1[j]
				} else {
					f1[j+1] = (f1[j] + f0[min(j+d+1, m)] - f0[j+1] + f0[j] - f0[max(j-d, 0)]) % mod
				}
			}
		}
		Fprintln(out, (f0[m]+f1[m]+mod*2)%mod)
	}
}

//func main() { cf2091F(bufio.NewReader(os.Stdin), os.Stdout) }
