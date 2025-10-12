package main

// https://space.bilibili.com/206214
func longestBalanced(s string) (ans int) {
	n := len(s)

	// 一种字母
	for i := 0; i < n; {
		start := i
		for i++; i < n && s[i] == s[i-1]; i++ {
		}
		ans = max(ans, i-start)
	}

	// 两种字母
	f := func(x, y byte) {
		for i := 0; i < n; i++ {
			pos := map[int]int{0: i - 1} // 前缀和数组的首项是 0，位置相当于在 i-1
			d := 0 // x 的个数减去 y 的个数
			for ; i < n && (s[i] == x || s[i] == y); i++ {
				if s[i] == x {
					d++
				} else {
					d--
				}
				if j, ok := pos[d]; ok {
					ans = max(ans, i-j)
				} else {
					pos[d] = i
				}
			}
		}
	}
	f('a', 'b')
	f('a', 'c')
	f('b', 'c')

	// 三种字母
	type pair struct{ diffAB, diffBC int }
	pos := map[pair]int{{}: -1} // 前缀和数组的首项是 0，位置相当于在 -1
	cnt := [3]int{}
	for i, b := range s {
		cnt[b-'a']++
		p := pair{cnt[0] - cnt[1], cnt[1] - cnt[2]}
		if j, ok := pos[p]; ok {
			ans = max(ans, i-j)
		} else {
			pos[p] = i
		}
	}
	return
}
