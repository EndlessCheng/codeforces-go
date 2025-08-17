package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minCost(grid [][]int, k int) int {
	n := len(grid[0])
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}

	sufMinF := make([]int, mx+2)
	for i := range sufMinF {
		sufMinF[i] = math.MaxInt
	}
	minF := make([]int, mx+1)
	f := make([]int, n+1)

	for range k + 1 {
		for i := range minF {
			minF[i] = math.MaxInt
		}

		// 64. 最小路径和（空间优化写法）
		for i := range f {
			f[i] = math.MaxInt / 2
		}
		f[1] = -grid[0][0] // 起点的成本不算
		for _, row := range grid {
			for j, x := range row {
				f[j+1] = min(f[j]+x, f[j+1]+x, sufMinF[x])
				minF[x] = min(minF[x], f[j+1])
			}
		}

		tmp := slices.Clone(sufMinF)
		// 计算 minF 的后缀最小值
		for i := mx; i >= 0; i-- {
			sufMinF[i] = min(sufMinF[i+1], minF[i])
		}
		if slices.Equal(sufMinF, tmp) {
			// 收敛了：传送一次不改变 sufMinF，那么无论传送多少次都不会改变 sufMinF
			break
		}
	}
	return f[n]
}

func main() {
	m, n := 80, 80
	const u = 10000
	rg := testutil.NewRandGenerator()
	s := 0
	for range u {
		a := make([][]int, m)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				a[i][j] = rg.IntOnly(0, 1e4)
			}
		}
		k := minCost(a, -1)
		s += k
	}
	fmt.Println(float64(s) / float64(u))
}
