package main

// https://space.bilibili.com/206214
func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	if a == e && (c != e || (d-b)*(d-f) > 0) || // 车直接攻击到皇后（同一行）
		b == f && (d != f || (c-a)*(c-e) > 0) || // 车直接攻击到皇后（同一列）
		c+d == e+f && (a+b != e+f || (a-c)*(a-e) > 0) || // 象直接攻击到皇后
		c-d == e-f && (a-b != e-f || (a-c)*(a-e) > 0) {
		return 1
	}
	return 2
}

// m 在 l 和 r 之间（写不写等号都可以）
func inBetween(l, m, r int) bool {
	return min(l, r) < m && m < max(l, r)
}

func minMovesToCaptureTheQueen2(a, b, c, d, e, f int) int {
	if a == e && (c != e || !inBetween(b, d, f)) || // 车直接攻击到皇后（同一行）
		b == f && (d != f || !inBetween(a, c, e)) || // 车直接攻击到皇后（同一列）
		c+d == e+f && (a+b != e+f || !inBetween(c, a, e)) || // 象直接攻击到皇后
		c-d == e-f && (a-b != e-f || !inBetween(c, a, e)) {
		return 1
	}
	return 2
}
