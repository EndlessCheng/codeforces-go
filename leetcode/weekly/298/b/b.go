package main

// https://space.bilibili.com/206214/dynamic
func minimumNumbers(num, k int) int {
	if num == 0 {
		return 0
	}
	for n := 1; n <= 10 && n*k <= num; n++ {
		if (num-n*k)%10 == 0 {
			return n
		}
	}
	return -1
}
