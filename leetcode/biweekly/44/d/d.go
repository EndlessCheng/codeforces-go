package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7
const mx int = 1e4 + 20

var F, invF, lpf [mx + 1]int

func init() {
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	lpf[1] = 1
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func waysToFillArray(qs [][]int) []int {
	ans := make([]int, len(qs))
	for i, q := range qs {
		res, n, x := 1, q[0], q[1]
		for x > 1 {
			p := lpf[x]
			e := 1
			for x /= p; lpf[x] == p; x /= p {
				e++
			}
			res = res * F[n+e-1] % mod * invF[e] % mod * invF[n-1] % mod
		}
		ans[i] = res
	}
	return ans
}

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
