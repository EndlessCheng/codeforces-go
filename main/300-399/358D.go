package _00_399

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF358D(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][3]int, n)
	for j := 0; j < 3; j++ {
		for i := range a {
			Fscan(in, &a[i][j])
		}
	}
	dp := make([][2]int, n+1)
	dp[0][0] = -1e9
	for i, v := range a {
		dp[i+1][0] = max(dp[i][0]+v[1], dp[i][1]+v[0])
		dp[i+1][1] = max(dp[i][0]+v[2], dp[i][1]+v[1])
	}
	Fprint(out, dp[n][0])
}

//func main() { CF358D(os.Stdin, os.Stdout) }
