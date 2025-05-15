package main

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxScore(n int, edges [][]int) int64 {
	ans := (n*n*2 + n*5 - 6) * (n - 1) / 6
	if n == len(edges) { // 环
		ans += 2
	}
	return int64(ans)
}

func maxScoreOld(n int, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var cycle, chain []int
	var cntV, cntE int
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		cntV++
		cntE += len(g[x])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if b {
			continue
		}
		cntV, cntE = 0, 0
		dfs(i)
		if cntV*2 == cntE { // 环
			cycle = append(cycle, cntV)
		} else if cntE > 0 { // 链，但不考虑孤立点
			chain = append(chain, cntV)
		}
	}

	ans := 0
	cur := n
	f := func(sz int, isCycle bool) {
		l, r := cur-sz+1, cur
		for i := l; i < r-1; i++ {
			ans += i * (i + 2)
		}
		ans += r * (r - 1)
		if isCycle {
			ans += l * (l + 1)
		}
		cur -= sz
	}

	slices.Sort(cycle)
	for _, sz := range cycle {
		f(sz, true)
	}

	slices.SortFunc(chain, func(a, b int) int { return b - a })
	for _, sz := range chain {
		f(sz, false)
	}

	return int64(ans)
}

// 保证 size[i] >= 2, n >= sum(size)
func chainGreedy(size []int, n int) (ans int) {
	slices.Sort(size)
	tot := 0
	for _, v := range size {
		tot += v
	}

	low := n - tot + 1
	a := make([][]int, len(size))
	for i, sz := range size {
		a[i] = make([]int, sz)
		a[i][0] = low
		low++
		a[i][sz-1] = low
		low++
	}

	for p, sz := range size {
		i, j := 1, sz-2
		row := a[p]
		for i <= j {
			row[i] = low
			low++
			i++
			if i > j {
				break
			}
			row[j] = low
			low++
			j--
		}
	}

	for _, row := range a {
		for i := 1; i < len(row); i++ {
			v, w := row[i-1], row[i]
			ans += v * w
		}
	}
	return
}

func runAC(size []int, n int) (ans int) {
	size = slices.Clone(size)
	tot := 0
	ban := make([]bool, n+1)
	ban[0] = true
	for _, v := range size {
		tot += v
		ban[tot] = true
	}

	base := n - tot + 1

	dp := make([][]int, 1<<tot)
	for i := range dp {
		dp[i] = make([]int, tot)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}
	var f func(int, int) int
	f = func(left, pre int) (res int) {
		if left < 0 {
			return
		}
		dv := &dp[left][pre]
		if *dv != math.MinInt {
			return *dv
		}
		defer func() { *dv = res }()

		choose := tot - bits.OnesCount(uint(left))
		for _s := uint(left); _s > 0; _s &= _s - 1 {
			p := bits.TrailingZeros(_s)
			r := f(left^1<<p, p)
			if !ban[choose] {
				r += (p + base) * (pre + base)
			}
			res = max(res, r)
		}

		return
	}
	ans = f(1<<tot-1, 0)

	return
}

func main() {
	fmt.Println(chainGreedy([]int{3, 4}, 11))
	fmt.Println(runAC([]int{3, 4}, 11))
}
