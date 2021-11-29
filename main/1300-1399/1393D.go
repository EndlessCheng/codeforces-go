package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1393D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := m + n*2 - 2
	pre := make([][26]int, m)
	for j, b := range a[0] {
		pre[j][b-'a'] = 1
	}
	for i := 1; i < n; i++ {
		cur := make([][26]int, m)
		cur[0][a[i][0]-'a'] = 1
		for j := 1; j < m-1; j++ {
			b := a[i][j] - 'a'
			cur[j][b] = min(min(pre[j-1][b], pre[j][b]), pre[j+1][b]) // 视作上面三个菱形的并
			if p := i - cur[j][b]*2; p >= 0 && a[p][j]-'a' == b {
				cur[j][b]++
			}
			ans += cur[j][b]
		}
		cur[m-1][a[i][m-1]-'a'] = 1
		pre = cur
	}
	Fprint(out, ans)
}

//func main() { CF1393D(os.Stdin, os.Stdout) }
