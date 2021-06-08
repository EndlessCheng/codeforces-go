package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 另一种做法是 k=2 时答案肯定是 ABAB... 与 BABA... 中的一个；
// k=3 时若当前字母与左侧相邻字母不同，则修改其为与左右相邻字母均不同的字母

// github.com/EndlessCheng/codeforces-go
func CF219C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, miJ int
	var s string
	Fscan(in, &n, &k, &s)
	dp := make([]int, k)
	fa := make([][26]int, n)
	for i, b := range s {
		b := int(b - 'A')
		dp2 := make([]int, k)
		for j := range dp2 {
			dp2[j] = 1e9
			add := 0
			if j != b {
				add = 1
			}
			for k, v := range dp {
				if k != j && v+add < dp2[j] {
					dp2[j] = v + add
					fa[i][j] = k
				}
			}
		}
		dp = dp2
	}
	for j, v := range dp {
		if v < dp[miJ] {
			miJ = j
		}
	}
	Fprintln(out, dp[miJ])
	ans := make([]byte, n)
	for i, j := n-1, miJ; i >= 0; i-- {
		ans[i] = 'A' + byte(j)
		j = fa[i][j]
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF219C(os.Stdin, os.Stdout) }
