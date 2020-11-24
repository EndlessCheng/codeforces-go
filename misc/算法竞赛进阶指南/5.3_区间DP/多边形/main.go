package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9 // 1e18
	type pair struct{ min, max int }

	var n int
	Fscan(in, &n)
	op := make([]string, 2*n)
	a := make([]int, 2*n)
	for i := 0; i < n; i++ {
		Fscan(in, &op[i], &a[i])
	}
	copy(op[n:], op[:n])
	copy(a[n:], a[:n])

	m := 2 * n
	dp := make([][]pair, m)
	for i := range dp {
		dp[i] = make([]pair, m)
		for j := range dp[i] {
			dp[i][j] = pair{inf, -inf}
		}
	}
	var f func(int, int) pair
	f = func(l, r int) (res pair) {
		if l == r {
			return pair{a[l], a[l]}
		}
		dv := &dp[l][r]
		if dv.min != inf {
			return *dv
		}
		defer func() { *dv = res }()
		res = *dv
		for i := l + 1; i <= r; i++ { // 枚举右区间左端点
			p, q := f(l, i-1), f(i, r)
			if op[i][0] == 't' { // add
				res.min = min(res.min, p.min+q.min)
				res.max = max(res.max, p.max+q.max)
			} else { // mul
				a, b, c, d := p.min*q.min, p.min*q.max, p.max*q.min, p.max*q.max
				res.min = min(res.min, a, b, c, d)
				res.max = max(res.max, a, b, c, d)
			}
		}
		return
	}
	ans := []int{}
	mx := -inf
	for i := 0; i < n; i++ {
		if v := f(i, i+n-1).max; v > mx {
			mx = v
			ans = []int{i}
		} else if v == mx {
			ans = append(ans, i)
		}
	}
	Fprintln(out, mx)
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
