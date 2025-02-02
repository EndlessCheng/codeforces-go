package main

// https://space.bilibili.com/206214
func maxDistance(s string, k int) int {
	ans, x, y := 0, 0, 0
	for i, c := range s {
		switch c {
		case 'N': y++
		case 'S': y--
		case 'E': x++
		default:  x--
		}
		ans = max(ans, min(abs(x)+abs(y)+k*2, i+1))
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }

func maxDistance2(s string, k int) (ans int) {
	cnt := ['Y']int{}
	for _, ch := range s {
		cnt[ch]++
		left := k
		f := func(a, b int) int {
			d := min(a, b, left)
			left -= d
			return abs(b-a) + d*2
		}
		ans = max(ans, f(cnt['N'], cnt['S'])+f(cnt['E'], cnt['W']))
	}
	return
}
