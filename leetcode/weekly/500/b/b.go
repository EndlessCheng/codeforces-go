package main

// https://space.bilibili.com/206214
const mx = 1001
var isPrime [mx]int

func init() {
	for i := 2; i < mx; i++ {
		isPrime[i] = 1
	}
	for i := 2; i*i < mx; i++ {
		if isPrime[i] > 0 {
			for j := i * i; j < mx; j += i {
				isPrime[j] = 0
			}
		}
	}

	// 原地计算 isPrime 的质数前缀和
	for i := 1; i < mx; i++ {
		if isPrime[i] > 0 {
			isPrime[i] = isPrime[i-1] + i
		} else {
			isPrime[i] = isPrime[i-1]
		}
	}
}

func sumOfPrimesInRange(n int) int {
	r := 0
	for x := n; x > 0; x /= 10 {
		r = r*10 + x%10
	}
	return isPrime[max(n, r)] - isPrime[min(n, r)-1]
}
