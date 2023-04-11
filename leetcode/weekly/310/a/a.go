package main

// https://space.bilibili.com/206214
func mostFrequentEven(nums []int) int {
	ans := -1
	cnt := map[int]int{}
	for _, x := range nums {
		if x%2 == 0 { // 统计偶数
			cnt[x]++
			if ans < 0 || cnt[x] > cnt[ans] || cnt[x] == cnt[ans] && x < ans {
				ans = x // 出现次数最大的数中，值最小的
			}
		}
	}
	return ans
}
