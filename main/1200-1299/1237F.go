package main

import (
	. "fmt"
	"io"
)

func cf1237F(in io.Reader, out io.Writer) {
	const mod = 998244353
	const mx = 3600
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
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	perm := func(n, k int) int {
		return F[n] * invF[n-k] % mod
	}

	var n, m, k, r1, c1, r2, c2, ans int
	Fscan(in, &n, &m, &k)
	banR := make([]bool, n)
	banC := make([]bool, m)
	for ; k > 0; k-- {
		Fscan(in, &r1, &c1, &r2, &c2)
		banR[r1-1] = true
		banR[r2-1] = true
		banC[c1-1] = true
		banC[c2-1] = true
	}

	calc := func(ban []bool) ([]int, int) {
		n := len(ban)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n/2+1)
			f[i][0] = 1
		}
		for i := 1; i < n; i++ {
			for j := 1; j <= (i+1)/2; j++ {
				f[i+1][j] = f[i][j]
				if !ban[i] && !ban[i-1] {
					f[i+1][j] = (f[i+1][j] + f[i-1][j-1]) % mod
				}
			}
		}
		empty := 0
		for _, b := range ban {
			if !b {
				empty++
			}
		}
		return f[n], empty
	}

	f, emptyR := calc(banR)
	g, emptyC := calc(banC)
	for i, v := range f { // i 个竖放
		for j, w := range g { // j 个横放
			if j > emptyR-i*2 || i > emptyC-j*2 {
				break
			}
			ans = (ans + v*w%mod*perm(emptyR-i*2, j)%mod*perm(emptyC-j*2, i)) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { cf1237F(bufio.NewReader(os.Stdin), os.Stdout) }
