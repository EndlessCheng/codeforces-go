package main

// https://space.bilibili.com/206214
func validSubstringCount(s, t string) (ans int64) {
	if len(s) < len(t) {
		return 0
	}
	cnt := [26]int{} // t 的字母出现次数与 s 的字母出现次数之差
	for _, b := range t {
		cnt[b-'a']++
	}
	less := 0 // 统计窗口内有多少个字母的出现次数比 t 的少
	for _, c := range cnt {
		if c > 0 {
			less++
		}
	}

	left := 0
	for _, b := range s {
		cnt[b-'a']--
		if cnt[b-'a'] == 0 {
			// 窗口内 b 的出现次数和 t 一样
			less--
		}
		for less == 0 {
			if cnt[s[left]-'a'] == 0 {
				// s[left] 移出窗口后，窗口内 s[left] 的出现次数比 t 的少
				less++
			}
			cnt[s[left]-'a']++
			left++
		}
		ans += int64(left)
	}
	return
}
