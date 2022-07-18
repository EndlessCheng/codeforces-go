package main

// https://space.bilibili.com/206214/dynamic
func maximumSum(nums []int) (ans int) {
	mx := map[int][2]int{}
	for _, v := range nums {
		s := 0
		for x := v; x > 0; x /= 10 {
			s += x % 10
		}
		p := mx[s]
		if v > p[0] {
			p[0], p[1] = v, p[0]
		} else if v > p[1] {
			p[1] = v
		}
		mx[s] = p
	}
	ans = -1
	for _, p := range mx {
		if p[1] > 0 {
			ans = max(ans, p[0]+p[1])
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
