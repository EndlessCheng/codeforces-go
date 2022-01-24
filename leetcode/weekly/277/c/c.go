package main

// 哈希表模拟

// github.com/EndlessCheng/codeforces-go
func findLonely(nums []int) (ans []int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for v, c := range cnt {
		if c == 1 && cnt[v+1] == 0 && cnt[v-1] == 0 {
			ans = append(ans, v)
		}
	}
	return
}
