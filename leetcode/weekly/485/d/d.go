package main

// https://space.bilibili.com/206214
func lexSmallestAfterDeletion(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}

	st := []rune{}
	for _, ch := range s {
		// 如果 ch 比栈顶小，用 ch 代替栈顶，可以让字典序更小
		for len(st) > 0 && ch < st[len(st)-1] && left[st[len(st)-1]-'a'] > 1 {
			left[st[len(st)-1]-'a']--
			st = st[:len(st)-1]
		}
		st = append(st, ch)
	}

	// 最后，移除末尾的重复元素
	for left[st[len(st)-1]-'a'] > 1 {
		left[st[len(st)-1]-'a']--
		st = st[:len(st)-1]
	}

	return string(st)
}
