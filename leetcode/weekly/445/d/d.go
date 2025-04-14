package main

import (
	"fmt"
	"math/big"
	"strings"
)

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const maxN = 333 // 进制转换后的最大长度
const maxB = 10

var comb [maxN + maxB][maxB]int

func init() {
	// 预处理组合数
	for i := 0; i < len(comb); i++ {
		comb[i][0] = 1
		for j := 1; j < min(i+1, maxB); j++ {
			// 注意本题组合数较小，无需取模
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}
}

func trans(s string, b int, inc bool) string {
	x := &big.Int{}
	fmt.Fscan(strings.NewReader(s), x)
	if inc {
		x.Add(x, big.NewInt(1))
	}
	return x.Text(b) // 转成 b 进制
}

func calc(s string, b int, inc bool) (res int) {
	s = trans(s, b, inc)
	// 计算小于 s 的合法数字个数
	// 为什么是小于？注意下面的代码，我们没有统计每个数位都填 s[i] 的情况
	pre := 0
	for i, d := range s {
		hi := int(d - '0')
		if hi < pre {
			break
		}
		m := len(s) - 1 - i
		res += comb[m+b-pre][b-1-pre] - comb[m+b-hi][b-1-hi]
		pre = hi
	}
	return
}

func countNumbers(l, r string, b int) int {
	// 小于 r+1 的合法数字个数 - 小于 l 的合法数字个数
	return (calc(r, b, true) - calc(l, b, false)) % mod
}
