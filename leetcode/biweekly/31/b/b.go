package main

func numOfSubarrays(a []int) (ans int) {
	sum, odd, even := 0, 0, 1
	for _, v := range a {
		sum += v
		if sum&1 > 0 {
			ans += even
			odd++
		} else {
			ans += odd
			even++
		}
	}
	return ans % (1e9 + 7)
}
