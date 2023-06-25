package main

// https://space.bilibili.com/206214
func countBeautifulPairs(nums []int) (ans int) {
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 { // 这里需要 O(log x) 的时间
			x /= 10
		}
		cnt[x]++
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
