package main

// https://space.bilibili.com/206214
func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	for i, x := range skill {
		s[i+1] = s[i] + x
	}

	suf := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		if skill[i] > skill[suf[len(suf)-1]] {
			suf = append(suf, i)
		}
	}

	pre := []int{0}
	for i := 1; i < n; i++ {
		if skill[i] > skill[pre[len(pre)-1]] {
			pre = append(pre, i)
		}
	}

	start := 0
	for j := 1; j < m; j++ {
		record := suf
		if mana[j-1] < mana[j] {
			record = pre
		}
		mx := 0
		for _, i := range record {
			mx = max(mx, mana[j-1]*s[i+1]-mana[j]*s[i])
		}
		start += mx
	}
	return int64(start + mana[m-1]*s[n])
}
