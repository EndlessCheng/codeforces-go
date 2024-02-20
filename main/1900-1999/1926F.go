package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1926F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T int
	var s string
	const N = 7
	a := [N]int{}
	dp := [N][1 << N][1 << (N - 2)]int{}
	for Fscan(in, &T); T > 0; T-- {
		for i := range a {
			Fscan(in, &s)
			a[i] = 0
			for j, b := range s {
				a[i] |= int(b&1^1) << j
			}
		}
		for i := range dp {
			for j := range dp[i] {
				for k := range dp[i][j] {
					dp[i][j][k] = -1
				}
			}
		}
		var f func(int, int, int) int
		f = func(i, pre, pre5 int) int {
			if i == N {
				return 0
			}
			p := &dp[i][pre][pre5]
			if *p != -1 {
				return *p
			}
			res := N * N
			for cur, ok := a[i], true; ok; ok = cur != a[i] {
				if cur>>2&cur&pre5 == 0 {
					cur5 := pre >> 2 & pre & (cur >> 1)
					res = min(res, f(i+1, cur, cur5)+bits.OnesCount(uint(cur^a[i])))
				}
				cur = (cur - 1) & a[i]
			}
			*p = res
			return res
		}
		Fprintln(out, f(0, 0, 0))
	}
}

//func main() { cf1926F(os.Stdin, os.Stdout) }
