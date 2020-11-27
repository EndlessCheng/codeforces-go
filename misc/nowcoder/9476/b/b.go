package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7
const mod2 = (mod + 1) / 2

func getSum(a []int, query []int) []int {
	n := len(a)
	sum := make([]int, n+1)
	sum2 := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
		sum2[i+1] = sum2[i] + v*v
	}
	ans := make([]int, len(query)/2)
	for i := 0; i < len(query); i += 2 {
		l, r := query[i], query[i+1]
		// 注意 sum[i] 是 n*max(a[i]) = 1e10 级别的，在相乘前需要取模，否则可能会达到 1e20
		s := (sum[r] - sum[l-1]) % mod
		s = s*s - (sum2[r] - sum2[l-1])
		s = (s%mod + mod) % mod
		s = s * mod2 % mod
		ans[i/2] = s
	}
	return ans
}
