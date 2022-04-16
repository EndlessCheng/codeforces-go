package main

// github.com/EndlessCheng/codeforces-go
func waysToBuyPensPencils(total, cost1, cost2 int) (ans int64) {
	for i := 0; i <= total/cost1; i++ {
		ans += int64((total-cost1*i)/cost2) + 1
	}
	return
}
