package main

// https://space.bilibili.com/206214
func numberOfSubstrings(s string, k int) (ans int) {
	cnt := [26]int{}
	left := 0
	for _, c := range s {
		cnt[c-'a']++
		for cnt[c-'a'] >= k {
			cnt[s[left]-'a']--
			left++
		}
		ans += left
	}
	return
}
