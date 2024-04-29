package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 1001

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func numberOfStableArrays(zero, one, limit int) (ans int) {
	if zero > one {
		zero, one = one, zero
	}
	f0 := make([]int, zero+3)
	for i := (zero-1)/limit + 1; i <= zero; i++ {
		f0[i] = comb(zero-1, i-1)
		for j := 1; j <= (zero-i)/limit; j++ {
			f0[i] = (f0[i] + (1-j%2*2)*comb(i, j)*comb(zero-j*limit-1, i-1)) % mod
		}
	}

	for i := (one-1)/limit + 1; i <= min(one, zero+1); i++ {
		f1 := comb(one-1, i-1)
		for j := 1; j <= (one-i)/limit; j++ {
			f1 = (f1 + (1-j%2*2)*comb(i, j)*comb(one-j*limit-1, i-1)) % mod
		}
		ans = (ans + (f0[i-1]+f0[i]*2+f0[i+1])*f1) % mod
	}
	return (ans + mod) % mod
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

//

func numberOfStableArrays2(zero, one, limit int) int {
	const mod = 1_000_000_007
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i == 0 { // 递归边界
			if k == 1 && j <= limit {
				return 1
			}
			return
		}
		if j == 0 { // 递归边界
			if k == 0 && i <= limit {
				return 1
			}
			return
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		if k == 0 {
			// +mod 保证答案非负
			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
			if i > limit {
				res = (res - dfs(i-limit-1, j, 1) + mod) % mod
			}
		} else {
			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1)) % mod
			if j > limit {
				res = (res - dfs(i, j-limit-1, 0) + mod) % mod
			}
		}
		*p = res // 记忆化
		return
	}
	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
}

func numberOfStableArrays3(zero, one, limit int) (ans int) {
	const mod = 1_000_000_007
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				// + mod 保证答案非负
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}
