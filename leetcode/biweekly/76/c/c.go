package main

// github.com/EndlessCheng/codeforces-go
var a = [5]int{20, 50, 100, 200, 500}

type ATM [5]int

func Constructor() ATM { return ATM{} }

func (left *ATM) Deposit(banknotesCount []int) {
	for i, count := range banknotesCount {
		left[i] += count // 存钱
	}
}

func (left *ATM) Withdraw(amount int) []int {
	ans := make([]int, 5)
	for i := 4; i >= 0; i-- {
		ans[i] = min(amount/a[i], left[i])
		amount -= ans[i] * a[i] // 取钱
	}
	if amount > 0 { // 没法取恰好 amount
		return []int{-1}
	}
	for i, count := range ans {
		left[i] -= count
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
