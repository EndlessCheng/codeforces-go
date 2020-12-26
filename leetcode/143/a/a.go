package main

// github.com/EndlessCheng/codeforces-go
func distributeCandies(candies, n int) (ans []int) {
	ans = make([]int, n)
	for i := 0; ; i++ {
		if candies < i+1 {
			ans[i%n] += candies
			break
		}
		ans[i%n] += i + 1
		candies -= i + 1
	}
	return
}
