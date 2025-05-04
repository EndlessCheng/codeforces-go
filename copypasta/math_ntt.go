package copypasta

import (
	"math/big"
	"math/bits"
	"slices"
)

/* NTT: number-theoretic transform 快速数论变换
https://en.wikipedia.org/wiki/Discrete_Fourier_transform_(general)#Number-theoretic_transform
从傅里叶变换到 998244353 https://www.bilibili.com/read/cv2289955/
硬核理解快速数论变换 https://www.bilibili.com/video/BV1eT411M7Fp/

NTT 和 FFT 类似，下面的实现在 FFT 代码的基础上稍微修改了下
https://oi-wiki.org/math/poly/ntt/
包含应用及习题 https://cp-algorithms.com/algebra/fft.html#toc-tgt-6
常用素数及原根 http://blog.miskcoo.com/2014/07/fft-prime-table
2281701377 =  17*2^27+1, g = 3, invG = 760567126
1004535809 = 479*2^21+1, g = 3, invG = 334845270
 998244353 = 119*2^23+1, g = 3, invG = 332748118
 167772161 =   5*2^25+1, g = 3, invG = 55924054
P-1 包含大量因子 2，便于分治

模数任意的解决方案 http://blog.miskcoo.com/2015/04/polynomial-multiplication-and-fast-fourier-transform
任意模数 NTT https://www.luogu.com.cn/problem/P4245

NTT vs FFT：对于模板题 https://www.luogu.com.cn/problem/P3803 NTT=1.98s(750ms) FFT=3.63s(1.36s) 括号内是最后一个 case 的运行时间

卡常技巧
A modulo multiplication method that is 2x faster than compiler implementation https://codeforces.com/blog/entry/111566

多项式全家桶
【推荐】https://www.luogu.com.cn/blog/command-block/ntt-yu-duo-xiang-shi-quan-jia-tong
https://blog.orzsiyuan.com/search/%E5%A4%9A%E9%A1%B9%E5%BC%8F/2/
模板 https://blog.orzsiyuan.com/archives/Polynomial-Template/
jiangly 模板 https://atcoder.jp/contests/arc163/submissions/45737810
https://blog.csdn.net/weixin_43973966/article/details/88996932
https://cp-algorithms.com/algebra/polynomial.html
http://blog.miskcoo.com/2015/05/polynomial-inverse
http://blog.miskcoo.com/2015/05/polynomial-division
http://blog.miskcoo.com/2015/05/polynomial-multipoint-eval-and-interpolation
关于优化形式幂级数计算的 Newton 法的常数 http://negiizhao.blog.uoj.ac/blog/4671
todo 卡常板子 https://judge.yosupo.jp/submission/65290

从拉插到快速插值求值 https://www.luogu.com.cn/blog/command-block/zong-la-cha-dao-kuai-su-cha-zhi-qiu-zhi
浅谈多项式复合和拉格朗日反演 https://www.luogu.com.cn/blog/your-alpha1022/qian-tan-duo-xiang-shi-fu-ge-hu-la-ge-lang-ri-fan-yan
快速阶乘算法 https://www.luogu.com.cn/problem/P5282
调和级数求和 https://www.luogu.com.cn/problem/P5702

生成函数/母函数 GF generating function
https://en.wikipedia.org/wiki/Generating_function
todo generatingfunctionology https://www2.math.upenn.edu/~wilf/gfologyLinked2.pdf

GF 右移 k 位：乘以 x^k，例如 [3,1,4,1,5] * [0,0,1] = [0,0,3,1,4,1,5]
GF 左移 k 位：减去前 k 项后，再除以 x^k

GF 的前缀和：除以 1-x，例如 [3,1,4,1,5] * [1,1,1,...] = [3,4,8,9,14,14,14,...]
GF 的差分：乘以 1-x，例如 [3,1,4,1,5] * [1,-1] = [3,-2,3,-3,4,-5] （去掉首尾两项就是一阶差分）

普通生成函数 OGF Ordinary generating function
指数生成函数 EGF Exponential generating function 
- 入门题 https://codeforces.com/problemset/problem/891/E 3000
狄利克雷生成函数 DGFs
todo 【推荐】https://www.luogu.com.cn/blog/command-block/sheng-cheng-han-shuo-za-tan
 【推荐】数数入门 https://www.luogu.com.cn/blog/CJL/conut-ru-men
 https://www.bilibili.com/video/BV1Zg411T7Eq
https://oi-wiki.org/math/gen-func/intro/
[Tutorial] Generating Functions in Competitive Programming https://codeforces.com/blog/entry/77468 https://codeforces.com/blog/entry/77551
Some useful power series https://www2.math.upenn.edu/~wilf/gfologyLinked2.pdf
OGF 展开方式 https://oi-wiki.org/math/gen-func/ogf/#_5
【数学理论】浅谈 OI 中常用的一些生成函数运算的合法与正确性 https://rqy.moe/Math/gf_correct/ https://www.luogu.com.cn/blog/lx-2003/gf-correct
一些常见数列的生成函数推导 https://www.luogu.com.cn/blog/nederland/girl-friend
狄利克雷相关（含 DGFs）https://www.luogu.com.cn/blog/command-block/gcd-juan-ji-xiao-ji
狄利克雷生成函数浅谈 https://www.luogu.com.cn/blog/gxy001/di-li-ke-lei-sheng-cheng-han-shuo-qian-tan
生成函数在背包问题中的应用 https://zykykyk.github.io/post/%E7%94%9F%E6%88%90%E5%87%BD%E6%95%B0%E5%9C%A8%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98%E4%B8%AD%E7%9A%84%E5%BA%94%E7%94%A8/
生成函数的背包计数问题 https://www.cnblogs.com/ErkkiErkko/p/10838697.html
OGFs, EGFs, differentiation and Taylor shifts https://codeforces.com/blog/entry/99646
A problem collection of ODE and differential technique https://codeforces.com/blog/entry/76447
Optimal Algorithm on Polynomial Composite Set Power Series https://codeforces.com/blog/entry/92183
On linear recurrences and the math behind them https://codeforces.com/blog/entry/100158
载谭 Binomial Sum：多项式复合、插值与泰勒展开 https://www.luogu.com.cn/blog/EntropyIncreaser/zai-tan-binomial-sum-duo-xiang-shi-fu-ge-cha-zhi-yu-tai-lei-zhan-kai
How to composite (some) polynomials faster? https://codeforces.com/blog/entry/126124
https://maspypy.com/%e5%a4%9a%e9%a0%85%e5%bc%8f%e3%83%bb%e5%bd%a2%e5%bc%8f%e7%9a%84%e3%81%b9%e3%81%8d%e7%b4%9a%e6%95%b0%e6%95%b0%e3%81%88%e4%b8%8a%e3%81%92%e3%81%a8%e3%81%ae%e5%af%be%e5%bf%9c%e4%bb%98%e3%81%91
https://maspypy.com/%e5%a4%9a%e9%a0%85%e5%bc%8f%e3%83%bb%e5%bd%a2%e5%bc%8f%e7%9a%84%e3%81%b9%e3%81%8d%e7%b4%9a%e6%95%b0%ef%bc%88%ef%bc%92%ef%bc%89%e5%bc%8f%e5%a4%89%e5%bd%a2%e3%81%ab%e3%82%88%e3%82%8b%e8%a7%a3%e6%b3%95
https://maspypy.com/%e5%a4%9a%e9%a0%85%e5%bc%8f%e3%83%bb%e5%bd%a2%e5%bc%8f%e7%9a%84%e3%81%b9%e3%81%8d%e7%b4%9a%e6%95%b0%ef%bc%88%ef%bc%93%ef%bc%89%e7%b7%9a%e5%bd%a2%e6%bc%b8%e5%8c%96%e5%bc%8f%e3%81%a8%e5%bd%a2%e5%bc%8f

炫酷反演魔术 https://www.luogu.com.cn/blog/command-block/xuan-ku-fan-yan-mo-shu

Min-Max容斥
https://www.luogu.com.cn/blog/Troverld/Min-Max-Inclusion-and-Exclusion
https://www.luogu.com.cn/blog/command-block/min-max-rong-chi-xiao-ji
https://lnrbhaw.github.io/2019/01/05/Min-Max%E5%AE%B9%E6%96%A5%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/

拉格朗日反演 扩展拉格朗日反演
证明 https://www.cnblogs.com/judge/p/10652738.html
多项式拉格朗日反演与复合逆 https://blog.csdn.net/C20190102/article/details/107279319
点双连通图计数 https://www.luogu.com.cn/problem/P5827
边双连通图计数 https://www.luogu.com.cn/problem/P5828

分治 FFT
todo 半在线卷积小记 https://www.luogu.com.cn/blog/command-block/ban-zai-xian-juan-ji-xiao-ji
CDQ FFT 半在线卷积的O(nlog^2/loglogn)算法 https://www.qaq-am.com/cdqFFT/
模板题 https://www.luogu.com.cn/problem/P4721
https://atcoder.jp/contests/abc267/tasks/abc267_h

---

入门题 https://atcoder.jp/contests/abc248/tasks/abc248_c
- 容斥做法 https://atcoder.jp/contests/abc248/editorial/3818
https://atcoder.jp/contests/abc276/tasks/abc276_g https://atcoder.jp/contests/abc276/editorial/5185
常系数齐次线性递推 https://www.luogu.com.cn/problem/P4723
题单 https://zhuanlan.zhihu.com/p/575825170
题单 https://ac.nowcoder.com/acm/contest/24710#question
https://codeforces.com/problemset/problem/958/F3
todo https://codeforces.com/contest/438/problem/E
todo https://leetcode.cn/contest/hust_1024_2023/problems/kzxnaX/
 https://leetcode.cn/circle/discuss/NEDYEC/ 
 https://oi-wiki.org/math/combinatorics/partition/#%E4%BA%94%E8%BE%B9%E5%BD%A2%E6%95%B0%E5%AE%9A%E7%90%86
 https://leetcode.cn/circle/discuss/Qvv72W/view/DJalmi/
题单 https://www.luogu.com.cn/training/128735
题单 https://www.luogu.com.cn/training/1008
题单 https://www.luogu.com.cn/training/3295
*/

const P = 998244353

func nttPow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % P
		}
		x = x * x % P
	}
	return
}

var omega, omegaInv [31]int // 多开一点空间

func init() {
	const g, invG = 3, 332748118
	for i := 1; i < len(omega); i++ {
		omega[i] = nttPow(g, (P-1)/(1<<i))
		omegaInv[i] = nttPow(invG, (P-1)/(1<<i))
	}
}

type ntt struct {
	n    int
	invN int
}

func newNTT(n int) ntt { return ntt{n, nttPow(n, P-2)} }

// 注：下面 swap 的代码，另一种写法是初始化每个 i 对应的 j https://blog.csdn.net/Flag_z/article/details/99163939
// 由于不是性能瓶颈，实测对性能影响不大
func (t ntt) transform(a, omega []int) {
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
	for l, li := 2, 1; l <= t.n; l <<= 1 {
		m := l >> 1
		wn := omega[li]
		li++
		for st := 0; st < t.n; st += l {
			b := a[st:]
			for i, w := 0, 1; i < m; i++ {
				d := b[m+i] * w % P
				b[m+i] = (b[i] - d + P) % P
				b[i] = (b[i] + d) % P
				w = w * wn % P
			}
		}
	}
}

func (t ntt) dft(a []int) {
	t.transform(a, omega[:])
}

func (t ntt) idft(a []int) {
	t.transform(a, omegaInv[:])
	for i, v := range a {
		a[i] = v * t.invN % P
	}
}

type poly []int

func (a poly) resize(n int) poly {
	b := make(poly, n)
	copy(b, a)
	return b
}

// 计算 A(x) 和 B(x) 的卷积 (convolution)
// c[k] = ∑a[i]*b[k-i], i=0..k
// 如果求 ∑a[i]*b[i]，可以把 b 反转后再求卷积
// 入参出参都是次项从低到高的系数
// 模板题 https://judge.yosupo.jp/problem/convolution_mod
//       https://www.luogu.com.cn/problem/P3803
//       https://www.luogu.com.cn/problem/P1919 
//       https://atcoder.jp/contests/practice2/tasks/practice2_f
func (a poly) conv(b poly) poly {
	n, m := len(a), len(b)
	limit := 1 << bits.Len(uint(n+m-1))
	A := a.resize(limit)
	B := b.resize(limit)
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
func polyConvNTTs(coefs []poly) poly {
	n := len(coefs)
	if n == 1 {
		return coefs[0]
	}
	return polyConvNTTs(coefs[:n/2]).conv(polyConvNTTs(coefs[n/2:]))
}

func (a poly) reverse() poly {
	slices.Reverse(a)
	return a
}

func (a poly) reverseCopy() poly {
	b := slices.Clone(a)
	slices.Reverse(b)
	return b
}

func (a poly) neg() poly {
	b := make(poly, len(a))
	for i, v := range a {
		if v > 0 {
			b[i] = P - v
		}
	}
	return b
}

func (a poly) add(b poly) poly {
	c := make(poly, len(a))
	for i, v := range a {
		c[i] = (v + b[i]) % P
	}
	return c
}

func (a poly) sub(b poly) poly {
	c := make(poly, len(a))
	for i, v := range a {
		c[i] = (v - b[i] + P) % P
	}
	return c
}

func (a poly) mul(k int) poly {
	k %= P
	b := make(poly, len(a))
	for i, v := range a {
		b[i] = v * k % P
	}
	return b
}

func (a poly) lsh(k int) poly {
	b := make(poly, len(a))
	if k > len(a) {
		return b
	}
	copy(b[k:], a)
	return b
}

func (a poly) rsh(k int) poly {
	b := make(poly, len(a))
	if k > len(a) {
		return b
	}
	copy(b, a[k:])
	return b
}

func (a poly) derivative() poly {
	n := len(a)
	d := make(poly, n)
	for i := 1; i < n; i++ {
		d[i-1] = a[i] * i % P
	}
	return d
}

func (a poly) integral() poly {
	n := len(a)
	s := make(poly, n)
	s[0] = 0 // C
	// 线性求逆元，详见 math.go 中的 initAllInv
	inv := make([]int, n)
	inv[1] = 1
	for i := 2; i < n; i++ {
		inv[i] = (P - P/i) * inv[P%i] % P
	}
	for i := 1; i < n; i++ {
		s[i] = a[i-1] * inv[i] % P
	}
	return s
}

// 多项式乘法逆 (mod x^n, 下同)
// 参考 https://blog.orzsiyuan.com/archives/Polynomial-Inversion/
// https://oi-wiki.org/math/poly/inv/
// 注：也可以用牛顿迭代法
// 模板题 https://www.luogu.com.cn/problem/P4238
func (a poly) inv() poly {
	n := len(a)
	m := 1 << bits.Len(uint(n))
	A := a.resize(m)
	invA := make(poly, m)
	invA[0] = nttPow(A[0], P-2)
	for l := 2; l <= m; l <<= 1 {
		ll := l << 1
		b := A[:l].resize(ll)
		iv := invA[:l].resize(ll)
		t := newNTT(ll)
		t.dft(b)
		t.dft(iv)
		for i, v := range iv {
			b[i] = v * (2 - v*b[i]%P + P) % P
		}
		t.idft(b)
		copy(invA, b[:l])
	}
	return invA[:n]
}

// 多项式除法
// https://blog.orzsiyuan.com/archives/Polynomial-Division-and-Modulo/
// https://oi-wiki.org/math/poly/div-mod/
// 模板题 https://www.luogu.com.cn/problem/P4512
func (a poly) div(b poly) poly {
	k := len(a) - len(b) + 1
	if k <= 0 {
		return make(poly, 1)
	}
	A := a.reverseCopy().resize(k)
	B := b.reverseCopy().resize(k)
	return A.conv(B.inv())[:k].reverse()
}

// 多项式取模
func (a poly) mod(b poly) poly {
	m := len(b)
	return a[:m-1].sub(a.div(b).conv(b)[:m-1])
}

func (a poly) divmod(b poly) (quo, rem poly) {
	m := len(b)
	quo = a.div(b)
	rem = a[:m-1].sub(quo.conv(b)[:m-1])
	return
}

// 多项式开根
// 参考 https://blog.orzsiyuan.com/archives/Polynomial-Square-Root/
// https://oi-wiki.org/math/poly/sqrt/
// 模板题 https://www.luogu.com.cn/problem/P5205
// 模板题（二次剩余）https://www.luogu.com.cn/problem/P5277
func (a poly) sqrt() poly {
	const inv2 = (P + 1) / 2
	n := len(a)
	m := 1 << bits.Len(uint(n))
	A := a.resize(m)
	rt := make(poly, m)
	rt[0] = 1
	if a[0] != 1 {
		rt[0] = int(new(big.Int).ModSqrt(big.NewInt(int64(a[0])), big.NewInt(P)).Int64())
		//if 2*rt[0] > P { // P5277 需要
		//	rt[0] = P - rt[0]
		//}
	}
	for l := 2; l <= m; l <<= 1 {
		ll := l << 1
		b := A[:l].resize(ll)
		r := rt[:l].resize(ll)
		ir := rt[:l].inv().resize(ll)
		t := newNTT(ll)
		t.dft(b)
		t.dft(r)
		t.dft(ir)
		for i, v := range r {
			b[i] = (b[i] + v*v%P) * inv2 % P * ir[i] % P
		}
		t.idft(b)
		copy(rt, b[:l])
	}
	return rt[:n]
}

// 多项式对数函数
// https://blog.orzsiyuan.com/archives/Polynomial-Natural-Logarithm/
// https://oi-wiki.org/math/poly/ln-exp/
// 模板题 https://www.luogu.com.cn/problem/P4725
func (a poly) ln() poly {
	if a[0] != 1 {
		panic(a[0])
	}
	return a.derivative().conv(a.inv())[:len(a)].integral()
}

// 多项式指数函数
// https://blog.orzsiyuan.com/archives/Polynomial-Exponential/
// https://oi-wiki.org/math/poly/ln-exp/
// 模板题 https://www.luogu.com.cn/problem/P4726
func (a poly) exp() poly {
	if a[0] != 0 {
		panic(a[0])
	}
	n := len(a)
	m := 1 << bits.Len(uint(n))
	A := a.resize(m)
	e := make(poly, m)
	e[0] = 1
	for l := 2; l <= m; l <<= 1 {
		b := e[:l].ln()
		b[0]--
		for i, v := range b {
			b[i] = (A[i] - v + P) % P
		}
		copy(e, b.conv(e[:l])[:l])
	}
	return e[:n]
}

// 多项式幂函数
// https://blog.orzsiyuan.com/archives/Polynomial-Power/
// https://oi-wiki.org/math/poly/ln-exp/#_5
// 模板题 https://www.luogu.com.cn/problem/P5245
// 模板题（a[0] != 1）https://www.luogu.com.cn/problem/P5273
func (a poly) pow(k int) poly {
	n := len(a)
	if k >= n && a[0] == 0 {
		return make(poly, n)
	}
	k1 := k % (P - 1)
	k %= P
	if a[0] == 1 {
		return a.ln().mul(k).exp()
	}
	shift := 0
	for ; shift < n && a[shift] == 0; shift++ {
	}
	if shift*k >= n {
		return make(poly, n)
	}
	a = a.rsh(shift)         // a[0] != 0
	a.mul(nttPow(a[0], P-2)) // a[0] == 1
	return a.ln().mul(k).exp().mul(nttPow(a[0], k1)).lsh(shift * k)
}

// 多项式三角函数
// 模意义下的单位根 i = w4 = g^((P-1)/4), 其中 g 为 P 的原根
// https://blog.orzsiyuan.com/archives/Polynomial-Trigonometric-Function/
// https://oi-wiki.org/math/poly/tri-func/
// 模板题 https://www.luogu.com.cn/problem/P5264
func (a poly) sincos() (sin, cos poly) {
	if a[0] != 0 {
		panic(a[0])
	}
	const i = 911660635    // pow(g, (P-1)/4)
	const inv2i = 43291859 // pow(2*i, P-2)
	const inv2 = (P + 1) / 2
	e := a.mul(i).exp()
	invE := e.inv()
	sin = e.sub(invE).mul(inv2i)
	cos = e.add(invE).mul(inv2)
	return
}

func (a poly) tan() poly {
	sin, cos := a.sincos()
	return sin.conv(cos.inv())
}

// 多项式反三角函数
// https://oi-wiki.org/math/poly/inv-tri-func/
// 模板题 https://www.luogu.com.cn/problem/P5265
func (a poly) asin() poly {
	if a[0] != 0 {
		panic(a[0])
	}
	n := len(a)
	b := a.conv(a)[:n].neg()
	b[0] = 1
	return a.derivative().conv(b.sqrt().inv())[:n].integral()
}

func (a poly) acos() poly {
	return a.asin().neg()
}

func (a poly) atan() poly {
	if a[0] != 0 {
		panic(a[0])
	}
	n := len(a)
	b := a.conv(a)[:n]
	b[0] = 1
	return a.derivative().conv(b.inv())[:n].integral()
}

// 多项式复合逆
// todo https://blog.csdn.net/weixin_43973966/article/details/88998646
//  模板题 https://www.luogu.com.cn/problem/P5809

// todo k 阶差分
// https://chatgpt.com/c/67afdb47-d464-8011-adc1-0f8639d3a546
