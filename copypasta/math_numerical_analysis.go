package copypasta

import "math"

// 数值分析
// https://en.wikipedia.org/wiki/Numerical_analysis

type mathF func(x float64) float64

// Simpson's 1/3 rule
// https://en.wikipedia.org/wiki/Simpson%27s_rule
// 证明过程 https://phqghume.github.io/2018/05/19/%E8%87%AA%E9%80%82%E5%BA%94%E8%BE%9B%E6%99%AE%E6%A3%AE%E6%B3%95/
func simpson(l, r float64, f mathF) float64 {
	h := (r - l) / 2
	return h * (f(l) + 4*f(l+h) + f(r)) / 3
}

// 不放心的话还可以设置一个最大递归深度 maxDeep
// 15eps 的证明过程 http://www2.math.umd.edu/~mariakc/teaching/adaptive.pdf
func asr(l, r, eps, A float64, f mathF) float64 {
	mid := l + (r-l)/2
	L := simpson(l, mid, f)
	R := simpson(mid, r, f)
	if math.Abs(L+R-A) <= 15*eps {
		return L + R + (L+R-A)/15
	}
	return asr(l, mid, eps/2, L, f) + asr(mid, r, eps/2, R, f)
}

// 自适应辛普森积分 Adaptive Simpson's Rule
// https://en.wikipedia.org/wiki/Adaptive_Simpson%27s_method
// https://oi-wiki.org/math/integral/
// https://cp-algorithms.com/num_methods/simpson-integration.html
// 模板题 https://www.luogu.com.cn/problem/P4525 https://www.luogu.com.cn/problem/P4526 https://www.acwing.com/problem/content/3077/
func Asr(a, b, eps float64, f mathF) float64 {
	return asr(a, b, eps, simpson(a, b, f), f)
}

//

// 多项式插值
// https://en.wikipedia.org/wiki/Polynomial_interpolation

// 拉格朗日插值
// 给定多项式上的 n 个点 (xi,yi)，求 f(k)
// https://en.wikipedia.org/wiki/Lagrange_polynomial
// https://oi-wiki.org/math/poly/lagrange/
// 浅谈几种插值方法 https://www.luogu.com.cn/blog/zhang-xu-jia/ji-zhong-cha-zhi-fang-fa-yang-xie
//
// 模板题 https://www.luogu.com.cn/problem/P4781
// todo https://www.luogu.com.cn/problem/P5667
// 等幂和 https://codeforces.com/problemset/problem/622/F
func lagrangePolynomialInterpolation(xs, ys []int64, k int64) int64 {
	const mod = 998244353

	pow := func(x, n int64) int64 {
		x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	inv := func(a int64) int64 { return pow(a, mod-2) }
	div := func(a, b int64) int64 { return a % mod * inv(b) % mod }

	ans := int64(0)
	for i, xi := range xs {
		a, b := ys[i]%mod, int64(1)
		for j, x := range xs {
			if j != i {
				a = a * (k - x) % mod
				b = b * (xi - x) % mod
			}
		}
		ans += div(a, b)
	}
	ans = (ans%mod + mod) % mod
	return ans
}
