package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumBeauty(flowers []int, newFlowers int64, target, full, partial int) int64 {
	sort.Ints(flowers)
	n := len(flowers)
	if flowers[0] >= target { // 剪枝，此时所有花园都是完善的
		return int64(n * full)
	}

	leftFlowers := int(newFlowers) - target*n // 剩余可以种植的花
	for i, f := range flowers {
		flowers[i] = min(f, target)
		leftFlowers += flowers[i]
	}

	ans := 0
	for i, x, sumFlowers := 0, 0, 0; i <= n; i++ { // 枚举后缀长度 n-i
		if leftFlowers >= 0 {
			// 计算最长前缀的长度
			for ; x < i && flowers[x]*x-sumFlowers <= leftFlowers; x++ {
				sumFlowers += flowers[x] // 注意 x 只增不减，这部分的时间复杂度为 O(n)
			}
			// 计算总美丽值
			beauty := (n - i) * full
			if x > 0 {
				beauty += min((leftFlowers+sumFlowers)/x, target-1) * partial
			}
			ans = max(ans, beauty)
		}
		if i < n {
			leftFlowers += target - flowers[i]
		}
	}
	return int64(ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
