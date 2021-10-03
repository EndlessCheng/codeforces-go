package main

/* 关键在于求出回合数

由于我们只关心总和能否被 $3$ 整除，我们可以将 $\textit{stones}[i]$ 按照模 $3$ 的结果分为 $3$ 组，即 $0$、$1$ 和 $2$。

根据题意，第一回合不能移除 $0$，否则直接输掉游戏，因此第一回合只能移除 $1$ 或者 $2$。我们可以枚举这两种情况，如果其中一种可以让 Alice 获胜就返回 $\texttt{true}$，否则返回 $\texttt{false}$。

以第一回合移除 $1$ 为例，在不考虑移除 $0$ 的前提下，后面的移除由于要满足不能被 $3$ 整除，因此移除的石子是固定的，整体构成一个 $1121212\dots$ 循环的序列。

对于 $0$，由于移除之后不会改变总和模 $3$ 的结果，因此不会改变后续 $1$ 和 $2$ 的移除顺序，所以可以在任意非开头的位置插入 $0$。

Bob 为了不让 Alice 获胜，若游戏还可以进行，那么 Bob 必然会继续移除石子，因此我们要求的就是最大的回合数，这相当于 $1121212\dots$ 序列的最长长度，加上 $0$ 的个数。

若回合数为奇数，且还有剩余石子，那么 Bob 只能移除一枚让总和被 $3$ 整除的石子，Alice 获胜；否则 Bob 获胜。

*/

// github.com/EndlessCheng/codeforces-go
func check(c [3]int) bool {
	if c[1] == 0 {
		return false
	}
	c[1]-- // 开头为 1
	turn := 1 + min(c[1], c[2])*2 + c[0] // 计算回合数
	if c[1] > c[2] {
		turn++
		c[1]--
	}
	return turn%2 == 1 && c[1] != c[2] // 回合数为奇数，且还有剩余石子
}

func stoneGameIX(stones []int) bool {
	c := [3]int{}
	for _, v := range stones {
		c[v%3]++
	}
	return check(c) || check([3]int{c[0], c[2], c[1]}) // 枚举第一回合移除的是 1 还是 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
