package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

func maxNiceDivisors(n int) int {
	ans := 1
	if n > 4 {
		k := (n-5)/3 + 1
		ans = pow(3, k)
		n -= k * 3
	}
	if n > 1 {
		ans = ans * n % mod
	}
	return ans
}

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}
