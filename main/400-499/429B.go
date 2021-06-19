package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF429B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	const mx = 1000
	var n, m, ans int
	var a, dp1, dp2, dp3, dp4 [mx + 2][mx + 2]int
	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			Fscan(in, &a[i][j])
		}
	}
	// A 从起点到相遇点
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp1[i][j] = a[i][j] + max(dp1[i-1][j], dp1[i][j-1])
		}
	}
	// A 从终点到相遇点
	for i := n; i > 0; i-- {
		for j := m; j > 0; j-- {
			dp2[i][j] = a[i][j] + max(dp2[i+1][j], dp2[i][j+1])
		}
	}
	// B 从起点到相遇点
	for i := n; i > 0; i-- {
		for j := 1; j <= m; j++ {
			dp3[i][j] = a[i][j] + max(dp3[i+1][j], dp3[i][j-1])
		}
	}
	// B 从终点到相遇点
	for i := 1; i <= n; i++ {
		for j := m; j > 0; j-- {
			dp4[i][j] = a[i][j] + max(dp4[i-1][j], dp4[i][j+1])
		}
	}
	for i := 2; i < n; i++ {
		for j := 2; j < m; j++ {
			ans = max(max(ans, dp1[i][j-1]+dp2[i][j+1]+dp3[i+1][j]+dp4[i-1][j]), max(ans, dp1[i-1][j]+dp2[i+1][j]+dp3[i][j-1]+dp4[i][j+1]))
		}
	}
	Fprint(out, ans)
}

//func main() { CF429B(os.Stdin, os.Stdout) }
