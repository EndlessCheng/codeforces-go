package main

// github.com/EndlessCheng/codeforces-go
type FindSumPairs struct{}

var (
	x, y []int
	cnt  map[int]int
)

func Constructor(nums1, nums2 []int) (_ FindSumPairs) {
	cnt = map[int]int{}
	for _, v := range nums2 {
		cnt[v]++
	}
	x, y = nums1, nums2
	return
}

func (FindSumPairs) Add(i, val int) {
	cur := y[i]
	cnt[cur]--
	cnt[cur+val]++
	y[i] += val
}

func (FindSumPairs) Count(tot int) (ans int) {
	for _, v := range x {
		ans += cnt[tot-v]
	}
	return
}
