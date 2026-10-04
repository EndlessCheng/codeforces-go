package main

// https://space.bilibili.com/206214
func minRotations1(n int, s string) int {
	suf := 0
	for i := 1; i < n; i++ {
		suf += rotate(s[i], s[i-1])
	}

	ans := rotate('0', s[n-1]) + suf // k=0 的情况

	pre := rotate('0', s[0])
	for k := 1; k < n; k++ {
		op := rotate(s[k], s[k-1])
		suf -= op // 撤销
		ans = min(ans, pre+rotate(s[k-1], s[n-1])+suf)
		pre += op
	}
	return ans
}

// 指针从数字 x 旋转到数字 y 的最少旋转次数
func rotate(x, y byte) int {
	d := abs(int(x) - int(y))
	return min(d, 10-d)
}

func minRotations(n int, s string) int {
	op0 := rotate('0', s[0])
	base := op0
	mn := min(rotate('0', s[n-1])-op0, 0) // k=0 时的增量（但不能超过 0）
	for k := 1; k < n; k++ {
		op := rotate(s[k-1], s[k])
		base += op
		mn = min(mn, rotate(s[k-1], s[n-1])-op)
	}
	return base + mn
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
