package main

// https://space.bilibili.com/206214
func maxSum(nums []int, m, k int) (ans int64) {
	cnt := map[int]int{}
	sum := int64(0)
	for _, x := range nums[:k-1] { // 统计 k-1 个数
		cnt[x]++
		sum += int64(x)
	}
	for i, in := range nums[k-1:] {
		cnt[in]++ // 再添加一个数就是 k 个数了
		sum += int64(in)
		if len(cnt) >= m && sum > ans {
			ans = sum
		}

		out := nums[i-k+1]
		sum -= int64(out)
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return
}
