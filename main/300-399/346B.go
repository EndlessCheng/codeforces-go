package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF346B(in io.Reader, out io.Writer) {
	var s, t, virus []byte
	Fscan(bufio.NewReader(in), &s, &t, &virus)
	n, m, l := int8(len(s)), int8(len(t)), int8(len(virus))
	match := make([]int8, l)
	for i, c := int8(1), int8(0); i < l; i++ {
		v := virus[i]
		for c > 0 && virus[c] != v {
			c = match[c-1]
		}
		if virus[c] == v {
			c++
		}
		match[i] = c
	}

	dp := make([][][]int8, n+1)
	to := make([][][]int8, n+1)
	for i := range dp {
		dp[i] = make([][]int8, m+1)
		to[i] = make([][]int8, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int8, l)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
			to[i][j] = make([]int8, l)
		}
	}
	var f func(i, j, k int8) int8
	f = func(i, j, k int8) (res int8) {
		if i == n || j == m {
			return
		}
		dv := &dp[i][j][k]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = -1
		if r := f(i+1, j, k); r > res {
			res = r
			to[i][j][k] = -1
		}
		if r := f(i, j+1, k); r > res {
			res = r
			to[i][j][k] = -2
		}
		if v := s[i]; v == t[j] {
			c := k
			for c > 0 && virus[c] != v {
				c = match[c-1]
			}
			if virus[c] == v {
				c++
			}
			if c < l {
				if r := 1 + f(i+1, j+1, c); r > res {
					res = r
					to[i][j][k] = c
				}
			}
		}
		return
	}
	ans := f(0, 0, 0)
	if ans <= 0 {
		Fprint(out, 0)
		return
	}

	lcs := make([]byte, 0, ans)
	for i, j, k := int8(0), int8(0), int8(0); i < n && j < m; {
		if t := to[i][j][k]; t == -1 {
			i++
		} else if t == -2 {
			j++
		} else {
			lcs = append(lcs, s[i])
			i++
			j++
			k = t
		}
	}
	Fprintf(out, "%s", lcs)
}

//func main() { CF346B(os.Stdin, os.Stdout) }
