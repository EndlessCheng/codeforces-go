package main

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k, m int) int64 {
	// 子数组至少有 distinctLimit 个不同的整数，其中至少有 k 个不同的好整数
	calc := func(distinctLimit int) (ans int64) {
		cnt := map[int]int{} // 用来统计不同元素个数
		goodNumbers := 0 // 窗口中的好整数（出现至少 m 次的整数）的个数
		left := 0
		for _, x := range nums {
			// 1. 入
			cnt[x]++
			if cnt[x] == m {
				goodNumbers++
			}

			// 2. 出
			for len(cnt) >= distinctLimit && goodNumbers >= k {
				out := nums[left]
				if cnt[out] == m {
					goodNumbers--
				}
				cnt[out]--
				if cnt[out] == 0 {
					delete(cnt, out)
				}
				left++
			}

			// 3. 更新答案
			ans += int64(left)
		}
		return
	}

	return calc(k) - calc(k+1)
}

func countSubarrays2(nums []int, k, m int) (ans int64) {
	cnt := map[int]int{}
	geM := 0     // 窗口中的出现次数 >= m 的元素个数
	minLeft := 0 // 窗口左端点的最小值
	maxLeft := 0 // 窗口左端点的最大值
	for _, x := range nums {
		// 1. 入
		cnt[x]++
		if cnt[x] == m {
			geM++
		}

		// 2. 出（更新 minLeft）
		for len(cnt) > k {
			out := nums[minLeft]
			if cnt[out] == m {
				geM--
			}
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			minLeft++
		}

		// 2. 出（更新 maxLeft）
		maxLeft = max(maxLeft, minLeft)
		// 窗口恰好有 k 个不同元素，且这些数都至少出现 m 次，且左端点元素的出现次数大于 m，则可以把左端点元素移出窗口
		for len(cnt) == k && geM == k && cnt[nums[maxLeft]] > m {
			cnt[nums[maxLeft]]--
			maxLeft++
		}

		// 3. 更新答案（前提是窗口恰好有 k 个不同元素，且这些数都至少出现 m 次）
		if len(cnt) == k && geM == k {
			// 右端点固定时，在 [minLeft, maxLeft] 中的左端点都符合要求
			ans += int64(maxLeft - minLeft + 1)
		}
	}
	return
}
