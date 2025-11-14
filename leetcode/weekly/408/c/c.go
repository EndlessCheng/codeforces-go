package main

// https://space.bilibili.com/206214
func numberOfSubstrings(s string) (ans int) {
	pos0 := []int{-1} // 哨兵，方便处理 cnt0 达到最大时的计数
	total1 := 0 // [0,r] 中的 1 的个数
	for r, ch := range s {
		if ch == '0' {
			pos0 = append(pos0, r) // 记录 0 的下标
		} else {
			total1++
			ans += r - pos0[len(pos0)-1] // 单独计算不含 0 的子串个数
		}

		m := len(pos0)
		// 倒着遍历 pos0，那么 cnt0 = m-i
		for i := m - 1; i > 0 && (m-i)*(m-i) <= total1; i-- {
			p, q := pos0[i-1], pos0[i]
			cnt0 := m - i
			cnt1 := r - q + 1 - cnt0 // [q,r] 中的 1 的个数 = [q,r] 的长度 - cnt0
			ans += max(q-max(cnt0*cnt0-cnt1, 0)-p, 0)
		}
	}
	return
}
