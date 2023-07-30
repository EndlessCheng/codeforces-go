package main

// https://space.bilibili.com/206214
func countCompleteSubarrays(nums []int) (ans int) {
	set := map[int]struct{}{}
	for _, v := range nums {
		set[v] = struct{}{}
	}
	m := len(set)

	cnt := map[int]int{}
	left := 0
	for _, v := range nums { // 枚举子数组右端点 i
		ans += left // 子数组左端点 < left 的都是合法的
		cnt[v]++
		for len(cnt) == m {
			ans++ // 子数组左端点等于 left 是合法的
			x := nums[left]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			left++
		}
	}
	return
}
