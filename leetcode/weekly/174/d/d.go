package main

func maxJumps(arr []int, d int) (ans int) {
	n := len(arr)

	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if memo[i] > 0 { // 之前计算过
			return memo[i]
		}
		defer func() { memo[i] = res }()

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

	for i := range arr {
		ans = max(ans, dfs(i))
	}
	return
}
