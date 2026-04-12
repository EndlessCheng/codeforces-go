package main

// https://space.bilibili.com/206214
const mx = 100_004 // 1e5 的下一个质数是 1e5 + 3
var notPrime = [mx]int{1, 1}

func init() {
	for i := 2; i*i < mx; i++ {
		if notPrime[i] == 0 {
			for j := i * i; j < mx; j += i {
				notPrime[j] = 1
			}
		}
	}
}

func minOperations(nums []int) (ans int) {
	for i, x := range nums {
		// 如果 i 是偶数，那么循环直到 is_prime[x] == 1（x 是质数）
		// 如果 i 是奇数，那么循环直到 is_prime[x] == 0（x 不是质数）
		for notPrime[x] != i%2 {
			ans++
			x++
		}
	}
	return
}
