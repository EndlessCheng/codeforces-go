package main

// https://space.bilibili.com/206214/dynamic
const mod, mx int = 1e9 + 7, 1e4 + 20

var ks [mx][]int
var F, invF [mx]int

func init() {
	for i := 2; i < mx; i++ {
		x := i
		for p := 2; p*p <= x; p++ {
			k := 0
			for ; x%p == 0; x /= p {
				k++
			}
			if k > 0 {
				ks[i] = append(ks[i], k)
			}
		}
		if x > 1 {
			ks[i] = append(ks[i], 1)
		}
	}
	F[0] = 1
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx-1] = pow(F[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func idealArrays(n, maxValue int) (ans int) {
	for m := 1; m <= maxValue; m++ {
		mul := 1
		for _, k := range ks[m] {
			comb := F[n+k-1] * invF[k] % mod * invF[n-1] % mod
			mul = mul * comb % mod
		}
		ans = (ans + mul) % mod
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
