package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p6669(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, k, n, m int
	for Fscan(in, &T, &k); T > 0; T-- {
		Fscan(in, &n, &m)
		m = min(m, n)
		a := [][2]int{}
		for ; n > 0; n /= k {
			a = append(a, [2]int{n % k, m % k})
			m /= k
		}
		type args struct {
			p                int
			g, l, limN, limM bool
		}
		memo := map[args]int{}
		var dfs func(int, bool, bool, bool, bool) int
		dfs = func(p int, greater, less, limN, limM bool) int {
			if p < 0 {
				if greater {
					return 1
				}
				return 0
			}
			t := args{p, greater, less, limN, limM}
			if v, ok := memo[t]; ok {
				return v
			}
			hiN, hiM := k-1, k-1
			if limN {
				hiN = a[p][0]
			}
			if limM {
				hiM = a[p][1]
			}
			res := 0
			for i := 0; i <= hiN; i++ {
				for j := 0; (less || j <= i) && j <= hiM; j++ {
					res += dfs(p-1, greater || j > i, less || j < i, limN && i == hiN, limM && j == hiM)
				}
			}
			res %= mod
			memo[t] = res
			return res
		}
		Fprintln(out, dfs(len(a)-1, false, false, true, true))
	}
}

//func main() { p6669(os.Stdin, os.Stdout) }
