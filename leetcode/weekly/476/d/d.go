package main

import "sort"

// https://space.bilibili.com/206214
func countStableSubarrays1(nums []int, queries [][]int) []int64 {
	n := len(nums)
	// 找递增段
	left := []int{}   // 递增段的左端点
	sum := []int64{0} // 递增子数组个数的前缀和
	start := 0
	for i, x := range nums {
		if i == n-1 || x > nums[i+1] {
			// 找到了一个递增段 [start, i]
			left = append(left, start)
			m := int64(i - start + 1)
			// 长为 m 的子数组中有 m*(m+1)/2 个递增子数组
			// 计算 m*(m+1)/2 的前缀和
			sum = append(sum, sum[len(sum)-1]+m*(m+1)/2)
			start = i + 1 // 下一个递增段的左端点
		}
	}

	ans := make([]int64, len(queries))
	for k, q := range queries {
		l, r := q[0], q[1]
		i := sort.SearchInts(left, l+1)     // 左端点严格大于 l 的第一个区间
		j := sort.SearchInts(left, r+1) - 1 // 包含 r 的最后一个区间

		// l 和 r 在同一个区间
		if i > j {
			m := int64(r - l + 1)
			ans[k] = m * (m + 1) / 2
			continue
		}

		// l 和 r 在不同区间
		// 分成三段 [l, left[i]) + [left[i], left[j]) + [left[j], r]
		// 中间那段的子数组个数用前缀和计算
		m := int64(left[i] - l)
		m2 := int64(r - left[j] + 1)
		ans[k] = m*(m+1)/2 + (sum[j] - sum[i]) + m2*(m2+1)/2
	}
	return ans
}

func countStableSubarrays(nums []int, queries [][]int) []int64 {
	n := len(nums)
	// 计算递增子数组个数的前缀和
	sum := make([]int64, n+1)
	cnt := 0
	for i, x := range nums {
		if i > 0 && x < nums[i-1] {
			cnt = 0
		}
		cnt++
		// 现在 cnt 表示以 i 为右端点的递增子数组个数
		sum[i+1] = sum[i] + int64(cnt)
	}

	// nxt[i] 表示 i 右边下一个递增段的左端点，若不存在则为 n
	nxt := make([]int, n)
	nxt[n-1] = n
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			nxt[i] = nxt[i+1]
		} else {
			nxt[i] = i + 1
		}
	}

	ans := make([]int64, len(queries))
	for k, q := range queries {
		l, r := q[0], q[1]
		l2 := nxt[l]
		if l2 > r { // l 和 r 在同一个区间
			m := int64(r - l + 1)
			ans[k] = m * (m + 1) / 2
		} else { // l 和 r 在不同区间
			// 分成 [l, l2) + [l2, r]
			// 由于 [l2, r] 中的每个右端点对应的左端点都在 [l2, r] 内，所以可以用前缀和计算
			m := int64(l2 - l)
			ans[k] = m*(m+1)/2 + sum[r+1] - sum[l2]
		}
	}
	return ans
}
