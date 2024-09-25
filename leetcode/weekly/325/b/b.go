package main

// https://space.bilibili.com/206214
func takeCharacters(s string, k int) int {
	cnt := [3]int{}
	for _, c := range s {
		cnt[c-'a']++ // 一开始，把所有字母都取走
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1 // 字母个数不足 k
	}

	mx, left := 0, 0
	for right, c := range s {
		c -= 'a'
		cnt[c]-- // 移入窗口，相当于不取走 c
		for cnt[c] < k { // 窗口之外的 c 不足 k
			cnt[s[left]-'a']++ // 移出窗口，相当于取走 s[left]
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}
