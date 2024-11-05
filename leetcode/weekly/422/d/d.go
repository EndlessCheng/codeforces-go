package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 40

var fac, invF [mx + 1]int

func init() {
	fac[0] = 1
	for i := 1; i <= mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invF[mx] = pow(fac[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func countBalancedPermutations(num string) int {
	cnt := [10]int{}
	total := 0
	for _, c := range num {
		cnt[c-'0']++
		total += int(c - '0')
	}

	if total%2 > 0 {
		return 0
	}

	n := len(num)
	n1 := n / 2
	f := make([][]int, n1+1)
	for i := range f {
		f[i] = make([]int, total/2+1)
	}
	f[0][0] = 1
	sc := 0
	s := 0
	for i, c := range cnt {
		sc += c
		s += c * i
		// 保证 left2 <= n-n1，即 left1 >= sc-(n-n1)
		for left1 := min(sc, n1); left1 >= max(sc-(n-n1), 0); left1-- {
			left2 := sc - left1
			// 保证分给第二个集合的元素和 <= total/2，即 leftS >= s-total/2
			for leftS := min(s, total/2); leftS >= max(s-total/2, 0); leftS-- {
				res := 0
				for k := max(c-left2, 0); k <= min(c, left1) && k*i <= leftS; k++ {
					res = (res + f[left1-k][leftS-k*i]*invF[k]%mod*invF[c-k]) % mod
				}
				f[left1][leftS] = res
			}
		}
	}
	return fac[n1] * fac[n-n1] % mod * f[n1][total/2] % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func countBalancedPermutations2(num string) int {
	cnt := [10]int{}
	total := 0
	for _, c := range num {
		cnt[c-'0']++
		total += int(c - '0')
	}

	if total%2 > 0 {
		return 0
	}

	for i := 1; i < 10; i++ {
		cnt[i] += cnt[i-1]
	}

	n := len(num)
	n1 := n / 2
	memo := [10][][]int{}
	for i := range memo {
		memo[i] = make([][]int, n1+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, total/2+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 // -1 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, left1, leftS int) int {
		if i < 0 {
			if leftS > 0 {
				return 0
			}
			return 1
		}
		p := &memo[i][left1][leftS]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		c := cnt[i]
		if i > 0 {
			c -= cnt[i-1]
		}
		left2 := cnt[i] - left1
		for k := max(c-left2, 0); k <= min(c, left1) && k*i <= leftS; k++ {
			r := dfs(i-1, left1-k, leftS-k*i)
			res = (res + r*invF[k]%mod*invF[c-k]) % mod
		}
		*p = res // 记忆化
		return res
	}
	return fac[n1] * fac[n-n1] % mod * dfs(9, n1, total/2) % mod
}
