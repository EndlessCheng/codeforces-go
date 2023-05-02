package copypasta

import (
	"math"
	"math/bits"
)

/* FFT: fast Fourier transform 快速傅里叶变换
https://en.wikipedia.org/wiki/Fast_Fourier_transform

【推荐】一小时学会快速傅里叶变换 https://zhuanlan.zhihu.com/p/31584464
多项式基础：插值、函数逼近、快速傅里叶变换 (蒋炎岩) https://www.bilibili.com/video/BV1a14y1M7v1/
为什么 FFT 可以加速卷积运算 https://www.zhihu.com/question/394657296/answer/2329522108
傅里叶变换学习笔记 https://www.luogu.com.cn/blog/command-block/fft-xue-xi-bi-ji
从多项式乘法到快速傅里叶变换 http://blog.miskcoo.com/2015/04/polynomial-multiplication-and-fast-fourier-transform
优化技巧 https://www.luogu.com.cn/blog/105254/qian-tan-fft-zong-ft-dao-fft
https://codeforces.com/blog/entry/43499 https://codeforces.com/blog/entry/48798
https://oi-wiki.org/math/poly/fft/
https://cp-algorithms.com/algebra/fft.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/FFT.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/Polynomial.java.html
用 FFT 做字符串匹配 https://zhuanlan.zhihu.com/p/267765026
Arrow product: How to enumerate directed graphs https://codeforces.com/blog/entry/115617

todo https://github.com/OI-wiki/gitment/discussions/670#discussioncomment-4496021
 若多项式系数没有复数的话，可以构造多项式 H(x)=F(x)+G(x)i，把 F(x) 放实部，把 G(x) 放虚部，然后对 H(x) 跑一遍 DFT，之后把 H(x) 平方一下，可以得到
 H2(x)=F2(x)−G2(x)+2F(x)G(x)i
 可以发现它的虚部就是多项式的乘积，于是对 H2(x) 跑一遍 IDFT，之后把虚部取出来除以 2，得到的就是 F(x)G(x)
 这样只需要跑 2 次 FFT，常数更小一些（如果是对 F(x) 和 G(x) 分别跑 DFT，然后乘起来，再跑 IDFT 的话，需要跑 3 次 FFT）

有关快速数论变换（NTT）以及多项式运算的内容见 math_ntt.go

模板题 https://www.luogu.com.cn/problem/P3803
todo 推式子 https://www.luogu.com.cn/problem/P3338 花絮 https://zhuanlan.zhihu.com/p/349249817
todo https://codeforces.com/problemset/problem/993/E
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
	for i, j := 0, 0; i < t.n; i++ {
		if i > j {
			a[i], a[j] = a[j], a[i]
		}
		for l := t.n >> 1; ; l >>= 1 {
			j ^= l
			if j >= l {
				break
			}
		}
	}
	for l := 2; l <= t.n; l <<= 1 {
		m := l >> 1
		for st := 0; st < t.n; st += l {
			b := a[st:]
			for i := 0; i < m; i++ {
				d := omega[t.n/l*i] * b[m+i]
				b[m+i] = b[i] - d
				b[i] += d
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

// 计算 A(x) 和 B(x) 的卷积 (convolution)
// c[i] = ∑a[k]*b[i-k], k=0..i
// 入参出参都是次项从低到高的系数
// 建议全程用 int64
func polyConvFFT(a, b []int64) []int64 {
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
	conv := make([]int64, n+m-1)
	for i := range conv {
		conv[i] = int64(math.Round(real(A[i]))) // % mod
	}
	return conv
}

// 计算多个多项式的卷积
// 入参出参都是次项从低到高的系数
// 可重集大小为 k 的不同子集个数 https://codeforces.com/contest/958/problem/F3
func polyConvFFTs(coefs [][]int64) []int64 {
	n := len(coefs)
	if n == 1 {
		return coefs[0]
	}
	return polyConvFFT(polyConvFFTs(coefs[:n/2]), polyConvFFTs(coefs[n/2:]))
}

// 有关快速数论变换（NTT）以及更多多项式运算的内容见 math_ntt.go
