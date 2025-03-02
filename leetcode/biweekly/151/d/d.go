package main

import "slices"

// https://space.bilibili.com/206214
// 预处理交替排列的方案数
var f = []int{1}

func init() {
	for i := 1; f[len(f)-1] < 1e15; i++ {
		f = append(f, f[len(f)-1]*i)
		f = append(f, f[len(f)-1]*i)
	}
}

func permute(n int, K int64) []int {
	// k 改成从 0 开始，方便计算
	k := int(K - 1)
	if n < len(f) && k >= f[n]*(2-n%2) { // n 是偶数的时候，方案数乘以 2
		return nil
	}

	// cand 表示剩余未填入 ans 的数字
	// cand[0] 保存偶数，cand[1] 保存奇数
	cand := [2][]int{}
	for i := 2; i <= n; i += 2 {
		cand[0] = append(cand[0], i)
	}
	for i := 1; i <= n; i += 2 {
		cand[1] = append(cand[1], i)
	}

	ans := make([]int, n)
	parity := 1 // 当前要填入 ans[i] 的数的奇偶性
	for i := range n {
		j := 0
		if n-1-i < len(f) {
			// 比如示例 1，按照第一个数分组，每一组的大小都是 size=2
			// 知道 k 和 size 就知道我们要去哪一组
			size := f[n-1-i]
			j = k / size // 去第 j 组
			k %= size
			// n 是偶数的情况，第一个数既可以填奇数又可以填偶数，要特殊处理
			if n%2 == 0 && i == 0 {
				parity = 1 - j%2
				j /= 2
			}
		} // else n 很大的情况下，只能按照 1,2,3,... 的顺序填
		ans[i] = cand[parity][j]
		cand[parity] = slices.Delete(cand[parity], j, j+1)
		parity ^= 1 // 下一个数的奇偶性
	}
	return ans
}
