package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1051D(in io.Reader, out io.Writer) {
	const mod = 998244353
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, k int
	Fscan(in, &n, &k)
	if k == 1 {
		Fprint(out, 2)
		return
	}
	f := make([][2]uint, k+1)
	f[1] = [2]uint{2, 0}
	f[2] = [2]uint{0, 2}
	for i := 2; i <= n; i++ {
		for j := min(k, i*2); j > 1; j-- {
			f[j][0] = (f[j][0] + f[j][1]*2 + f[j-1][0]) % mod
			f[j][1] = (f[j][1] + f[j-1][0]*2 + f[j-2][1]) % mod
		}
	}
	Fprint(out, (f[k][0]+f[k][1])%mod)
}

//func main() { CF1051D(os.Stdin, os.Stdout) }
