package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2081A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	const inv2 = (mod + 1) / 2
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt1 := make([]int, n)
		c1 := 0
		for i, b := range s {
			if b == '0' {
				c1 = 0
			} else {
				c1++
				cnt1[i] = c1
			}
		}

		dp := make([][2]int, n)
		for i := range dp {
			dp[i] = [2]int{-1, -1}
		}
		var f func(int, int) int
		f = func(i, j int) (res int) {
			if i <= 0 {
				return
			}
			p := &dp[i][j]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
			res = f(i-1, 0) + 1
			if j == 0 && s[i] == '0' {
				return
			}
			c := cnt1[i-j] + j
			res2 := f(i-c, 1) + c
			return (res + res2) * inv2 % mod
		}
		Fprintln(out, f(n-1, 0))
	}
}

//func main() { cf2081A(bufio.NewReader(os.Stdin), os.Stdout) }
