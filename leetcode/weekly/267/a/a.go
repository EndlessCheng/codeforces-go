package main

/* O(n) 一次遍历

根据题意，当第 $k$ 个人买完票时，在 $k$ 前面的人买的票不会超过 $\textit{tickets}[k]$，在 $k$ 后面的人买的票不会超过 $\textit{tickets}[k]-1$，累加所有购票数即为答案。

*/

// github.com/EndlessCheng/codeforces-go
func timeRequiredToBuy(tickets []int, k int) (ans int) {
	for i, t := range tickets {
		if i <= k {
			ans += min(t, tickets[k])
		} else {
			ans += min(t, tickets[k]-1)
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
