package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	next := []byte{'Q': 'W', 'W': 'E', 'E': 'R', 'R': 'Q'}

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	// 记忆化，根据题意模拟
	// 由于路径可能会形成环，可以额外用一个 inStk 数组来标记该点是否在递归栈中，若访问到一个在栈中的点，则说明有环
	dp := make([][]int, n)
	inStk := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
		inStk[i] = make([]bool, m)
	}
	var f func(int, int) int
	f = func(x, y int) (res int) {
		dv := &dp[x][y]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if inStk[x][y] {
			return -2
		}
		inStk[x][y] = true
		for _, d := range dir4 {
			if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m {
				if a[xx][yy] == next[a[x][y]] {
					r := f(xx, yy)
					if r == -2 {
						return -2
					}
					res = max(res, r)
				}
			}
		}
		inStk[x][y] = false
		return res + 1
	}
	for i, r := range a {
		for j, b := range r {
			if b == 'Q' {
				res := f(i, j)
				if res == -2 {
					Fprint(out, "infinity")
					return
				}
				ans = max(ans, res)
			}
		}
	}
	if ans < 4 {
		Fprint(out, "none")
	} else {
		Fprint(out, ans/4)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
