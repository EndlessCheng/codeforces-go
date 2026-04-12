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
		for notPrime[x] != i%2 {
			ans++
			x++
		}
	}
	return
}
