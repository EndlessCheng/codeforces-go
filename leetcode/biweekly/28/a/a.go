package main

// github.com/EndlessCheng/codeforces-go
func finalPrices1(prices []int) []int {
	n := len(prices)
	st := []int{0} // 哨兵，作为没有折扣时的栈顶值
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for st[len(st)-1] > p {
			st = st[:len(st)-1] // p 比栈顶小，比栈顶更能成为左侧元素的折扣值
		}
		// 循环结束后，栈顶的价格 <= p，作为折扣值
		ans[i] = p - st[len(st)-1]
		st = append(st, p)
	}
	return ans
}

func finalPrices(prices []int) []int {
	st := []int{} // todolist
	for i, p := range prices {
		for len(st) > 0 && prices[st[len(st)-1]] >= p {
			prices[st[len(st)-1]] -= p // p 是栈顶的折扣值
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return prices
}
