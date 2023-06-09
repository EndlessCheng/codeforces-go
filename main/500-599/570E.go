package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF570E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	if a[0][0] != a[n-1][m-1] {
		Fprint(out, 0)
		return
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+2)
	}
	f[1][n] = 1
	for i := 1; i < (n+m)/2; i++ {
		for r1 := min(n, i+1); r1 > 0 && r1 >= i+2-m; r1-- {
			c1 := i + 2 - r1
			for r2 := max(1, max(n-i, r1-2)); r2 <= n && r2 < n+m-i; r2++ {
				c2 := n + m - i - r2
				if a[r1-1][c1-1] == a[r2-1][c2-1] {
					f[r1][r2] = (((f[r1][r2]+f[r1][r2+1])%mod+f[r1-1][r2+1])%mod + f[r1-1][r2]) % mod
				} else {
					f[r1][r2] = 0
				}
			}
		}
	}

	ans := int64(0)
	if (n+m)%2 > 0 {
		for i := 1; i <= n; i++ {
			ans += int64(f[i][i] + f[i][i+1])
		}
	} else {
		for i := 1; i <= n; i++ {
			ans += int64(f[i][i])
		}
	}
	Fprint(out, ans%mod)
}

//func main() { CF570E(os.Stdin, os.Stdout) }
