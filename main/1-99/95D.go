package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func cf95D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, k int
	var s []byte
	Fscan(in, &T, &k)
	// todo 去掉 ok，改成 pow10 数组提前返回
	dp := make([][1001][2]int, k+1)
	var f func(int, int, int, bool) int
	f = func(i, left, ok int, isLimit bool) (res int) {
		if i < 0 {
			return ok
		}
		if !isLimit {
			p := &dp[left][i][ok]
			if *p > 0 {
				return *p - 1
			}
			defer func() { *p = res + 1 }()
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := range up + 1 {
			if d == 4 || d == 7 {
				newOk := 0
				if ok > 0 || left > 0 {
					newOk = 1
				}
				res += f(i-1, k, newOk, isLimit && d == up)
			} else {
				res += f(i-1, max(left-1, 0), ok, isLimit && d == up)
			}
		}
		res %= mod
		return
	}

	for range T {
		Fscan(in, &s)
		slices.Reverse(s)
		ans := mod - f(len(s)-1, 0, 0, true)

		pre := -k - 1
		for i, b := range s {
			if b == '4' || b == '7' {
				if i-pre <= k {
					ans++
					break
				}
				pre = i
			}
		}

		Fscan(in, &s)
		slices.Reverse(s)
		ans += f(len(s)-1, 0, 0, true)
		Fprintln(out, ans%mod)
	}
}

func main() { cf95D(bufio.NewReader(os.Stdin), os.Stdout) }
