package main

// github.com/EndlessCheng/codeforces-go
func find(a, b []int) (res []int) {
	set := map[int]bool{}
	for _, v := range b {
		set[v] = true
	}
	inRes := map[int]bool{}
	for _, v := range a {
		if !inRes[v] && !set[v] {
			inRes[v] = true
			res = append(res, v)
		}
	}
	return
}

func findDifference(nums1, nums2 []int) [][]int {
	return [][]int{find(nums1, nums2), find(nums2, nums1)}
}
