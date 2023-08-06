package main

// https://space.bilibili.com/206214
func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		a = append(a, a[0]+n)
		mx := 0
		for i := 1; i < len(a); i++ {
			mx = max(mx, (a[i]-a[i-1])/2)
		}
		ans = min(ans, mx)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
