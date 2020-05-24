package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 原题 https://www.luogu.com.cn/problem/P1541
func run(_r io.Reader, _w io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	cardCnts := [4]uint8{}
	for i := 0; i < m; i++ {
		var tp int
		Fscan(in, &tp)
		cardCnts[tp-1]++
	}
	dp := make([]map[[4]uint8]int, n)
	dp[0] = map[[4]uint8]int{cardCnts: 0}
	for i, v := range a {
		for used, s := range dp[i] {
			for step, left := range used {
				if left > 0 {
					newUsed := used
					newUsed[step]--
					if dp[i+step+1] == nil {
						dp[i+step+1] = map[[4]uint8]int{}
					}
					dp[i+step+1][newUsed] = max(dp[i+step+1][newUsed], s+v)
				}
			}
		}
	}
	ans := 0
	for _, s := range dp[n-1] {
		ans = max(ans, s)
	}
	Fprint(_w, ans+a[n-1])
}

func main() { run(os.Stdin, os.Stdout) }
