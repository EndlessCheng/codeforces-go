package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF743E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const mx = 8
	const m = 1 << mx

	var n, v int
	Fscan(in, &n)
	pos := [mx][]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		pos[v-1] = append(pos[v-1], i)
	}

	minC := n
	for _, ps := range pos {
		minC = min(minC, len(ps))
	}
	// 注：更快的做法是二分 c
	for c := minC + 1; ; c-- { // 当前枚举的每个数的个数必须是 c 或 c-1
		// 状态定义为连续元素为 mask 下，最后一个元素的下标的最小值
		// 第二维度表示 c-1 的出现次数
		dp := make([][mx]int, m)
		for i := 1; i < m; i++ {
			for j := range dp[i] {
				dp[i][j] = 1e9
			}
		}
		for s, dv := range dp {
			for t, lb := m-1^s, 0; t > 0; t ^= lb {
				lb = t & -t
				ss := s | lb
				ps := pos[bits.TrailingZeros(uint(lb))]
				for k, p := range dv {
					i := sort.SearchInts(ps, p+1) + c - 1 // 取 c 个数后的位置
					if i-1 < len(ps) {
						if i < len(ps) {
							dp[ss][k] = min(dp[ss][k], ps[i])
						}
						if k+1 < mx {
							// 取 c-1 个数的位置，特判 c=1 的情况
							if c == 1 {
								dp[ss][k+1] = min(dp[ss][k+1], p)
							} else {
								dp[ss][k+1] = min(dp[ss][k+1], ps[i-1])
							}
						}
					}
				}
			}
		}
		for i, p := range dp[m-1] {
			if p <= n {
				Fprint(out, c*mx-i)
				return
			}
		}
	}
}

//func main() { CF743E(os.Stdin, os.Stdout) }
