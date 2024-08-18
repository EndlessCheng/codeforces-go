package main

// https://space.bilibili.com/206214
func countKConstraintSubstrings(s string, k int) (ans int) {
	cnt := [2]int{}
	left := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[left]&1]--
			left++
		}
		ans += i - left + 1
	}
	return
}
