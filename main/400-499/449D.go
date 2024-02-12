package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF449D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod, mx = 1_000_000_007, 20
	var n, v int
	Fscan(in, &n)
	p2 := make([]int, n+1)
	p2[0] = 1
	dp := make([]int, 1<<mx)
	for i := 0; i < n; i++ {
		p2[i+1] = p2[i] << 1 % mod
		Fscan(in, &v)
		dp[v]++
	}
	for i := 0; i < mx; i++ {
		for s := 1<<mx - 1; s >= 0; s-- {
			if s>>i&1 == 0 {
				dp[s] += dp[s|1<<i]
			}
		}
	}
	ans := int64(0)
	for i, v := range dp {
		if bits.OnesCount(uint(i))&1 > 0 {
			ans -= int64(p2[v] - 1)
		} else {
			ans += int64(p2[v] - 1)
		}
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF449D(os.Stdin, os.Stdout) }
