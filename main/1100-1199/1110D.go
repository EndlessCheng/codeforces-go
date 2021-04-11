package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1110D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, v int
	Fscan(in, &n, &m)
	c := make([]int, m+3)
	for ; n > 0; n-- {
		Fscan(in, &v)
		c[v]++
	}
	dp := make([][3][3]int, m+1)
	for i := 1; i <= m; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if c[i] >= j+k+l && c[i+1] >= k+j && c[i+2] >= k {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][l][j]+k+(c[i]-j-k-l)/3)
					}
				}
			}
		}
	}
	Fprint(out, dp[m][0][0])
}

//func main() { CF1110D(os.Stdin, os.Stdout) }
