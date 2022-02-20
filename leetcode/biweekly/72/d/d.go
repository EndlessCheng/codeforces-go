package main

// github.com/EndlessCheng/codeforces-go
func goodTriplets(nums1, nums2 []int) (ans int64) {
	n := len(nums1)
	p := make([]int, n)
	for i, v := range nums1 {
		p[v] = i
	}
	tree := make([]int, n+1)
	for i := 1; i < n-1; i++ {
		for j := p[nums2[i-1]] + 1; j <= n; j += j & -j {
			tree[j]++
		}
		less, y := 0, p[nums2[i]]
		for j := y; j > 0; j &= j - 1 {
			less += tree[j]
		}
		ans += int64(less) * int64(n-1-y-(i-less))
	}
	return
}
