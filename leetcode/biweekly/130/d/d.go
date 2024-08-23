package main

import "math/bits"

// https://space.bilibili.com/206214
func sumE(k int) (res int) {
	var n, cnt1, sumI int
	for i := bits.Len(uint(k+1)) - 1; i >= 0; i-- {
		c := cnt1<<i + i<<i>>1 // 新增的幂次个数
		if c <= k {
			k -= c
			res += sumI<<i + i*(i-1)/2<<i>>1
			sumI += i   // 之前填的 1 的幂次之和
			cnt1++      // 之前填的 1 的个数
			n |= 1 << i // 填 1
		}
	}
	// 剩余的 k 个幂次，由 n 的低 k 个 1 补充
	for ; k > 0; k-- {
		res += bits.TrailingZeros(uint(n))
		n &= n - 1 // 去掉最低位的 1（置为 0）
	}
	return
}

func pow(x, n, mod int) int {
	res := 1 % mod // 注意 mod 可能等于 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func findProductsOfElements(queries [][]int64) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		er := sumE(int(q[1]) + 1)
		el := sumE(int(q[0]))
		ans[i] = pow(2, er-el, int(q[2]))
	}
	return ans
}
