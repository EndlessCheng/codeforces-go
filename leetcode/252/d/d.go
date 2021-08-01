package main

/* 动态规划
定义：
- $f[i][0] 表示前 i 项得到的全 $0$ 子序列个数
- f[i][1] 表示前 i 项得到的先 $0$ 后 $1$ 的子序列个数
- f[i][2] 表示前 i 项得到的特殊子序列个数

对于 f[i][0]，当遇到 $0$ 时，有选或不选两种方案，不选 0 时有 f[i][0] = f[i-1][0]，选 0 时，可以单独组成一个子序列，也可以可以与前面的 0 组合，因此有 f[i][0] = f[i-1][0] + 1，两者相加得 f[i][0] = 2\cdot f[i-1][0] + 1。

对于 f[i][1]，当遇到 $1$ 时，有选或不选两种方案，不选 1 时有 f[i][1] = f[i-1][1]，选 1 时，可以单独与前面的 0 组成一个子序列，也可以与前面的 1 组合，因此有 f[i][1] = f[i-1][1] + f[i-1][0]，两者相加得 f[i][1] = 2\cdot f[i-1][1] + f[i-1][0]

f[i][2] 和 f[i][1] 类似，有 f[i][2] = 2\cdot f[i-1][2] + f[i-1][1]

代码实现时，可以把第一维压缩掉。

*/

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

func countSpecialSubsequences(nums []int) int {
	f := [3]int{}
	for _, v := range nums {
		if v == 0 {
			f[0] = (f[0]*2 + 1) % mod
		} else if v == 1 {
			f[1] = (f[1]*2 + f[0]) % mod
		} else {
			f[2] = (f[2]*2 + f[1]) % mod
		}
	}
	return f[2]
}
