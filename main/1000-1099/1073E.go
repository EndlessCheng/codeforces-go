package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF1073E(in io.Reader, out io.Writer) {
	const mod = 998244353
	var l, r int64
	var k int
	Fscan(in, &l, &r, &k)
	calc := func(s string) int64 {
		n := len(s)
		type pair struct{ cnt, sum int64 }
		dp := make([][1024]pair, n)
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = pair{-1, -1}
			}
		}
		var f func(int, uint16, bool, bool) pair
		f = func(p int, mask uint16, limitUp, hasD bool) (res pair) {
			if p == n {
				if !hasD {
					return
				}
				return pair{1, 0}
			}
			if !limitUp && hasD {
				dv := &dp[p][mask]
				if dv.cnt >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			up := 9
			if limitUp {
				up = int(s[p] & 15)
			}
			for ch := 0; ch <= up; ch++ {
				tmp := mask
				if hasD || ch > 0 {
					tmp |= 1 << ch
				}
				if bits.OnesCount16(tmp) <= k {
					pr := f(p+1, tmp, limitUp && ch == up, hasD || ch > 0)
					res.cnt = (res.cnt + pr.cnt) % mod
					res.sum = (res.sum + int64(math.Pow10(n-1-p))%mod*pr.cnt%mod*int64(ch) + pr.sum) % mod
				}
			}
			return
		}
		return f(0, 0, true, false).sum
	}
	ansLower := calc(strconv.FormatInt(l-1, 10))
	ansUpper := calc(strconv.FormatInt(r, 10))
	ans := ansUpper - ansLower
	ans = (ans%mod + mod) % mod
	Fprint(out, ans)
}

//func main() { CF1073E(os.Stdin, os.Stdout) }
