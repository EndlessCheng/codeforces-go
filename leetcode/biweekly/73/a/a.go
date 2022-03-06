package main

// github.com/EndlessCheng/codeforces-go
func mostFrequent(nums []int, key int) (ans int) {
	cnt := map[int]int{}
	maxCnt := 0
	for i, num := range nums[:len(nums)-1] {
		if num == key {
			v := nums[i+1]
			cnt[v]++
			if cnt[v] > maxCnt {
				maxCnt = cnt[v]
				ans = v
			}
		}
	}
	return
}
