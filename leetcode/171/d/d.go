package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func do(s []byte, _p2 byte) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	dis := func(p1, p2 int) int {
		x1, y1 := p1/6, p1%6
		x2, y2 := p2/6, p2%6
		v := abs(x1-x2) + abs(y1-y2)
		return v
	}

	const mx = 305
	dp := [mx][26][26]int{}
	vis := [mx][26][26]bool{}
	n := len(s)
	var f func(i int, p1, p2 byte) int
	f = func(i int, p1, p2 byte) (ans int) {
		if i >= n {
			return 0
		}
		if vis[i][p1][p2] {
			return dp[i][p1][p2]
		}
		vis[i][p1][p2] = true
		defer func() { dp[i][p1][p2] = ans }()
		c := s[i]
		if c == p1 || c == p2 {
			return f(i+1, p1, p2)
		}
		f1 := f(i+1, c, p2) + dis(int(p1), int(c))
		f2 := f(i+1, p1, c) + dis(int(p2), int(c))
		return min(f1, f2)
	}
	return f(0, s[0], _p2)
}

func minimumDistance(ss string) int {
	s := []byte(ss)
	for i := range s {
		s[i] -= 'A'
	}
	ans := int(1e9)
	for i := byte(0); i < 26; i++ {
		ans = min(ans, do(s, i))
	}
	return ans
}
