package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1341D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	ds := [10]string{"1110111", "0010010", "1011101", "1011011", "0111010", "1101011", "1101111", "1010010", "1111111", "1111011"}

	var n, k int
	var s []byte
	Fscan(in, &n, &k)
	a := make([][10]int, n)
	for i := range a {
		Fscan(in, &s)
		for j, d := range ds {
			for k, b := range s {
				if b > d[k] {
					a[i][j] = 1e9
					break
				}
				a[i][j] += int(d[k] - b)
			}
		}
	}

	dp := make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	ans := make([]byte, n)
	var f func(int, int) int8
	f = func(p, left int) (res int8) {
		if p >= n {
			if left > 0 {
				return
			}
			return 1
		}
		v := &dp[p][left]
		if *v >= 0 {
			return *v
		}
		defer func() { *v = res }()
		for i := 9; i >= 0; i-- {
			if a[p][i] <= left {
				ans[p] = byte('0' + i)
				if f(p+1, left-a[p][i]) == 1 {
					return 1
				}
			}
		}
		return
	}
	if f(0, k) == 0 {
		Fprint(_w, -1)
	} else {
		Fprint(_w, string(ans))
	}
}

//func main() { CF1341D(os.Stdin, os.Stdout) }
