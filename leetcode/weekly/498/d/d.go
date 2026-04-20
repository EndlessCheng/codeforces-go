package main

import (
	"math"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func countGoodIntegersOnPath1(l, r int64, directions string) int64 {
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)

	inPath := make([]bool, n)
	pos := n - 16 // 右下角是下标 n-1，那么左上角是下标 n-16
	for _, d := range directions {
		if pos >= 0 { // 只需要对网格图中的后 n 个格子做标记
			inPath[pos] = true // 标记在路径中的格子
		}
		if d == 'R' { // 往右
			pos++
		} else { // 往下
			pos += 4 // 相当于往右数 4 个位置
		}
	}
	inPath[n-1] = true // 终点一定在路径中

	diffLH := n - len(lowS)
	memo := make([][10]int64, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int64
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int64) {
		if i == n { // 成功到达终点
			return 1 // 找到了一个好整数
		}

		if !limitLow && !limitHigh {
			p := &memo[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		if inPath[i] { // 当前位置在路径中
			d = max(d, pre) // 当前位置填的数必须 >= pre
		}
		for ; d <= hi; d++ {
			p := pre
			if inPath[i] {
				p = d
			}
			res += dfs(i+1, p, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	return dfs(0, 0, true, true)
}

const maxM = 7

var comb [maxM + 10][maxM + 1]int

func init() {
	// 预处理组合数
	for i := range comb {
		comb[i][0] = 1
		for j := 1; j < min(i+1, len(comb[i])); j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}
}

func countGoodIntegersOnPath(l, r int64, directions string) int64 {
	highS := strconv.FormatInt(r+1, 10) // 注意这里加一了
	n := len(highS)
	lowS := strconv.FormatInt(l, 10)
	lowS = strings.Repeat("0", n-len(lowS)) + lowS

	inPath := make([]bool, n)
	pos := n - 16 // 右下角是下标 n-1，那么左上角是下标 n-16
	for _, d := range directions {
		if pos >= 0 { // 只需要对网格图中的后 n 个格子做标记
			inPath[pos] = true // 标记在路径中的格子
		}
		if d == 'R' { // 往右
			pos++
		} else { // 往下
			pos += 4 // 相当于往右数 4 个位置
		}
	}
	inPath[n-1] = true // 终点一定在路径中

	// suf[i] 表示后缀 [i, n-1] 在路径中的下标个数
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		if inPath[i] {
			suf[i]++
		}
	}

	// 计算小于 r 的合法整数个数
	solve := func(r string) (res int) {
		pre := 0
		for i, ch := range r {
			hi := int(ch - '0')
			m := suf[i+1]
			if !inPath[i] {
				res += hi * comb[m+9-pre][m] * int(math.Pow10(n-1-i-m))
				continue
			}
			if hi < pre {
				break
			}
			res += (comb[m+10-pre][m+1] - comb[m+10-hi][m+1]) * int(math.Pow10(n-1-i-m))
			pre = hi // 这一位填 hi，继续计算剩余数位的方案数
		}
		return res
	}

	return int64(solve(highS) - solve(lowS))
}
