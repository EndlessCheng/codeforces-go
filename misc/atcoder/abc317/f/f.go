package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var n, b1, b2, b3 int
	Fscan(in, &n, &b1, &b2, &b3)

	s := strconv.FormatUint(uint64(n), 2)

	const limit = 10
	dp := make([][limit][limit][limit][2][2][2][2][2][2]int, len(s))
	vis := make([][limit][limit][limit][2][2][2][2][2][2]bool, len(s))

	var f func(int, int, int, int, int, int, int, int, int, int) int
	f = func(p, a1, a2, a3 int, lim1, lim2, lim3, z1, z2, z3 int) (res int) {
		if p == len(s) {
			if a1 == 0 && a2 == 0 && a3 == 0 && z1 > 0 && z2 > 0 && z3 > 0 {
				return 1
			}
			return 0
		}

		dv := &dp[p][a1][a2][a3][lim1][lim2][lim3][z1][z2][z3]
		if vis[p][a1][a2][a3][lim1][lim2][lim3][z1][z2][z3] {
			return *dv
		}
		vis[p][a1][a2][a3][lim1][lim2][lim3][z1][z2][z3] = true
		defer func() { *dv = res }()

		up1 := 1
		if lim1 > 0 {
			up1 = int(s[p] - '0')
		}
		up2 := 1
		if lim2 > 0 {
			up2 = int(s[p] - '0')
		}
		up3 := 1
		if lim3 > 0 {
			up3 = int(s[p] - '0')
		}
		for d1 := 0; d1 <= up1; d1++ {
			for d2 := 0; d2 <= up2; d2++ {
				d3 := d1 ^ d2
				if lim3 > 0 && d3 > up3 {
					continue
				}
				l1 := 0
				if lim1 > 0 && d1 == up1 {
					l1 = 1
				}
				l2 := 0
				if lim2 > 0 && d2 == up2 {
					l2 = 1
				}
				l3 := 0
				if lim3 > 0 && d3 == up3 {
					l3 = 1
				}
				cnt := f(p+1, (a1<<1|d1)%b1, (a2<<1|d2)%b2, (a3<<1|d3)%b3, l1, l2, l3, z1|d1, z2|d2, z3|d3)
				res += cnt
			}
		}
		res %= mod
		return
	}
	ans := f(0, 0, 0, 0, 1, 1, 1, 0, 0, 0)
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
