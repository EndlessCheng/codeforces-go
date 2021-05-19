package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF401D(in io.Reader, out io.Writer) {
	var num []byte
	var m int
	Fscan(in, &num, &m)
	n, dup, cnt := len(num), int64(1), [10]int64{}
	for i := range num {
		num[i] &= 15
		cnt[num[i]]++
		dup *= cnt[num[i]]
	}

	dp := make([][]int64, 1<<n)
	for i := range dp {
		dp[i] = make([]int64, m)
	}
	dp[0][0] = 1
	for s, dv := range dp {
		for t := s ^ (1<<n - 1); t > 0; {
			lb := t & -t
			if p := bits.TrailingZeros(uint(lb)); s > 0 || num[p] > 0 {
				for j, v := range dv {
					dp[s|lb][(j*10+int(num[p]))%m] += v
				}
			}
			t ^= lb
		}
	}
	Fprint(out, dp[1<<n-1][0]/dup)
}

//func main() { CF401D(os.Stdin, os.Stdout) }
