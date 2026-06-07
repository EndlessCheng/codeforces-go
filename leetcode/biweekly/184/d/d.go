package main

import "slices"

// https://space.bilibili.com/206214
const MX = 100_001

var mu = [MX]int{1: 1}     // 莫比乌斯函数
var divisors = [MX][]int{} // 不含平方因子的因子列表，用于容斥

func init() {
	// 预处理莫比乌斯函数
	// 当 n > 1 时，sum_{d|n} mu[d] = 0
	// 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
	for i := 1; i < MX; i++ {
		for j := i * 2; j < MX; j += i {
			mu[j] -= mu[i] // i 是 j 的真因子
		}
	}

	// 预处理不含平方因子的因子列表
	// 本题不需要因子 1
	for i := 2; i < MX; i++ {
		if mu[i] == 0 {
			continue
		}
		for j := i; j < MX; j += i {
			divisors[j] = append(divisors[j], i) // i 是 j 的因子，且 mu[i] != 0
		}
	}
}

func maxScore(nums []int, maxVal int) (ans int) {
	maxNum := slices.Max(nums)
	cnt := make([]int, maxNum+1)
	for _, x := range nums {
		cnt[x]++
	}
	cntMulti := make([]int, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		for j := i; j <= maxNum; j += i {
			cntMulti[i] += cnt[j] // 统计 nums 中有多少个数是 i 的倍数
		}
	}

	if cnt[1] > 0 {
		ans = 1 // selectedValue = 1 时，无需修改，得分为 1
	}

	// 枚举 selectedValue
	// 优化：如果 selectedValue <= ans，那么 ans 不会变大，跳出循环
	for selectedValue := max(maxNum, maxVal); selectedValue > ans; selectedValue-- {
		if selectedValue > maxVal && cnt[selectedValue] == 0 {
			continue // 无法改成 selectedValue
		}

		// 与 selectedValue 不互质的数，其中一个数改成 selectedValue，其余数都改成 1
		cost := 0
		for _, d := range divisors[selectedValue] {
			if d > maxNum {
				break
			}
			cost -= mu[d] * cntMulti[d]
		}

		if selectedValue <= maxNum && cnt[selectedValue] > 0 {
			cost-- // 如果某个 nums[i] 恰好等于 selectedValue，可以少改一次
		} else if cost == 0 {
			cost = 1 // 至少要有一个数改成 selectedValue
		}

		ans = max(ans, selectedValue-cost)
	}
	return
}
