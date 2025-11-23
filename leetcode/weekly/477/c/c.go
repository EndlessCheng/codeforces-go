package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const maxN = 100_001

var pow10 = [maxN]int{1}

func init() {
	// 预处理 10 的幂
	for i := 1; i < maxN; i++ {
		pow10[i] = pow10[i-1] * 10 % mod
	}
}

func sumAndMultiply(s string, queries [][]int) []int {
	n := len(s)
	sumD := make([]int, n+1)       // s 的前缀和
	preNum := make([]int, n+1)     // s 的前缀对应的数字（模 mod）
	sumNonZero := make([]int, n+1) // s 的前缀的非零数字个数
	for i, ch := range s {
		d := int(ch - '0')
		sumD[i+1] = sumD[i] + d
		preNum[i+1] = preNum[i]
		sumNonZero[i+1] = sumNonZero[i]
		if d > 0 {
			preNum[i+1] = (preNum[i]*10 + d) % mod
			sumNonZero[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]+1
		length := sumNonZero[r] - sumNonZero[l]
		x := preNum[r] - preNum[l]*pow10[length]%mod // 注意结果可能是负数，所以下面 +mod
		ans[i] = (x + mod) * (sumD[r] - sumD[l]) % mod
	}
	return ans
}
