package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF757D(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	const mx = 20
	var n, ans int
	var s []byte
	Fscan(in, &n, &s)
	for i := range s {
		s[i] &= 1
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1<<mx)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, mask int) (res int) {
		if mask&(mask+1) == 0 && mask > 0 {
			res = 1
		}
		if p == n {
			return
		}
		dv := &dp[p][mask]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for j, v := p, byte(0); j < n; j++ {
			v = v<<1 | s[j]
			if v > mx {
				break
			}
			if v > 0 {
				res = (res + f(j+1, mask|1<<(v-1))) % mod
			}
		}
		return
	}
	for i := range s {
		ans = (ans + f(i, 0)) % mod
	}
	Fprint(out, ans)
}

//func main() { CF757D(os.Stdin, os.Stdout) }
