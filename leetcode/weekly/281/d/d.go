package main

// github.com/EndlessCheng/codeforces-go
func coutPairs(nums []int, k int) (ans int64) {
	divisors := []int{}
	for d := 1; d*d <= k; d++ {
		if k%d == 0 {
			divisors = append(divisors, d)
			if d*d < k {
				divisors = append(divisors, k/d)
			}
		}
	}
	cnt := map[int]int{}
	for _, v := range nums {
		ans += int64(cnt[k/gcd(v, k)])
		for _, d := range divisors {
			if v%d == 0 {
				cnt[d]++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
