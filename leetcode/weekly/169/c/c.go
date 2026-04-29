package main

func canReach(arr []int, start int) bool {
	n := len(arr)
	vis := make([]bool, n)

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i < 0 || i >= n || vis[i] { // 出界，或者之前访问过（没找到）
			return false
		}
		if arr[i] == 0 { // 找到了
			return true
		}
		vis[i] = true // 避免重复访问
		return dfs(i+arr[i]) || dfs(i-arr[i])
	}

	return dfs(start)
}
