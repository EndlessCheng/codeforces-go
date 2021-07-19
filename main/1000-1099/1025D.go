package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1025D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	ok := make([][]bool, n)
	for i := range a {
		Fscan(in, &a[i])
		ok[i] = make([]bool, n)
	}
	for i, v := range a {
		for j, w := range a[:i] {
			if gcd(v, w) > 1 {
				ok[i][j] = true
				ok[j][i] = true
			}
		}
	}

	dp := make([][][2]int8, n)
	for i := range dp {
		dp[i] = make([][2]int8, n)
		for j := range dp[i] {
			dp[i][j] = [2]int8{-1, -1}
		}
	}
	var f func(int, int, int) int8
	f = func(l, r, side int) (res int8) {
		if l > r {
			return 1
		}
		dv := &dp[l][r][side]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := l; i <= r; i++ {
			if side == 0 && (l == 0 || ok[l-1][i]) && f(l, i-1, 1) > 0 && f(i+1, r, 0) > 0 ||
				side == 1 && (r == n-1 || ok[r+1][i]) && f(l, i-1, 1) > 0 && f(i+1, r, 0) > 0 {
				return 1
			}
		}
		return
	}
	if f(0, n-1, 0) > 0 {
		Fprint(out, "Yes")
	} else {
		Fprint(out, "No")
	}
}

//func main() { CF1025D(os.Stdin, os.Stdout) }
