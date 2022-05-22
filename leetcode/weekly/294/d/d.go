package main

// github.com/EndlessCheng/codeforces-go
func totalStrength(strength []int) (ans int) {
	const mod int = 1e9 + 7

	n := len(strength)
	left := make([]int, n) // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
	st := []int{}
	for i, v := range strength {
		for len(st) > 0 && strength[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n) // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		v := strength[i]
		for len(st) > 0 && strength[st[len(st)-1]] > v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	s := make([]int, n+1) // 前缀和
	for i, v := range strength {
		s[i+1] = (s[i] + v) % mod
	}
	ss := make([]int, n+2) // 前缀和的前缀和
	for i, v := range s {
		ss[i+1] = (ss[i] + v) % mod
	}
	for i, v := range strength {
		l, r := left[i]+1, right[i]-1 // [l,r] 左闭右闭
		tot := ((i-l+1)*(ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % mod
		ans = (ans + v*tot) % mod // 累加贡献
	}
	return (ans + mod) % mod // 防止算出负数（因为上面算 tot 有个减法）
}
