package main

// https://space.bilibili.com/206214
// 438. 找到字符串中所有字母异位词
func findAnagrams(s, p []int) []bool {
	// 统计 p 的每种整数的出现次数
	diff := map[int]int{}
	for _, x := range p {
		diff[x]++
	}

	validL := make([]bool, len(s)-len(p)+1)
	left := 0
	for right, x := range s {
		diff[x]--         // x 进入窗口
		for diff[x] < 0 { // x 太多了
			diff[s[left]]++ // 左端点离开窗口
			left++
		}
		if right-left+1 == len(p) { // 当前窗口与 p 的每种整数的出现次数都相同（证明见 438 题解）
			validL[left] = true
		}
	}
	return validL
}

// 计算并返回 z 数组，其中 z[i] = |LCP(s[i:], s)|
func calcZ(s []int) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func countSuffix(nums []int, target []int) int {
	n, m := len(nums), len(target)
	validL := findAnagrams(target, nums)
	z := calcZ(append(nums, target...))

	lastInc := make([]int, m)
	for i := 1; i < m; i++ {
		if target[i-1] < target[i] {
			lastInc[i] = i
		} else {
			lastInc[i] = lastInc[i-1]
		}
	}

	ans := make([]int, len(validL))
	for k, valid := range validL {
		if valid {
			// target[k,k+n-1] 中的 i 必须满足 max(lastInc[k+n-1], k) <= i <= min(k+z[n+k], k+n-1)
			ans[k] = max(min(k+z[n+k], k+n-1)-max(lastInc[k+n-1], k)+1, 0)
		}
	}
	return ans[0]
}
