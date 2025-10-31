package copypasta

import (
	"math"
	"math/big"
)

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
// 给定（同余）多项式上的 n 个点 (xi,yi)，我们可以得到一个 n-1 次多项式，
// 利用拉格朗日插值，可以在不用高斯消元的情况下，求出 f(k) 的值
// 时间复杂度 O(n^2)
//
// https://en.wikipedia.org/wiki/Lagrange_polynomial
// https://oi-wiki.org/math/poly/lagrange/
// 浅谈几种插值方法 https://www.luogu.com.cn/blog/zhang-xu-jia/ji-zhong-cha-zhi-fang-fa-yang-xie
//
// 模板题 https://www.luogu.com.cn/problem/P4781
// todo https://www.luogu.com.cn/problem/P5667
// https://codeforces.com/problemset/problem/1155/E 2200 交互 找零点
// https://codeforces.com/problemset/problem/622/F 2600 等幂和
// https://codeforces.com/problemset/problem/995/F 2700
// https://projecteuler.net/problem=101
func lagrangePolynomialInterpolation(xs, ys []int, k int) (fk int) {
	for i, xi := range xs {
		a, b := 1, 1
		for j, xj := range xs {
			if j != i {
				a = a * (k - xj) % mod  // 分子
				b = b * (xi - xj) % mod // 分母
			}
		}
		fk += a * pow(b, mod-2) % mod * ys[i] % mod // 也可以把 a 初始化成 ys[i]%mod
	}
	fk = (fk%mod + mod) % mod
	return
}

func lagrangePolynomialInterpolationBig(xs, ys []int, k int) *big.Rat {
	fk := big.NewRat(0, 1)
	for i, xi := range xs {
		a, b := big.NewInt(int64(ys[i])), big.NewInt(1)
		for j, xj := range xs {
			if j != i {
				a.Mul(a, big.NewInt(int64(k-xj)))  // 分子
				b.Mul(b, big.NewInt(int64(xi-xj))) // 分母
			}
		}
		fk.Add(fk, new(big.Rat).SetFrac(a, b))
	}
	return fk
}
