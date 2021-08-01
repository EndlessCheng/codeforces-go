package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1234F(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var s string
	Fscan(bufio.NewReader(in), &s)

	const mx = 20
	dp := make([]int, 1<<mx)
	for i := range s {
		x := 0
		for j := i; j >= 0 && x>>(s[j]-'a')&1 == 0; j-- {
			x |= 1 << (s[j] - 'a')
			dp[x] = i - j + 1
		}
	}

	for i := 0; i < mx; i++ {
		for s := 0; s < 1<<mx; s++ {
			s |= 1 << i
			dp[s] = max(dp[s], dp[s^1<<i])
		}
	}

	ans := 0
	for s, v := range dp {
		ans = max(ans, v+dp[1<<mx-1^s])
	}
	Fprint(out, ans)
}

//func main() { CF1234F(os.Stdin, os.Stdout) }
