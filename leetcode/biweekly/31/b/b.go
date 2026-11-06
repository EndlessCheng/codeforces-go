package main

func numOfSubarrays1(arr []int) (ans int) {
	sum := make([]int, len(arr)+1)
	for i, x := range arr {
		sum[i+1] = (sum[i] + x) % 2
	}

	// cnt[0] 统计偶数 s[i] 的个数
	// cnt[1] 统计奇数 s[i] 的个数
	cnt := [2]int{}
	for _, s := range sum {
		ans += cnt[s^1] // 在 s 左边的奇偶性不同于 s 的个数
		cnt[s]++
	}
	return ans % 1_000_000_007
}

func numOfSubarrays(arr []int) (ans int) {
	cnt := [2]int{}
	s := 0
	for _, x := range arr {
		cnt[s]++
		s = (s + x) % 2
		ans += cnt[s^1]
	}
	return ans % 1_000_000_007
}
