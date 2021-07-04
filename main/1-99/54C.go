package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF54C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	f := func(n int64) (res int64) {
		x, p10 := n, int64(1)
		for ; x >= 10; x /= 10 { // 正反算
			res += p10
			p10 *= 10
		}
		if x > 1 {
			res += p10
		} else if x == 1 {
			res += n - p10 + 1
		}
		return
	}

	var n, k int
	var l, r int64
	Fscan(in, &n)
	dp := make([]float64, n+1) // dp[i][j] 表示在前 i 个数中取到 j 个 1 开头的数的概率
	dp[0] = 1
	for i := 0; i < n; i++ {
		Fscan(in, &l, &r)
		p := float64(f(r)-f(l-1)) / float64(r-l+1)
		for j := n; j > 0; j-- {
			dp[j] = dp[j]*(1-p) + dp[j-1]*p
		}
		dp[0] *= 1 - p
	}
	Fscan(in, &k)
	ans := .0
	for _, v := range dp[(n*k+99)/100:] {
		ans += v
	}
	Fprintf(out, "%.15f\n", ans)
}

//func main() { CF54C(os.Stdin, os.Stdout) }
