package main

// github.com/EndlessCheng/codeforces-go
func countDistinct(nums []int, k, p int) int {
	set := map[[200]int]struct{}{}
	for i := range nums {
		arr, idx, cnt := [200]int{}, 0, 0
		for _, v := range nums[i:] {
			if v%p == 0 {
				if cnt++; cnt > k {
					break
				}
			}
			arr[idx] = v
			idx++
			set[arr] = struct{}{}
		}
	}
	return len(set)
}
