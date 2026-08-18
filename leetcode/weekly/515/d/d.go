package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func elevatorRequests1(n int, start int, requests [][]int) int64 {
	m := len(requests)
	memo := make([][]int, 1<<m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	// 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
	var dfs func(int, int) int
	dfs = func(mask, i int) int {
		mask ^= 1 << i
		req := requests[i]
		t, x := req[0], req[1]
		if mask == 0 {
			// i 是第一个被处理的请求
			return max(abs(x-start), t)
		}

		p := &memo[mask][i]
		if *p != -1 { // 之前计算过
			return *p
		}

		res := math.MaxInt
		for j, r := range requests {
			if mask>>j&1 > 0 {
				// 处理完请求 j 的时间 + 从 j 到 i 的时间
				res = min(res, dfs(mask, j)+abs(x-r[1]))
			}
		}
		// 处理完请求 i 的时间不能早于 t
		res = max(res, t)

		*p = res // 记忆化
		return res
	}

	ans := math.MaxInt
	for i := range m {
		ans = min(ans, dfs(1<<m-1, i))
	}
	return int64(ans)
}

func elevatorRequests2(n int, start int, requests [][]int) int64 {
	m := len(requests)
	f := make([][]int, 1<<m)
	for i := range f {
		f[i] = make([]int, m)
	}

	for i, req := range requests {
		f[1<<i][i] = max(abs(req[1]-start), req[0])
	}

	for mask := 1; mask < 1<<m; mask++ {
		if mask&(mask-1) == 0 {
			continue
		}
		for i, req := range requests {
			if mask>>i&1 == 0 {
				continue
			}
			res := math.MaxInt
			msk := mask ^ 1<<i
			t, x := req[0], req[1]
			for j, r := range requests {
				if msk>>j&1 > 0 {
					res = min(res, f[msk][j]+abs(x-r[1]))
				}
			}
			f[mask][i] = max(res, t)
		}
	}

	return int64(slices.Min(f[1<<m-1]))
}

//

func elevatorRequests3(n int, start int, requests [][]int) int64 {
	requests = append(requests, []int{0, -1}, []int{0, n})                 // 插入两个哨兵
	slices.SortFunc(requests, func(a, b []int) int { return a[1] - b[1] }) // 按楼层排序
	m := len(requests)                                                     // 不含哨兵的下标范围是 [1, m-2]

	memo := make([][][2]int, m)
	for i := range memo {
		memo[i] = make([][2]int, m)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 表示该状态没有计算过
		}
	}

	// dfs(i, j, false) 返回完成请求 [1,i] ∪ [j+1,m-2] 所需的最短时间，此时电梯在 floor[i]（最后一个完成的请求是 i）
	// dfs(i, j, true)  返回完成请求 [1,i-1] ∪ [j,m-2] 所需的最短时间，此时电梯在 floor[j]（最后一个完成的请求是 j）
	var dfs func(int, int, uint8) int
	dfs = func(i, j int, isRight uint8) int {
		if i == 0 || j == m-1 { // 出界
			return math.MaxInt / 2
		}

		p := &memo[i][j][isRight]
		if *p != -1 {
			return *p
		}

		k := i
		if isRight > 0 {
			k = j
		}
		t, x := requests[k][0], requests[k][1]
		if i == 1 && j == m-2 { // 当前请求是第一个请求
			*p = max(abs(x-start), t) // 从 start 到当前楼层
		} else {
			*p = min(max(dfs(i-1, j, 0)+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
				max(dfs(i, j+1, 1)+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
		}
		return *p
	}

	ans := math.MaxInt
	// 枚举最后一个完成的请求
	for i := 1; i < m-1; i++ {
		ans = min(ans, dfs(i, i, 0))
	}
	return int64(ans)
}

func elevatorRequests4(n int, start int, requests [][]int) int64 {
	requests = append(requests, []int{0, -1}, []int{0, n})                 // 插入两个哨兵
	slices.SortFunc(requests, func(a, b []int) int { return a[1] - b[1] }) // 按楼层排序
	m := len(requests)                                                     // 不含哨兵的下标范围是 [1, m-2]

	f := make([][][2]int, m)
	for i := range f {
		f[i] = make([][2]int, m)
	}
	for j := range f[0] {
		f[0][j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i < m-1; i++ {
		f[i][m-1] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
		t, x := requests[i][0], requests[i][1]
		for j := m - 2; j >= i; j-- {
			t2, y := requests[j][0], requests[j][1]
			if i == 1 && j == m-2 { // 当前请求是第一个请求
				// 从 start 到当前楼层
				f[i][j][0] = max(abs(x-start), t)
				f[i][j][1] = max(abs(y-start), t2)
				continue
			}
			f[i][j][0] = min(max(f[i-1][j][0]+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
				max(f[i][j+1][1]+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
			f[i][j][1] = min(max(f[i-1][j][0]+y-requests[i-1][1], t2), // 从 floor[i-1] 到当前楼层
				max(f[i][j+1][1]+requests[j+1][1]-y, t2)) // 从 floor[j+1] 到当前楼层
		}
	}

	// 枚举最后一个完成的请求
	ans := math.MaxInt
	for i := 1; i < m-1; i++ {
		ans = min(ans, f[i][i][0])
	}
	return int64(ans)
}

func elevatorRequests(n int, start int, requests [][]int) int64 {
	requests = append(requests, []int{0, -1}, []int{0, n}) // 插入两个哨兵
	slices.SortFunc(requests, func(a, b []int) int { return a[1] - b[1] }) // 按楼层排序
	m := len(requests) // 不含哨兵的下标范围是 [1, m-2]

	f := make([][2]int, m)
	for j := range f {
		f[j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i < m-1; i++ {
		t, x := requests[i][0], requests[i][1]
		for j := m - 2; j >= i; j-- {
			t2, y := requests[j][0], requests[j][1]
			if i == 1 && j == m-2 { // 当前请求是第一个请求
				// 从 start 到当前楼层
				f[j][0] = max(abs(x-start), t)
				f[j][1] = max(abs(y-start), t2)
				continue
			}
			f[j][1] = min(max(f[j][0]+y-requests[i-1][1], t2), // 从 floor[i-1] 到当前楼层
						max(f[j+1][1]+requests[j+1][1]-y, t2)) // 从 floor[j+1] 到当前楼层
			f[j][0] = min(max(f[j][0]+x-requests[i-1][1], t), // 从 floor[i-1] 到当前楼层
						max(f[j+1][1]+requests[j+1][1]-x, t)) // 从 floor[j+1] 到当前楼层
		}
	}

	// 枚举最后一个完成的请求
	ans := math.MaxInt
	for i := 1; i < m-1; i++ {
		ans = min(ans, f[i][0])
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
