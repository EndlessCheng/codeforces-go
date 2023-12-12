package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf294C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	const mx int = 1000
	F := make([]int, mx+1)
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := make([]int, mx+1)
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	ans := F[n-m] * invF[a[0]-1] % mod * invF[n-a[m-1]] % mod
	e := 0
	for i := 1; i < m; i++ {
		d := a[i] - a[i-1] - 1
		if d > 0 {
			ans = ans * invF[d] % mod
			e += d - 1
		}
	}
	Fprint(out, ans*pow(2, e)%mod)
}

//func main() { cf294C(os.Stdin, os.Stdout) }
