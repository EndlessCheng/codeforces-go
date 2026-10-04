package main

// https://space.bilibili.com/206214
func minRotations1(n int, s string) int {
	suf := 0
	for i := 1; i < n; i++ {
		suf += dis(s[i], s[i-1])
	}

	ans := dis('0', s[n-1]) + suf // k=0 的情况

	pre := dis('0', s[0])
	for k := 1; k < n; k++ {
		op := dis(s[k], s[k-1])
		suf -= op // 撤销
		ans = min(ans, pre+dis(s[k-1], s[n-1])+suf)
		pre += op
	}
	return ans
}

// 指针从数字 x 旋转到数字 y 的最少旋转次数
func dis(x, y byte) int {
	d := abs(int(x) - int(y))
	return min(d, 10-d)
}

func minRotations(n int, s string) int {
	base, mn := 0, 0
	pre := byte('0')
	for _, cur := range s {
		op := dis(pre, byte(cur))
		base += op
		mn = min(mn, dis(pre, s[n-1])-op)
		pre = byte(cur)
	}
	return base + mn
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
