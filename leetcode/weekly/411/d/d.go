package main

// https://space.bilibili.com/206214
func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	left2 := make([]int, n)
	l = 0
	for i := range left2 {
		for l < n && left[l] < i {
			l++
		}
		left2[i] = l
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(left2[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
