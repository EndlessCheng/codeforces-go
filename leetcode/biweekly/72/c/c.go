package main

// 贪心

// github.com/EndlessCheng/codeforces-go
func maximumEvenSplit(n int64) (ans []int64) {
	if n&1 == 0 {
		for i := int64(2); n >= i; i += 2 {
			ans = append(ans, i)
			n -= i
		}
		ans[len(ans)-1] += n
	}
	return
}
