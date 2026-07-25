package main

import "math"

func stoneGameIII1(stoneValue []int) string {
	n := len(stoneValue)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = math.MinInt // MinInt 表示该状态没有计算过
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}

		p := &memo[i]
		if *p != math.MinInt { // 之前计算过
			// 之前计算过
			return *p
		}

		res := math.MinInt
		sum := 0
		for j := i; j < min(i+3, n); j++ {
			sum += stoneValue[j]
			res = max(res, sum-dfs(j+1))
		}
		*p = res // 记忆化
		return res
	}

	diff := dfs(0)
	if diff == 0 {
		return "Tie"
	}
	if diff > 0 {
		return "Alice"
	}
	return "Bob"
}

func stoneGameIII2(stoneValue []int) string {
	n := len(stoneValue)
	f := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		f[i] = math.MinInt
		sum := 0
		for j := i; j < min(i+3, n); j++ {
			sum += stoneValue[j]
			f[i] = max(f[i], sum-f[j+1])
		}
	}

	diff := f[0]
	if diff == 0 {
		return "Tie"
	}
	if diff > 0 {
		return "Alice"
	}
	return "Bob"
}

func stoneGameIII3(stoneValue []int) string {
	n := len(stoneValue)
	f := make([]int, n+3)
	sufSum := 0 // stoneValue 的后缀和
	for i := n - 1; i >= 0; i-- {
		sufSum += stoneValue[i]
		f[i] = sufSum - min(f[i+1], f[i+2], f[i+3])
	}

	diff := f[0] - (sufSum - f[0])
	if diff == 0 {
		return "Tie"
	}
	if diff > 0 {
		return "Alice"
	}
	return "Bob"
}

func stoneGameIII4(stoneValue []int) string {
	var sufSum, f1, f2, f3 int
	for i := len(stoneValue) - 1; i >= 0; i-- {
		sufSum += stoneValue[i]
		f1, f2, f3 = sufSum-min(f1, f2, f3), f1, f2
	}

	diff := f1 - (sufSum - f1)
	if diff == 0 {
		return "Tie"
	}
	if diff > 0 {
		return "Alice"
	}
	return "Bob"
}

func stoneGameIII(stoneValue []int) string {
	const k = 3
	n := len(stoneValue)
	f := make([]int, n+1)
	q := []int{n}
	sufSum := 0 // stoneValue 的后缀和

	for i := n - 1; i >= 0; i-- {
		sufSum += stoneValue[i]

		// 窗口向左滑动
		// 1. 右边出，保证窗口元素下标在 [i+1, i+k] 中
		if q[0] > i+k { // 离开窗口
			q = q[1:]
		}

		// 2. 计算转移
		f[i] = sufSum - f[q[0]]

		// 3. 左边入，先去掉无用数据，再把 i 入队
		for len(q) > 0 && f[q[len(q)-1]] >= f[i] { // f 更小（或者相等）
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	diff := f[0]*2 - sufSum
	if diff == 0 {
		return "Tie"
	}
	if diff > 0 {
		return "Alice"
	}
	return "Bob"
}
