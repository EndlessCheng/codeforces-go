package main

// https://space.bilibili.com/206214
// 指针从数字 x 旋转到数字 y 的最少旋转次数
func rotate(x, y byte) int {
	d := abs(int(x) - int(y))
	return min(d, 10-d)
}

func minRotations(n int, s string) int {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
