package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF940E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, c, v int
	Fscan(in, &n, &c)
	st := make([][17]int, n)
	s := int64(0)
	for i := range st {
		Fscan(in, &v)
		st[i][0] = v
		s += int64(v)
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }

	dp := make([]int64, n+1)
	for r := c; r <= n; r++ {
		dp[r] = max(dp[r-1], dp[r-c]+int64(q(r-c, r)))
	}
	Fprint(out, s-dp[n])
}

//func main() { CF940E(os.Stdin, os.Stdout) }
