package main

// https://space.bilibili.com/206214
func validSubstringCount(s, t string) (ans int64) {
	if len(s) < len(t) {
		return 0
	}

	diff := [26]int{} // t 的字母出现次数与 s 的字母出现次数之差
	for _, c := range t {
		diff[c-'a']++
	}

	// 统计窗口内有多少个字母的出现次数比 t 的少
	less := 0
	for _, d := range diff {
		if d > 0 {
			less++
		}
	}

	left := 0
	for _, c := range s {
		diff[c-'a']--
		if diff[c-'a'] == 0 {
			// c 移入窗口后，窗口内 c 的出现次数和 t 的一样
			less--
		}
		for less == 0 { // 窗口符合要求
			if diff[s[left]-'a'] == 0 {
				// s[left] 移出窗口之前，检查出现次数，
				// 如果窗口内 s[left] 的出现次数和 t 的一样，
				// 那么 s[left] 移出窗口后，窗口内 s[left] 的出现次数比 t 的少
				less++
			}
			diff[s[left]-'a']++
			left++
		}
		ans += int64(left)
	}
	return
}
