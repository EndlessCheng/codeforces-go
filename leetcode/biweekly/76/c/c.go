package main

// github.com/EndlessCheng/codeforces-go
var denominations = [...]int{20, 50, 100, 200, 500}

const kinds = len(denominations)

type ATM [kinds]int

func Constructor() ATM {
	return ATM{}
}

func (banknotes *ATM) Deposit(banknotesCount []int) {
	// 存钱
	for i, count := range banknotesCount {
		banknotes[i] += count
	}
}

func (banknotes *ATM) Withdraw(amount int) []int {
	ans := make([]int, kinds)

	// 计算每种钞票所需数量
	for i := kinds - 1; i >= 0; i-- {
		ans[i] = min(amount/denominations[i], banknotes[i])
		amount -= ans[i] * denominations[i]
	}

	// 无法取恰好 amount
	if amount > 0 {
		return []int{-1}
	}

	// 取钱
	for i, count := range ans {
		banknotes[i] -= count
	}

	return ans
}
