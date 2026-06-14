package main

import (
	"maps"
	"slices"
)

// https://space.bilibili.com/206214
func getLength1(nums []int) (ans int) {
	for i := range nums {
		cnt := map[int]int{} // 出现次数
		cc := map[int]int{}  // 出现次数的出现次数
		for j := i; j < len(nums); j++ {
			x := nums[j]
			if c := cnt[x]; c > 0 {
				cc[c]--
				if cc[c] == 0 {
					delete(cc, c) // 保证我们可以正确计算 cc 的大小
				}
			}
			cnt[x]++
			cc[cnt[x]]++

			if len(cnt) == 1 { // 子数组只有一种元素
				ans = max(ans, j-i+1)
			} else if len(cc) == 2 { // 子数组的元素出现次数恰好有两种
				c := slices.Sorted(maps.Keys(cc))
				if c[0]*2 == c[1] {
					ans = max(ans, j-i+1)
				}
			}
		}
	}
	return
}

func getLength(nums []int) (ans int) {
	n := len(nums)
	set := map[int]struct{}{}
	// 子数组只有一种元素
	same := 0
	for i, x := range nums {
		set[x] = struct{}{}
		same++
		if i == n-1 || x != nums[i+1] { // 到达连续相同段的末尾
			ans = max(ans, same)
			same = 0
		}
	}

	// 所有元素互不相同
	if len(set) == n {
		return 1
	}

	// 最优性剪枝：答案不会变大
	for i := 0; i < n-ans; i++ {
		cnt := map[int]int{} // 出现次数
		cc := map[int]int{}  // 出现次数的出现次数
		for j := i; j < n; j++ {
			x := nums[j]
			if c := cnt[x]; c > 0 {
				cc[c]--
				if cc[c] == 0 {
					delete(cc, c) // 保证我们可以正确计算 cc 的大小
				}
			}
			cnt[x]++
			cc[cnt[x]]++

			if len(cc) == 2 { // 子数组的元素出现次数恰好有两种
				c := slices.Sorted(maps.Keys(cc))
				if c[0]*2 == c[1] {
					ans = max(ans, j-i+1)
				}
			}
		}
	}
	return
}
