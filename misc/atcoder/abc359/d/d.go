package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, k int
	var s string
	Fscan(in, &n, &k, &s)

	pal := make([]bool, 1<<k)
o:
	for i := range pal {
		for j := 0; j < k/2; j++ {
			if i>>j&1 != i>>(k-1-j)&1 {
				continue o
			}
		}
		pal[i] = true
	}

	mask := 1<<(k-1) - 1
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<k)
	}
	for j := range f[0] {
		f[0][j] = 1
	}
	for i, c := range s {
		for j := 0; j <= mask; j++ {
			res := 0
			for b := 0; b < 2; b++ {
				if c != '?' && int(c&1) != b {
					continue
				}
				t := j<<1 | b
				if i > n-k || !pal[t] {
					res += f[i][t&mask]
				}
			}
			f[i+1][j] = res % mod
		}
	}
	Fprint(out, f[n][0])
}

func main() { run(os.Stdin, os.Stdout) }
