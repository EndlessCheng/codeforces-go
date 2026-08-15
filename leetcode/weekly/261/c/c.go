package main

// github.com/EndlessCheng/codeforces-go
func stoneGameIX(stones []int) bool {
	cnt := [3]int{}
	for _, x := range stones {
		cnt[x%3]++
	}

	if cnt[0]%2 == 0 {
		return cnt[1] > 0 && cnt[2] > 0
	}
	return abs(cnt[1]-cnt[2]) > 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
