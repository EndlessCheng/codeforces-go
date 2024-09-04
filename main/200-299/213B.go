package main

import (
	. "fmt"
	"io"
)

func cf213B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	const mx = 101
	C := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	var n, ans int
	Fscan(in, &n)
	a := [10]int{}
	for i := range a {
		Fscan(in, &a[i])
	}

	f := make([]int, n+1)
	for i := a[9]; i <= n; i++ {
		f[i] = 1
	}
	for j := 8; j > 0; j-- {
		for i := n; i >= 0; i-- {
			res := 0
			for k := a[j]; k <= i; k++ {
				res = (res + f[i-k]*C[i][k]) % mod
			}
			f[i] = res
		}
	}
	for i := 1; i <= n; i++ {
		for k := a[0]; k <= i; k++ {
			ans = (ans + f[i-k]*C[i-1][k]) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { cf213B(os.Stdin, os.Stdout) }
