package main

import "slices"

func maxJumps1(arr []int, d int) (ans int) {
	n := len(arr)
	memo := make([]int, n)

	var dfs func(int) int
	dfs = func(i int) (res int) {
		p := &memo[i]
		if *p > 0 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化

		// 往左跳
		for j := i - 1; j >= max(i-d, 0) && arr[j] < arr[i]; j-- {
			res = max(res, dfs(j))
		}

		// 往右跳
		for j := i + 1; j <= min(i+d, n-1) && arr[j] < arr[i]; j++ {
			res = max(res, dfs(j))
		}

		return res + 1 // +1 提到循环外面
	}

	// 枚举起点
	for i := range n {
		ans = max(ans, dfs(i))
	}
	return
}

func maxJumps(arr []int, d int) (ans int) {
	n := len(arr)
	// 计算 arr[i] 左边最近的更大元素 arr[left[i]]
	left := make([]int, n)
	st := []int{}
	for i, x := range arr {
		for len(st) > 0 && arr[st[len(st)-1]] <= x {
			st = st[:len(st)-1]
		}
		if len(st) > 0 && i-st[len(st)-1] <= d {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1 // 左边没有更大的数，或者跳跃距离超过 d，都标记为 -1
		}
		st = append(st, i)
	}

	// 计算 arr[i] 右边最近的更大元素 arr[right[i]]
	right := make([]int, n)
	st = st[:0]
	for i, x := range slices.Backward(arr) {
		for len(st) > 0 && arr[st[len(st)-1]] <= x {
			st = st[:len(st)-1]
		}
		if len(st) > 0 && st[len(st)-1]-i <= d {
			right[i] = st[len(st)-1]
		} else {
			right[i] = -1 // 右边没有更大的数，或者跳跃距离超过 d，都标记为 -1
		}
		st = append(st, i)
	}

	memo := make([]int, n)

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 { // 没有更大的数，或者跳跃距离超过 d
			return 0
		}
		p := &memo[i]
		if *p == 0 { // 没有计算过
			// 往左跳 vs 往右跳
			*p = max(dfs(left[i]), dfs(right[i])) + 1
		}
		return *p
	}

	// 枚举终点，倒着跳
	for i := range n {
		ans = max(ans, dfs(i))
	}
	return
}
