package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1580A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]string, n)
		sum := make([][]int, n+1)
		sum[0] = make([]int, m+1)
		for i := range a {
			Fscan(in, &a[i])
			sum[i+1] = make([]int, m+1)
			for j, v := range a[i] {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + int(v&1)
			}
		}
		query := func(r1, c1, r2, c2 int) int { r2++; c2++; return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1] }

		const minR, minC, inf = 5, 4, 99
		ans := inf
		for i := minR - 1; i < n; i++ {
			for j := i - minR + 1; j >= 0; j-- {
				mi := inf
				for k := minC - 1; k < m; k++ {
					mi = query(j+1, k-1, i-1, k-1) + int((a[j][k-1]&1^1)+(a[i][k-1]&1^1)) + // 第 k-1 列
						min(mi, query(j+1, k-minC+2, i-1, k-2)+ // 第 k-2 至 k-minC+2 列的中间部分
							minC-3-query(j, k-minC+2, j, k-2)+ // 的第 j 行
							minC-3-query(i, k-minC+2, i, k-2)+ // 的第 i 行
							i-j-1-query(j+1, k-minC+1, i-1, k-minC+1)) // 第 k-minC+1 列（去掉两角）
					ans = min(ans, mi+i-j-1-query(j+1, k, i-1, k))     // 第 k 列（去掉两角）
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1580A(os.Stdin, os.Stdout) }
