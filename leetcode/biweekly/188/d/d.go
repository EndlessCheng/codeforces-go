package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minMaxWaitingTime11(demand []int, fuel []int) int {
	type pair struct{ maxNum, bestWaitTime int }
	type args struct{ i, wait0, wait1, fuel0, fuel1 int }
	memo := map[args]pair{}

	// 加油机 0 在 wait0 秒后空闲，剩余燃料量 fuel0
	// 加油机 1 在 wait1 秒后空闲，剩余燃料量 fuel1
	var dfs func(int, int, int, int, int) pair
	dfs = func(i, wait0, wait1, fuel0, fuel1 int) pair {
		if i == len(demand) {
			return pair{}
		}

		key := args{i, wait0, wait1, fuel0, fuel1}
		if p, ok := memo[key]; ok {
			return p
		}

		maxNum, bestWaitTime := 0, 0
		d := demand[i]

		// 选择加油机 0，等 wait0 秒开始加油，加油机 1 的等待时间减少 wait0 秒
		if d <= fuel0 {
			p := dfs(i+1, d, max(wait1-wait0, 0), fuel0-d, fuel1)
			maxNum = p.maxNum + 1
			bestWaitTime = max(p.bestWaitTime, wait0)
		}

		// 选择加油机 1，等 wait1 秒开始加油，加油机 0 的等待时间减少 wait1 秒
		if d <= fuel1 {
			p := dfs(i+1, max(wait0-wait1, 0), d, fuel0, fuel1-d)
			num := p.maxNum + 1
			time := max(p.bestWaitTime, wait1)
			if num > maxNum || num == maxNum && time < bestWaitTime {
				maxNum, bestWaitTime = num, time
			}
		}

		res := pair{maxNum, bestWaitTime}
		memo[key] = res
		return res
	}

	ans := dfs(0, 0, 0, fuel[0], fuel[1])
	if ans.maxNum == 0 {
		return -1
	}
	return ans.bestWaitTime
}

func minMaxWaitingTime(demand []int, fuel []int) int {
	type pair struct{ maxNum, bestWaitTime int }
	type args struct{ i, wait1, fuel0, fuel1 int }
	memo := map[args]pair{}

	var dfs func(int, int, int, int) pair
	dfs = func(i, wait1, fuel0, fuel1 int) pair {
		if i == len(demand) {
			return pair{}
		}

		key := args{i, wait1, fuel0, fuel1}
		if p, ok := memo[key]; ok {
			return p
		}

		maxNum, bestWaitTime := 0, 0
		wait0 := 0
		if i > 0 {
			wait0 = demand[i-1]
		}
		d := demand[i]

		// 跟在车 i-1 后面加油
		if d <= fuel0 {
			p := dfs(i+1, max(wait1-wait0, 0), fuel0-d, fuel1)
			maxNum = p.maxNum + 1
			bestWaitTime = max(p.bestWaitTime, wait0)
		}

		// 不跟在车 i-1 后面加油
		if d <= fuel1 {
			p := dfs(i+1, max(wait0-wait1, 0), fuel1-d, fuel0) // 注意这里交换了 fuel0 和 fuel1
			num := p.maxNum + 1
			time := max(p.bestWaitTime, wait1)
			if num > maxNum || num == maxNum && time < bestWaitTime {
				maxNum, bestWaitTime = num, time
			}
		}

		res := pair{maxNum, bestWaitTime}
		memo[key] = res
		return res
	}

	ans := dfs(0, 0, fuel[0], fuel[1])
	if ans.maxNum == 0 {
		return -1
	}
	return ans.bestWaitTime
}

func minMaxWaitingTime21(demand []int, fuel []int) int {
	// 1. 求出最多能服务的车辆数 maxCars
	type fuelArgs struct{ i, fuel0, fuel1 int }
	memo := map[fuelArgs]int{}

	var calcMaxCars func(int, int, int) int
	calcMaxCars = func(i, fuel0, fuel1 int) (res int) {
		if i == len(demand) {
			return 0
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
		return
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

func minMaxWaitingTime22(demand []int, fuel []int) int {
	// 1. 求出最多能服务的车辆数 maxCars
	type fuelArgs struct{ i, fuel0, fuel1 int }
	memo := map[fuelArgs]int{}

	var calcMaxCars func(int, int, int) int
	calcMaxCars = func(i, fuel0, fuel1 int) (res int) {
		if i == len(demand) {
			return 0
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
		return
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
