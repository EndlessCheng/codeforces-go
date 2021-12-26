package main

// O(n) 做法：哈希表 + 枚举相同元素

// github.com/EndlessCheng/codeforces-go
func getDistances(arr []int) []int64 {
	pos := map[int][]int{}
	for i, v := range arr {
		pos[v] = append(pos[v], i) // 记录相同元素的位置
	}
	ans := make([]int64, len(arr))
	for _, p := range pos {
		sum := 0
		for _, i := range p {
			sum += i - p[0]
		}
		ans[p[0]] = int64(sum) // 最左侧元素的间隔和
		for i, n := 1, len(p); i < n; i++ { // 计算下一个相同元素的间隔和
			sum -= (n - i*2) * (p[i] - p[i-1]) // 到右边的 n-i 个点的距离更近了，同时到左边 i 个点的距离更远了
			ans[p[i]] = int64(sum)
		}
	}
	return ans
}
