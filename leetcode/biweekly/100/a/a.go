package main

// https://space.bilibili.com/206214
func distMoney(money, children int) int {
	money -= children // 每人至少 1 美元
	if money < 0 {
		return -1
	}
	ans := min(money/7, children) // 初步分配，让尽量多的孩子分配到 8 美元（注意前面已经分配了 1 美元）
	money -= ans * 7
	children -= ans
	if children == 0 && money > 0 || // 必须找一个前面分配了 8 美元的孩子，分配完剩余的钱
		children == 1 && money == 3 { // 不能有孩子分配恰好 4 美元
		ans--
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
