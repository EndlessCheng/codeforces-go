package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1272F(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	b2i := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)
	s += "$"
	t += "$"

	memo := make([][][]int, n+1)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n+m+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(i, j, k int) int {
		if k > n+m {
			return 1e9
		}
		if i == n && j == m && k == 0 {
			return 0
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		*p = 1e9 // 相当于 vis，防止 k +1 -1 死循环
		res := f(i+b2i(s[i] == '('), j+b2i(t[j] == '('), k+1) + 1
		if k > 0 {
			res = min(res, f(i+b2i(s[i] == ')'), j+b2i(t[j] == ')'), k-1)+1)
		}
		*p = res
		return res
	}

	ans := []byte{}
	var makeAns func(int, int, int)
	makeAns = func(i, j, k int) {
		if i == n && j == m && k == 0 {
			return
		}
		if f(i+b2i(s[i] == '('), j+b2i(t[j] == '('), k+1)+1 == f(i, j, k) {
			ans = append(ans, '(')
			makeAns(i+b2i(s[i] == '('), j+b2i(t[j] == '('), k+1)
		} else {
			ans = append(ans, ')')
			makeAns(i+b2i(s[i] == ')'), j+b2i(t[j] == ')'), k-1)
		}
	}
	makeAns(0, 0, 0)
	Fprintf(out, "%s", ans)
}

//func main() { CF1272F(os.Stdin, os.Stdout) }
