package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func elevatorRequests1(n int, start int, requests []int) int64 {
	requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
	slices.Sort(requests)
	m := len(requests)

	memo := make([][][2]int, m-1)
	for i := range memo {
		memo[i] = make([][2]int, m-1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}

	// 已处理完 requests 的子数组 [i, j]
	// isRight = 0 表示电梯在 requests[i]
	// isRight = 1 表示电梯在 requests[j]
	var dfs func(int, int, int) int
	dfs = func(i, j, isRight int) int {
		if i == 0 || j == m-1 { // 出界
			return math.MaxInt / 2
		}
		if i == 1 && j == m-2 { // 已处理完所有请求
			return 0
		}

		p := &memo[i][j][isRight]
		if *p != -1 {
			return *p
		}

		var x int
		if isRight > 0 {
			x = requests[j]
		} else {
			x = requests[i]
		}
		remain := m - 3 - j + i
		*p = min(dfs(i-1, j, 0)+(x-requests[i-1])*remain, // 往左
			dfs(i, j+1, 1)+(requests[j+1]-x)*remain) // 往右
		return *p
	}

	i := sort.SearchInts(requests, start)
	return int64(dfs(i, i, 0)) // 这里 0 和 1 是一样的
}

func elevatorRequests2(n int, start int, requests []int) int64 {
	requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
	slices.Sort(requests)
	m := len(requests)

	idx := sort.SearchInts(requests, start)
	f := make([][][2]int, idx+1)
	for i := range f {
		f[i] = make([][2]int, m)
	}
	for j := range f[0] {
		f[0][j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}

	for i := 1; i <= idx; i++ {
		f[i][m-1] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
		for j := m - 2; j >= idx; j-- {
			if i == 1 && j == m-2 {
				f[i][j] = [2]int{}
				continue
			}
			remain := m - 3 - j + i
			f[i][j][0] = min(f[i-1][j][0]+(requests[i]-requests[i-1])*remain, // 往左
				f[i][j+1][1]+(requests[j+1]-requests[i])*remain) // 往右
			f[i][j][1] = min(f[i-1][j][0]+(requests[j]-requests[i-1])*remain, // 往左
				f[i][j+1][1]+(requests[j+1]-requests[j])*remain) // 往右
		}
	}
	return int64(f[idx][idx][0])
}

func elevatorRequests(n int, start int, requests []int) int64 {
	requests = append(requests, start, -1, n) // 插入 start 和两个哨兵
	slices.Sort(requests)
	m := len(requests)

	idx := sort.SearchInts(requests, start)
	f := make([][2]int, m)
	for i := range f {
		f[i] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
	}
	for i := 1; i <= idx; i++ {
		for j := m - 2; j >= idx; j-- {
			if i == 1 && j == m-2 {
				f[j] = [2]int{}
				continue
			}
			remain := m - 3 - j + i
			f[j][1] = min(f[j][0]+(requests[j]-requests[i-1])*remain, // 往左
						f[j+1][1]+(requests[j+1]-requests[j])*remain) // 往右
			f[j][0] = min(f[j][0]+(requests[i]-requests[i-1])*remain, // 往左
						f[j+1][1]+(requests[j+1]-requests[i])*remain) // 往右
		}
	}
	return int64(f[idx][0])
}
