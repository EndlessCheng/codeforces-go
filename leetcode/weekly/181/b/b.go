package main

const mx = 100_001

var divisorNum, divisorSum [mx]int

func init() {
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisorNum[j]++ // i 是 j 的因子
			divisorSum[j] += i
		}
	}
}

func sumFourDivisors(nums []int) (ans int) {
	for _, x := range nums {
		if divisorNum[x] == 4 {
			ans += divisorSum[x]
		}
	}
	return
}
