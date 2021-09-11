package main

import "sort"

/*
排序+贪心

将 $\textit{cards}$ 从大到小排序，并累加前 $\textit{cnt}$ 个元素之和，记作 $\textit{sum}$，若 $\textit{sum}$ 是偶数则直接返回，若不是偶数，则我们需要从前 $\textit{cnt}$ 个元素中选一个元素 $x$，并从后面找一个最大的且奇偶性和 $x$ 不同的元素替换 x，这样就可以使 $\textit{sum}$ 为偶数。

为了使 $\textit{sum}$ 尽可能大，我们需要使替换前后的变化尽可能地小，我们分两种情况讨论：

- 替换 $\textit{card}[\textit{cnt}-1]$；
- 替换前 $\textit{cnt}$ 个元素中，最小的且奇偶性和 $\textit{card}[\textit{cnt}-1]$ 不同的元素。

*/

// github.com/EndlessCheng/codeforces-go
func maxmiumScore(cards []int, cnt int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	sum := 0
	for _, v := range cards[:cnt] {
		sum += v
	}
	if sum&1 == 0 {
		return sum
	}
	// 在 cards[cnt:] 中找一个最大的且奇偶性和 x 不同的元素，替换 x
	replace := func(x int) {
		for _, v := range cards[cnt:] {
			if v&1 != x&1 {
				ans = max(ans, sum-x+v)
				break
			}
		}
	}
	replace(cards[cnt-1]) // 替换 cards[cnt-1]
	for i := cnt - 2; i >= 0; i-- {
		if cards[i]&1 != cards[cnt-1]&1 { // 找一个最小的且奇偶性不同于 cards[cnt-1] 的元素，将其替换掉
			replace(cards[i])
			break
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
