package copypasta

import "math/bits"

/* NTT: number-theoretic transform 快速数论变换
https://en.wikipedia.org/wiki/Discrete_Fourier_transform_(general)#Number-theoretic_transform

NTT 和 FFT 类似，下面的实现在 FFT 代码的基础上稍微修改了下
https://oi-wiki.org/math/poly/ntt/
包含应用及习题 https://cp-algorithms.com/algebra/fft.html#toc-tgt-6
常用素数及原根 http://blog.miskcoo.com/2014/07/fft-prime-table
2281701377 =  17*2^27+1, g = 3, invG = 760567126
1004535809 = 479*2^21+1, g = 3, invG = 334845270
 998244353 = 119*2^23+1, g = 3, invG = 332748118
 167772161 =   5*2^25+1, g = 3, invG = 55924054

模数任意的解决方案 http://blog.miskcoo.com/2015/04/polynomial-multiplication-and-fast-fourier-transform
任意模数 NTT https://www.luogu.com.cn/problem/P4245

NTT vs FFT：对于模板题 https://www.luogu.com.cn/problem/P3803 NTT=1.98s(750ms) FFT=3.63s(1.36s) 括号内是最后一个 case 的运行时间
*/

/* 多项式全家桶
todo 【推荐】https://www.luogu.com.cn/blog/command-block/ntt-yu-duo-xiang-shi-quan-jia-tong
https://cp-algorithms.com/algebra/polynomial.html
http://blog.miskcoo.com/2015/05/polynomial-inverse
http://blog.miskcoo.com/2015/05/polynomial-division
http://blog.miskcoo.com/2015/05/polynomial-multipoint-eval-and-interpolation
关于优化形式幂级数计算的 Newton 法的常数 http://negiizhao.blog.uoj.ac/blog/4671

从拉插到快速插值求值 https://www.luogu.com.cn/blog/command-block/zong-la-cha-dao-kuai-su-cha-zhi-qiu-zhi
快速阶乘算法 https://www.luogu.com.cn/problem/P5282
*/

/* 分治 FFT
todo 半在线卷积小记 https://www.luogu.com.cn/blog/command-block/ban-zai-xian-juan-ji-xiao-ji
CDQ FFT 半在线卷积的O(nlog^2/loglogn)算法 https://www.qaq-am.com/cdqFFT/
模板题 https://www.luogu.com.cn/problem/P4721
*/

/* GF: generating function 生成函数/母函数/多项式计数
https://en.wikipedia.org/wiki/Generating_function

普通生成函数 OGF
指数生成函数 EGF
狄利克雷生成函数 DGFs
todo 【推荐】https://www.luogu.com.cn/blog/command-block/sheng-cheng-han-shuo-za-tan
https://oi-wiki.org/math/gen-func/intro/
【数学理论】浅谈 OI 中常用的一些生成函数运算的合法与正确性 https://rqy.moe/Math/gf_correct/
一些常见数列的生成函数推导 https://www.luogu.com.cn/blog/nederland/girl-friend
狄利克雷相关（含 DGFs）https://www.luogu.com.cn/blog/command-block/gcd-juan-ji-xiao-ji

炫酷反演魔术 https://www.luogu.com.cn/blog/command-block/xuan-ku-fan-yan-mo-shu
反演魔术：反演原理及二项式反演 http://blog.miskcoo.com/2015/12/inversion-magic-binomial-inversion

https://codeforces.com/problemset/problem/958/F3
*/

type ntt struct {
	n        int
	invN     int64
	omega    []int64
	omegaInv []int64
}

const P = 998244353

func pow(x int64, n int) (res int64) {
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % P
		}
		x = x * x % P
	}
	return
}

func newNTT(n int) *ntt {
	const g, invG = 3, 332748118
	omega := make([]int64, n+1)
	omegaInv := make([]int64, n+1)
	for i := 1; i <= n; i <<= 1 {
		omega[i] = pow(g, (P-1)/i)
		omegaInv[i] = pow(invG, (P-1)/i)
	}
	return &ntt{n, pow(int64(n), P-2), omega, omegaInv}
}

func (t *ntt) transform(a, omega []int64) {
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
		wn := omega[l]
		for st := 0; st < t.n; st += l {
			b := a[st:]
			for i, w := 0, int64(1); i < m; i++ {
				d := b[m+i] * w % P
				b[m+i] = (b[i] - d + P) % P
				b[i] = (b[i] + d) % P
				w = w * wn % P
			}
		}
	}
}

func (t *ntt) dft(a []int64) {
	t.transform(a, t.omega)
}

func (t *ntt) idft(a []int64) {
	t.transform(a, t.omegaInv)
	for i, v := range a {
		a[i] = v * t.invN % P
	}
}

// 计算 A(x) 和 B(x) 的卷积 (convolution)
// 入参出参都是次项从低到高的系数
// 模板题 https://www.luogu.com.cn/problem/P3803 https://www.luogu.com.cn/problem/P1919 https://atcoder.jp/contests/practice2/tasks/practice2_f
func polyConvNTT(a, b []int64) []int64 {
	n, m := len(a), len(b)
	limit := 1 << bits.Len(uint(n+m-1))
	A := make([]int64, limit)
	copy(A, a)
	B := make([]int64, limit)
	copy(B, b)
	t := newNTT(limit)
	t.dft(A)
	t.dft(B)
	for i, v := range A {
		A[i] = v * B[i] % P
	}
	t.idft(A)
	return A[:n+m-1]
}

// 计算多个多项式的卷积
// 入参出参都是次项从低到高的系数
func polyConvNTTs(coefs [][]int64) []int64 {
	var f func(l, r int) []int64
	f = func(l, r int) []int64 {
		if l == r {
			return coefs[l-1] // coefs start at 0
		}
		mid := (l + r) >> 1
		return polyConvNTT(f(l, mid), f(mid+1, r))
	}
	return f(1, len(coefs))
}
