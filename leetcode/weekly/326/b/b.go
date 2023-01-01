package main

// https://space.bilibili.com/206214
func distinctPrimeFactors(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				set[i] = struct{}{}
				for x /= i; x%i == 0; x /= i {
				}
			}
		}
		if x > 1 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}
