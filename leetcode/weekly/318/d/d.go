package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minimumTotalDistance1(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)

	n, m := len(factory), len(robot)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j < 0 { // 所有机器人都修完了
			return 0
		}
		if i < 0 { // 还有机器人没修，但没有工厂了
			return math.MaxInt / 2 // 避免加法溢出
		}

		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		// 工厂 i 不修机器人
		res = dfs(i-1, j)

		position, limit := factory[i][0], factory[i][1]
		disSum := 0
		// 枚举修 k 个机器人
		for k := 1; k <= min(j+1, limit); k++ {
			disSum += abs(robot[j-k+1] - position)
			res = min(res, dfs(i-1, j-k)+disSum)
		}
		return
	}

	return int64(dfs(n-1, m-1))
}

func minimumTotalDistance2(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)

	n, m := len(factory), len(robot)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = math.MaxInt / 2
	}

	for i, fac := range factory {
		position, limit := fac[0], fac[1]
		for j := 1; j <= m; j++ {
			// 工厂 i 不修机器人
			res := f[i][j]

			// 修理下标在 [k, j-1] 中的机器人
			disSum := 0
			for k := j - 1; k >= max(j-limit, 0); k-- {
				disSum += abs(robot[k] - position)
				res = min(res, f[i][k]+disSum)
			}

			f[i+1][j] = res
		}
	}

	return int64(f[n][m])
}

func minimumTotalDistance3(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)

	n, m := len(factory), len(robot)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = math.MaxInt / 2
	}

	for i, fac := range factory {
		position, limit := fac[0], fac[1]

		disSum := make([]int, m+1) // 前缀和
		for j, r := range robot {
			disSum[j+1] = disSum[j] + abs(r-position)
		}

		type pair struct{ i, v int }
		q := []pair{{0, 0}}

		for j := 1; j <= m; j++ {
			// 1. 入
			v := f[i][j] - disSum[j]
			for len(q) > 0 && q[len(q)-1].v >= v {
				q = q[:len(q)-1]
			}
			q = append(q, pair{j, v})

			// 2. 出
			for q[0].i < j-limit {
				q = q[1:]
			}

			// 3. 队首为滑动窗口最小值
			f[i+1][j] = disSum[j] + q[0].v
		}
	}

	return int64(f[n][m])
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)

	m := len(robot)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = math.MaxInt / 2
	}

	for _, fac := range factory {
		position, limit := fac[0], fac[1]

		disSum := 0
		type pair struct{ i, v int }
		q := []pair{{0, 0}}

		for j, r := range robot {
			j++
			disSum += abs(r - position)

			// 1. 入
			v := f[j] - disSum
			for len(q) > 0 && q[len(q)-1].v >= v {
				q = q[:len(q)-1]
			}
			q = append(q, pair{j, v})

			// 2. 出
			for q[0].i < j-limit {
				q = q[1:]
			}

			// 3. 队首为滑动窗口最小值
			f[j] = disSum + q[0].v
		}
	}

	return int64(f[m])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
