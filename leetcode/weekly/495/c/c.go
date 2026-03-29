package main

// https://space.bilibili.com/206214
func sortableIntegers(nums []int) (ans int) {
	n := len(nums)
	nextDec := make([]int, n) // nums[nextDec[i]] > nums[nextDec[i] + 1]
	nextDec[n-1] = n
	p := n
	// 对于每个 i，记录下一个递减的位置
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			p = i
		}
		nextDec[i] = p
	}

	solve := func(k int) {
		lastMax := 0 // 上一段的最大值
		for r := k - 1; r < n; r += k {
			l := r - k + 1
			m := nextDec[l]
			if m >= r {
				// [l,r] 是递增的，最小值为 nums[l]，最大值为 nums[r]
				// nums[l] 必须 >= 上一段的最大值
				if nums[l] < lastMax {
					return
				}
				lastMax = nums[r]
			} else {
				// [l, m] 是第一段，[m+1, r] 是第二段
				// 第二段必须是递增的，且第二段的最小值 >= 上一段的最大值，且第二段最大值 <= 第一段的最大值
				if nextDec[m+1] < r || nums[m+1] < lastMax || nums[r] > nums[l] {
					return
				}
				lastMax = nums[m]
			}
		}
		ans += k // 满足要求
	}

	// 枚举 n 的因子 k
	for k := 1; k*k <= n; k++ {
		if n%k == 0 {
			solve(k)
			if k*k < n {
				solve(n / k)
			}
		}
	}
	return
}
