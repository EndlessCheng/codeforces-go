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
		mx := n - a[len(a)-1] + a[0]
		for i := 1; i < len(a); i++ {
			mx = max(mx, a[i]-a[i-1])
		}
		ans = min(ans, mx)
	}
	return ans / 2 // 最后再除 2
}
