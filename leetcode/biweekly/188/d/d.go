package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minMaxWaitingTime1(demand []int, fuel []int) int {
	// 1. 求出最多能服务的车辆数 maxCars
	type fuelArgs struct{ i, fuel0, fuel1 int }
	memo := map[fuelArgs]int{}

	var calcMaxCars func(int, int, int) int
	calcMaxCars = func(i, fuel0, fuel1 int) (res int) {
		if i == len(demand) {
			return
		}

		args := fuelArgs{i, fuel0, fuel1}
		if v, ok := memo[args]; ok {
			return v
		}

		d := demand[i]
		if d <= fuel0 {
			res = calcMaxCars(i+1, fuel0-d, fuel1) + 1
		}
		if d <= fuel1 {
			res = max(res, calcMaxCars(i+1, fuel0, fuel1-d)+1)
		}

		memo[args] = res
		return res
	}

	maxCars := calcMaxCars(0, fuel[0], fuel[1])
	if maxCars == 0 {
		return -1
	}

	// 2. 二分最大等待时间
	ans := sort.Search(slices.Max(demand), func(maxWaitingTime int) bool {
		// 3. 判断在最大等待时间 <= maxWaitingTime 的约束下，能否服务 maxCars 辆车
		type state struct{ i, wait0, wait1, fuel0, fuel1 int }
		vis := map[state]bool{}

		// 加油机 0 在 wait0 秒后空闲，剩余燃料量 fuel0
		// 加油机 1 在 wait1 秒后空闲，剩余燃料量 fuel1
		var dfs func(int, int, int, int, int) bool
		dfs = func(i, wait0, wait1, fuel0, fuel1 int) bool {
			if i == maxCars {
				return true
			}

			st := state{i, wait0, wait1, fuel0, fuel1}
			if vis[st] {
				return false
			}
			vis[st] = true

			d := demand[i]

			// 选择加油机 0，等 wait0 秒开始加油，加油机 1 的等待时间减少 wait0 秒
			if wait0 <= maxWaitingTime && d <= fuel0 &&
				dfs(i+1, d, max(wait1-wait0, 0), fuel0-d, fuel1) {
				return true
			}

			// 选择加油机 1，等 wait1 秒开始加油，加油机 0 的等待时间减少 wait1 秒
			if wait1 <= maxWaitingTime && d <= fuel1 &&
				dfs(i+1, max(wait0-wait1, 0), d, fuel0, fuel1-d) {
				return true
			}

			return false
		}

		return dfs(0, 0, 0, fuel[0], fuel[1])
	})
	return ans
}

func minMaxWaitingTime(demand []int, fuel []int) int {
	// 1. 求出最多能服务的车辆数 maxCars
	type fuelArgs struct{ i, fuel0, fuel1 int }
	memo := map[fuelArgs]int{}

	var calcMaxCars func(int, int, int) int
	calcMaxCars = func(i, fuel0, fuel1 int) (res int) {
		if i == len(demand) {
			return
		}

		args := fuelArgs{i, fuel0, fuel1}
		if v, ok := memo[args]; ok {
			return v
		}

		d := demand[i]
		if d <= fuel0 {
			res = calcMaxCars(i+1, fuel0-d, fuel1) + 1
		}
		if d <= fuel1 {
			res = max(res, calcMaxCars(i+1, fuel0, fuel1-d)+1)
		}

		memo[args] = res
		return res
	}

	maxCars := calcMaxCars(0, fuel[0], fuel[1])
	if maxCars == 0 {
		return -1
	}

	// 2. 二分最大等待时间
	ans := sort.Search(slices.Max(demand), func(maxWaitingTime int) bool {
		// 3. 判断在最大等待时间 <= maxWaitingTime 的约束下，能否服务 maxCars 辆车
		type state struct{ i, wait1, fuel0, fuel1 int }
		vis := map[state]bool{}

		var dfs func(int, int, int, int) bool
		dfs = func(i, wait1, fuel0, fuel1 int) bool {
			if i == maxCars {
				return true
			}

			st := state{i, wait1, fuel0, fuel1}
			if vis[st] {
				return false
			}
			vis[st] = true

			wait0 := 0
			if i > 0 {
				wait0 = demand[i-1]
			}
			d := demand[i]

			// 跟在车 i-1 后面加油
			if wait0 <= maxWaitingTime && d <= fuel0 &&
				dfs(i+1, max(wait1-wait0, 0), fuel0-d, fuel1) {
				return true
			}

			// 不跟在车 i-1 后面加油
			if wait1 <= maxWaitingTime && d <= fuel1 &&
				dfs(i+1, max(wait0-wait1, 0), fuel1-d, fuel0) { // 注意这里交换了 fuel0 和 fuel1
				return true
			}

			return false
		}

		return dfs(0, 0, fuel[0], fuel[1])
	})
	return ans
}
