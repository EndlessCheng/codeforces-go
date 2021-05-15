package main

// github.com/EndlessCheng/codeforces-go
func sumOfFlooredPairs(a []int) (ans int) {
	cnt := make([]int, 2e5)
	for _, v := range a {
		cnt[v]++
	}
	sum := make([]int, 2e5+1)
	for i, v := range cnt {
		sum[i+1] = sum[i] + v
	}
	for i, c := range cnt {
		if c > 0 {
			for d := 1; d*i <= 1e5; d++ {
				ans += c * d * (sum[(d+1)*i] - sum[d*i])
			}
		}
	}
	return ans % (1e9 + 7)
}
