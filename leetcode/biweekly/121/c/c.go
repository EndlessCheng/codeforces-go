package main

// https://space.bilibili.com/206214
func minimumOperationsToMakeEqual(x, y int) int {
	memo := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if x <= y {
			return y - x
		}
		if v, ok := memo[x]; ok {
			return v
		}
		res := min(x-y,
			dfs(x/11)+x%11+1,
			dfs(x/11+1)+11-x%11+1,
			dfs(x/5)+x%5+1,
			dfs(x/5+1)+5-x%5+1)
		memo[x] = res
		return res
	}
	return dfs(x)
}
