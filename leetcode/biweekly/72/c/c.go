package main

// 贪心

// github.com/EndlessCheng/codeforces-go
func maximumEvenSplit(finalSum int64) (ans []int64) {
	if finalSum%2 == 0 {
		for i := int64(2); i <= finalSum; i += 2 {
			ans = append(ans, i)
			finalSum -= i
		}
		ans[len(ans)-1] += finalSum
	}
	return
}
