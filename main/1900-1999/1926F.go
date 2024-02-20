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
	const MASK = 1<<(N-2) - 1
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
		f = func(i, pre7, pre5 int) int {
			if i == N {
				return 0
			}
			p := &dp[i][pre7][pre5]
			if *p != -1 {
				return *p
			}
			res := N * N
		o:
			for cur, ok := a[i], true; ok; ok = cur != a[i] {
				for t, lb := pre5, 0; t > 0; t ^= lb {
					lb = t & -t
					if cur&lb > 0 && cur&(lb<<2) > 0 {
						cur = (cur - 1) & a[i]
						continue o
					}
				}
				cur5 := 0
				for t, lb := cur>>1&MASK, 0; t > 0; t ^= lb {
					lb = t & -t
					if pre7&lb > 0 && pre7&(lb<<2) > 0 {
						cur5 |= lb
					}
				}
				res = min(res, f(i+1, cur, cur5)+bits.OnesCount(uint(cur^a[i])))
				cur = (cur - 1) & a[i]
			}
			*p = res
			return res
		}
		ans := N * N
		for cur, ok := a[0], true; ok; ok = cur != a[0] {
			ans = min(ans, f(1, cur, 0)+bits.OnesCount(uint(cur^a[0])))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1926F(os.Stdin, os.Stdout) }
