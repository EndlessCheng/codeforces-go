package main

// github.com/EndlessCheng/codeforces-go
type FindSumPairs struct {
	nums1 []int
	nums2 []int
	cnt   map[int]int
}

func Constructor(nums1, nums2 []int) FindSumPairs {
	cnt := map[int]int{}
	for _, x := range nums2 {
		cnt[x]++
	}
	return FindSumPairs{nums1, nums2, cnt}
}

func (p *FindSumPairs) Add(index int, val int) {
	// 维护 nums2 每个元素的出现次数
	p.cnt[p.nums2[index]]--
	p.nums2[index] += val
	p.cnt[p.nums2[index]]++
}

func (p *FindSumPairs) Count(tot int) (ans int) {
	for _, x := range p.nums1 {
		ans += p.cnt[tot-x]
	}
	return
}
