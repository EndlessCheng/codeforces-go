package main

// github.com/EndlessCheng/codeforces-go
func findKDistantIndices(nums []int, key, k int) (ans []int) {
	vis := make([]bool, len(nums))
	for i, num := range nums {
		if num == key {
			for j := i; j >= i-k && j >= 0 && !vis[j]; j-- {
				vis[j] = true
			}
			for j := min(i+k, len(nums)-1); !vis[j]; j-- {
				vis[j] = true
			}
		}
	}
	for i, b := range vis {
		if b {
			ans = append(ans, i)
		}
	}
	return
}

func min(a, b int) int { if a > b { return b }; return a }
