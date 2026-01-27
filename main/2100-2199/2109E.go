package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2109E(in io.Reader, out io.Writer) {
	const mod = 998244353
	const mx = 501
	C := [mx][mx]int{}
	for i := range C {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, k+1)
		}
		f[n][0] = 1
		for i := n - 1; i >= 0; i-- {
			for j, fj := range f[i+1] {
				op := j
				if s[i] == '0' {
					op++
				}
				for l := range min(op, k-j) + 1 {
					f[i][j+l] = (f[i][j+l] + fj*C[(op+l)/2][l]) % mod
				}
			}
		}
		Fprintln(out, f[0][k])
	}
}

//func main() { cf2109E(bufio.NewReader(os.Stdin), os.Stdout) }
