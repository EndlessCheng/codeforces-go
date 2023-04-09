package main

// https://space.bilibili.com/206214
func distance(nums []int) []int64 {
	groups := map[int][]int{}
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := int64(0)
		for _, x := range a {
			s += int64(x - a[0]) // a[0] 到其它下标的距离之和
		}
		ans[a[0]] = s
		for i := 1; i < n; i++ {
			// 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
			s += int64(i*2-n) * int64(a[i]-a[i-1])
			ans[a[i]] = s
		}
	}
	return ans
}
