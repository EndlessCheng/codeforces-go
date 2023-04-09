package main

// https://space.bilibili.com/206214
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func diagonalPrime(nums [][]int) (ans int) {
	for i, row := range nums {
		if x := row[i]; x > ans && isPrime(x) {
			ans = x
		}
		if x := row[len(nums)-1-i]; x > ans && isPrime(x) {
			ans = x
		}
	}
	return
}
