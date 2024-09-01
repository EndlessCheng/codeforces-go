package main

// https://space.bilibili.com/206214
func maximumSubarrayXor(nums []int, queries [][]int) []int {
	n := len(nums)
	f := make([][]int, n)
	mx := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		mx[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = nums[i]
		mx[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i][j-1] ^ f[i+1][j]
			mx[i][j] = max(f[i][j], mx[i+1][j], mx[i][j-1])
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = mx[q[0]][q[1]]
	}
	return ans
}
