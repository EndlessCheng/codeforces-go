package main

// github.com/EndlessCheng/codeforces-go
func waysToBuyPensPencils(total, cost1, cost2 int) int64 {
	n := 1 + total/cost1
	ans := int64(n)
	for i := 0; i < n; i++ {
		ans += int64((total - cost1*i) / cost2)
	}
	return ans
}
