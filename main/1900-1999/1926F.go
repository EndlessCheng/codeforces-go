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
	dp := [7][1 << 7][1 << 5]int{}
	for Fscan(in, &T); T > 0; T-- {
		a := [7]int{}
		for i := range a {
			Fscan(in, &s)
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
		f = func(i, pre7, pre5 int) (res int) {
			if i == 7 {
				return
			}
			p := &dp[i][pre7][pre5]
			if *p != -1 {
				return *p
			}
			res = 99
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
				for t, lb := cur>>1&31, 0; t > 0; t ^= lb {
					lb = t & -t
					if pre7&lb > 0 && pre7&(lb<<2) > 0 {
						cur5 |= lb
					}
				}
				r := f(i+1, cur, cur5)
				res = min(res, r+bits.OnesCount(uint(cur^a[i])))
				cur = (cur - 1) & a[i]
			}
			*p = res
			return
		}
		ans := 99
		for cur, ok := a[0], true; ok; ok = cur != a[0] {
			res := f(1, cur, 0) + bits.OnesCount(uint(cur^a[0]))
			ans = min(ans, res)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1926F(os.Stdin, os.Stdout) }
