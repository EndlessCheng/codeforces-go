package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1152D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// https://www.luogu.com.cn/problemnew/solution/CF1152D
	const mod int = 1e9 + 7
	var n int
	Fscan(in, &n)
	dp := [1001][1001]int{}
	dp[0][0] = 1
	ans := 0
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0]
		if i&1 == 1 {
			ans = (ans + dp[i][0]) % mod
		}
		for j := 1; j <= i; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
			if (i+j)&1 == 1 {
				ans = (ans + dp[i][j]) % mod
			}
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF1152D(os.Stdin, os.Stdout)
//}
