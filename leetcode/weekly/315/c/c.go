package main

// https://space.bilibili.com/206214
func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		rev := 0
		for x := i; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		if i+rev == num {
			return true
		}
	}
	return false
}
