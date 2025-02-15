package copypasta

import (
	"math"
	"math/bits"
	"slices"
)

/* FFT: fast Fourier transform 快速傅里叶变换
https://en.wikipedia.org/wiki/Fast_Fourier_transform

视频讲解 https://www.youtube.com/watch?v=h7apO7q16V0
中文翻译 https://www.bilibili.com/video/BV1za411F76U/
代码参考：一小时学会快速傅里叶变换 https://zhuanlan.zhihu.com/p/31584464
todo 优化 https://github.com/atcoder/ac-library/blob/master/atcoder/convolution.hpp
https://www.luogu.com.cn/blog/105254/qian-tan-fft-zong-ft-dao-fft
多项式基础：插值、函数逼近、快速傅里叶变换 (蒋炎岩) https://www.bilibili.com/video/BV1a14y1M7v1/
从多项式乘法到快速傅里叶变换 http://blog.miskcoo.com/2015/04/polynomial-multiplication-and-fast-fourier-transform
整数乘法的长征 by EI https://www.cnblogs.com/Elegia/p/18020040/integer-multiplication
https://codeforces.com/blog/entry/43499 https://codeforces.com/blog/entry/48798
https://oi-wiki.org/math/poly/fft/
https://cp-algorithms.com/algebra/fft.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FFT.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/Polynomial.java.html
用 FFT 做字符串匹配 https://zhuanlan.zhihu.com/p/267765026
Arrow product: How to enumerate directed graphs https://codeforces.com/blog/entry/115617
能否在 O(nlogn) 时间内实现大数的进制转换？ https://www.zhihu.com/question/395557989
信息论的角度理解 https://www.zhihu.com/question/394657296/answer/2302858330

todo https://github.com/OI-wiki/gitment/discussions/670#discussioncomment-4496021
 若多项式系数没有复数的话，可以构造多项式 H(x)=F(x)+G(x)i，把 F(x) 放实部，把 G(x) 放虚部，然后对 H(x) 跑一遍 DFT，之后把 H(x) 平方一下，可以得到
 H2(x)=F2(x)−G2(x)+2F(x)G(x)i
 可以发现它的虚部就是多项式的乘积，于是对 H2(x) 跑一遍 IDFT，之后把虚部取出来除以 2，得到的就是 F(x)G(x)
 这样只需要跑 2 次 FFT，常数更小一些（如果是对 F(x) 和 G(x) 分别跑 DFT，然后乘起来，再跑 IDFT 的话，需要跑 3 次 FFT）
https://www.luogu.com.cn/blog/command-block/fft-xue-xi-bi-ji

有关快速数论变换（NTT）以及多项式运算的内容见 math_ntt.go

模板题 https://www.luogu.com.cn/problem/P3803
todo 推式子 https://www.luogu.com.cn/problem/P3338 花絮 https://zhuanlan.zhihu.com/p/349249817
 https://codeforces.com/problemset/problem/993/E
 https://codeforces.com/gym/104081/problem/K
 https://atcoder.jp/contests/abc196/tasks/abc196_f 子串失配个数
*/

type fft struct {
	n               int
	omega, omegaInv []complex128
}

func newFFT(n int) *fft {
	omega := make([]complex128, n)
	omegaInv := make([]complex128, n)
	for i := range omega {
		sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
		omega[i] = complex(cos, sin)
		omegaInv[i] = complex(cos, -sin)
	}
	return &fft{n, omega, omegaInv}
}

// 注：下面 swap 的代码，另一种写法是初始化每个 i 对应的 j https://blog.csdn.net/Flag_z/article/details/99163939
func (t *fft) transform(a, omega []complex128) {
	n := t.n
	for i, j := 0, 0; i < n; i++ {
		if i > j { // 保证同一对元素只交换一次
			a[i], a[j] = a[j], a[i]
		}
		for l := n / 2; ; l /= 2 {
			j ^= l
			if j >= l {
				break
			}
		}
	}
	for l := 2; l <= n; l *= 2 {
		m := l / 2
		for st := 0; st < n; st += l {
			b := a[st:]
			for i := range m {
				v := omega[n/l*i] * b[m+i]
				b[m+i] = b[i] - v
				b[i] += v
			}
		}
	}
}

func (t *fft) dft(a []complex128) {
	t.transform(a, t.omega)
}

func (t *fft) idft(a []complex128) {
	t.transform(a, t.omegaInv)
	cn := complex(float64(t.n), 0)
	for i := range a {
		a[i] /= cn
	}
}

// 计算两个多项式 A(x) 和 B(x) 的乘积，也叫卷积 (convolution)
// 结果多项式 C(x)，其中 x^k 的系数为
//     c[k] = ∑a[i]*b[k-i], i=0~k
// 入参出参都是次项从低到高的系数
//
// EXTRA: 对数组 a 的频率数组 F(x) 计算自卷积，
//        得到的结果 G(x) 表示两数之和等于 x 的方案数（这两个数之间没有位置约束）
// https://atcoder.jp/contests/abc392/tasks/abc392_g
//
// 关于滑动窗口点积，见后面
func polyConvFFT(a, b []int) []int {
	n, m := len(a), len(b)
	limit := 1 << bits.Len(uint(n+m-1))
	A := make([]complex128, limit)
	for i, v := range a {
		A[i] = complex(float64(v), 0)
	}
	B := make([]complex128, limit)
	for i, v := range b {
		B[i] = complex(float64(v), 0)
	}
	t := newFFT(limit)
	t.dft(A)
	t.dft(B)
	for i := range A {
		A[i] *= B[i]
	}
	t.idft(A)
	conv := make([]int, n+m-1)
	for i := range conv {
		conv[i] = int(math.Round(real(A[i]))) // % mod
	}
	return conv
}

// 滑动窗口点积
// 对 a 的每个长为 len(b) 的连续子数组 subarray，计算 subarray 与 b 的点积 c[i]
// 也就是求 c[i] = ∑a[i+j]*b[j], j=0~m-1
// 做法：把 b 反转后求卷积
// EXTRA：滑动窗口 0-1 异或，利用恒等式 a^b = a+b - (a&b)*2 转成加法（前缀和）与乘法（卷积）
// https://atcoder.jp/contests/abc196/tasks/abc196_f 2274=CF2431
func slidingWindowDotProduct(a, b []int) []int {
	b = slices.Clone(b) // 避免修改原数组
	slices.Reverse(b)
	c := polyConvFFT(a, b)
	return c[len(b)-1 : len(a)]
}

// 计算多个多项式的卷积
// 入参出参都是次项从低到高的系数
// https://codeforces.com/contest/958/problem/F3 可重集大小为 k 的不同子集个数
func polyConvFFTs(coefs [][]int) []int {
	n := len(coefs)
	if n == 1 {
		return coefs[0]
	}
	return polyConvFFT(polyConvFFTs(coefs[:n/2]), polyConvFFTs(coefs[n/2:]))
}

// 有关快速数论变换（NTT）以及更多多项式运算的内容见 math_ntt.go
// 如果题目没有取模，但保证答案小于 mod，也可以用 NTT 以加快速度
// 效率对比 (go 1.20.6)
// FFT 789ms https://atcoder.jp/contests/abc196/submissions/62740907
// NTT 584ms https://atcoder.jp/contests/abc196/submissions/62740827
