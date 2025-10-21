package copypasta

import (
	. "fmt"
	"io"
	"math"
	"slices"
)

// 3B1B 线性代数的本质 https://www.bilibili.com/video/BV1ys411472E

/* 矩阵
https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97#%E7%B7%9A%E6%80%A7%E4%BB%A3%E6%95%B8%E8%A7%A3%E6%B3%95
https://zhuanlan.zhihu.com/p/56444434
https://codeforces.com/blog/entry/80195 Matrix Exponentiation video + training contest
浅谈矩阵乘法在算法竞赛中的应用 https://zhuanlan.zhihu.com/p/631804105
F2 矩阵 有可能是可逆的，和或的 01 矩阵 似乎是肯定不可逆的，逆矩阵有时候也有一定的应用场景
除了直接的矩阵乘法，矩阵加法有时候也有用，有时候可以通过分块矩阵 或者逆矩阵 把连加表达成矩阵求幂
https://atcoder.jp/contests/abc299/tasks/abc299_h
这个开关灯问题 也涉及F2矩阵的逆矩阵（或高斯消元） https://github.com/tdzl2003/leetcode_live/blob/master/poj/1222_1753_3279.md
F2 矩阵 int64 to int64 的散列（可逆意味着一一映射，意味着无冲突） https://github.com/tdzl2003/leetcode_live/blob/master/other/int64_hash.md

Advanced Matrix Multiplication Optimization on Modern Multi-Core Processors
https://salykova.github.io/gemm-cpu

三对角矩阵算法（托马斯算法）https://en.wikipedia.org/wiki/Tridiagonal_matrix_algorithm
https://codeforces.com/contest/24/problem/D

浅谈范德蒙德(Vandermonde)方阵的逆矩阵与拉格朗日(Lagrange)插值的关系以及快速傅里叶变换(FFT)中IDFT的原理 https://www.cnblogs.com/gzy-cjoier/p/9741950.html

矩阵快速幂优化 DP
视频讲解：https://www.bilibili.com/video/BV1hn1MYhEtC/?t=21m27s
文字讲解：https://leetcode.cn/problems/student-attendance-record-ii/solutions/2885136/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-a8kj/
https://codeforces.com/problemset/problem/450/B 1300 也可以找规律
https://codeforces.com/problemset/problem/166/E 1500
https://codeforces.com/problemset/problem/691/E 1900
https://codeforces.com/problemset/problem/954/F 2100 3xM 的格子，其中有一些障碍物，求从第二行最左走到第二行最右的方案数，每次可以向右/右上/右下走一步
https://codeforces.com/problemset/problem/1117/D 2100 a(n) = a(n-1) + a(n-m)
https://codeforces.com/problemset/problem/1182/E 2300
https://codeforces.com/problemset/problem/226/C 2400
- https://www.luogu.com.cn/problem/P1306
https://codeforces.com/problemset/problem/593/E 2400 分段
https://codeforces.com/problemset/problem/93/D 2500 前 n 项之和
https://codeforces.com/problemset/problem/60/E 2600
https://codeforces.com/problemset/problem/575/A 2700 分段 倍增
https://atcoder.jp/contests/abc232/tasks/abc232_e
https://atcoder.jp/contests/dp/tasks/dp_r 有向图中长为 k 的路径数
https://www.luogu.com.cn/problem/P1939 https://ac.nowcoder.com/acm/contest/6357/A
https://www.luogu.com.cn/problem/P3216 12345678910111213...n % m
https://www.luogu.com.cn/problem/P10310
https://ac.nowcoder.com/acm/contest/9247/A
https://blog.csdn.net/zyz_bz/article/details/88993616 TR 的数列
http://poj.org/problem?id=3734 挑战 P202 一维方块染色

https://www.luogu.com.cn/problem/P9777
已知 f(1) = x + 1/x = k，计算 f(n) = x^n + 1/x^n
由于 f(n) * f(1) = f(n+1) + f(n-1)
所以 f(n+1) = k*f(n) - f(n-1)，矩阵快速幂解决

min max 矩阵快速幂
https://atcoder.jp/contests/abc236/tasks/abc236_g

& xor 矩阵快速幂
https://atcoder.jp/contests/abc009/tasks/abc009_4

todo poj 2345 3532 3526
*/
type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix) powMul(n int, f0 matrix) matrix {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

// 有两种类型的矩阵快速幂优化 DP
// 一种是多维 DP / 状态机 DP，转移系数写成一个 size*size 的矩阵，见下面的 solveDP
// 另一种是线性 DP，转移系数写在第一行，其余行 m[i+1][i] = 1，见下面的 calcFibonacci 
// - 更快的算法见下面的 kitamasa
//
// 特别地，对于多维 DP，如果要求计算前 n 项之和（前缀和），我们可以在列向量末尾添加一个前缀和项 s[i]
// 递推式 s[i] = sum(f[i]) + s[i-1] = sum(M @ f[i-1]) + s[i-1]
// 对应系数矩阵，在最下面新增一行，最右边新增一列
// 最下面新增一行：M 矩阵每一列的和，再添加一个 1，即 [sum(M[i][0]), sum(M[i][1]), ..., sum(M[i][-1]), 1]
// 最右边新增一列：除了最下面的 1 以外，其余全为 0
func solveDP(N int) (ans int) {
	const size = 26

	// 系数矩阵
	// 例如，递推式中的 f[i][j] += f[i-1][k] * 2，提取系数得 m[j][k] = 2
	m := newMatrix(size, size)
	for j := range m {
		// 根据题目修改
		for k := 0; k < j; k++ {
			m[j][k] = 1
		}
	}

	// 初始值
	// 一般是全为 1 的列向量，对应 f[0][j]=1 或者递归边界 dfs(0,j)=1
	f0 := newMatrix(size, 1)
	for j := range f0 {
		f0[j][0] = 1
	}

	// 答案
	// 一般来说需要迭代 N-1 次或者 N 次
	// fn[j][0] 对应 f[-1][j]
	fn := m.powMul(N-1, f0)
	for _, row := range fn {
		ans += row[0]
	}
	ans %= mod
	return
}

// -----------------------------------------------------------------------------

// 广义斐波那契数列
// a(n) = p*a(n-1) + q*a(n-2)
// 矩阵为 [f[n], f[n-1]]^T = [[p, q], [1, 0]] * [f[n-1], f[n-2]]^T
// 如果有常系数，例如 a(n) = p*a(n-1) + q*a(n-2) + C
// 矩阵为 [f[n], f[n-1], C]^T = [[p, q, 1], [1, 0, 0], [0, 0, 1]] * [f[n-1], f[n-2], C]^T
// ！数列下标从 1 开始，n 从 1 开始
// https://www.luogu.com.cn/problem/P1349
// https://www.luogu.com.cn/problem/P1939
// https://www.luogu.com.cn/problem/P1306
func calcFibonacci(p, q, a1, a2, n int) int {
	if n == 1 {
		return a1 % mod
	}
	// 变形得到 [f[n], f[n-1]]^T = [[p, q], [1, 0]] * [f[n-1], f[n-2]]^T
	// 也可以用打家劫舍的状态机写法理解，其中 f[i][0] 表示 i 可选可不选，f[i][1] 表示 i 一定不能选
	// f[i][0] += p*f[i-1][0] 不选 i
	// f[i][0] += q*f[i-1][1] 选 i，那么 i-1 一定不能选
	// f[i][1] = f[i-1][0]
	// 提取系数得 m[0][0] = p，m[0][1] = q，m[1][0] = 1
	m := matrix{
		{p, q},
		{1, 0},
	}
	f2 := matrix{
		{a2},
		{a1},
	}
	// 结果是列向量 [f[n], f[n-1]]，取第一项
	fn := m.powMul(n-2, f2)
	return fn[0][0]
}

// 给定常系数齐次递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
// 以及初始值 f(i) = a[i] (0 <= i < k)
// 返回 f(n)，其中参数 n 从 0 开始
// 注意入参 a 和 coef 的顺序
// Kitamasa 算法：如果只求第 n 项，可以做到 O(k^2 log n) 或者 O(k log k log n)，其中 k 是线性递推式的阶数，也是 coef 的长度
// 注：Kitamasa 译为「北正」，碰巧谐音「倍增」
// 另见 math_ntt.go 的 Bostan-Mori 算法
// https://codeforces.com/blog/entry/88760
// https://codeforces.com/blog/entry/97627
// https://misawa.github.io/others/fast_kitamasa_method.html
//
// https://atcoder.jp/contests/tdpc/tasks/tdpc_fibonacci
// https://www.luogu.com.cn/problem/P5487
func kitamasa(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans%mod + mod) % mod }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * pow(coef[0], n)
	}

	// 比如 f(4) = 3*f(2) + 2*f(1) + f(0)
	// 或者说 f(n) = 3*f(n-2) + 2*f(n-3) + f(n-4)
	// 那么 f(8) = 3*f(6) + 2*f(5) + f(4)
	// 其中 f(5) = 3*f(3) + 2*f(2) + f(1)
	//           = 3*(用 f(2) f(1) f(0) 表出) + 2*f(2) + f(1)
	// f(6) 同理
	// 这样可以用 f(2) f(1) f(0)，也就是 a[2] a[1] a[0] 表出 f(8)
	mul := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod
			}
			// 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
			bk := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk*coef[i]) % mod
			}
			b[0] = bk * coef[0] % mod
		}
		return c
	}

	// 计算 resC，以表出 f(n) = recC[k-1] * a[k-1] + recC[k-2] * a[k-2] + ... + resC[0] + a[0]
	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = mul(c, resC)
		}
		c = mul(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod
	}
	return
}

// Berlekamp-Massey 算法
// 给定数列的前 m 项，返回符合这个数列的最短常系数齐次递推式的系数 coef（设其长度为 k）
// 当 n >= k 时，有递推式 f(n) = coef[0] * f(n-1) + coef[1] * f(n-2) + ... + coef[k-1] * f(n-k)  （注意 coef 的顺序）
// 时间复杂度 O(mL)，其中 m 是 a 的长度，L 是最终 coef 的长度
// 关键思路：利用过去的失败，修正现在的失败
// ！如果模数不是质数，需要用 exgcd 或者其他方法求逆元
// 注：一种理解角度是，基于汉克尔矩阵的在线高斯消元
// 注：用 Cayley-Hamilton 定理可以证明，对于矩阵快速幂优化 DP，求出前 2k 项，就能得到递推方程（k 是系数矩阵的边长）
//    另见 https://www.luogu.com/paste/ytpmeswf 第 88 条
// https://en.wikipedia.org/wiki/Berlekamp%E2%80%93Massey_algorithm
// https://oi-wiki.org/math/berlekamp-massey/
// https://codeforces.com/blog/entry/61306
//
// https://www.luogu.com.cn/problem/P5487 模板题
// https://www.luogu.com.cn/problem/P7820
// https://codeforces.com/problemset/problem/1511/F 2700
// https://codeforces.com/problemset/problem/506/E 3000
// https://leetcode.cn/problems/total-characters-in-string-after-transformations-ii/
func berlekampMassey(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod
		}
		if d == 0 {
			continue
		}

		// 首次算错，初始化 coef
		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - 1 - preI
		oldSz := len(coef)
		sz := bias + len(preC) + 1
		var oldCoef []int
		if sz > oldSz {
			oldCoef = slices.Clone(coef)
			coef = slices.Grow(coef, sz-oldSz)[:sz]
		}

		// 上一次算错告诉我们，preD = a[preI] - sum_j preC[j]*a[preI-1-j]
		// 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
		// 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
		// 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias=i-1-preI 处
		// 注意：preI 之前的数据符合旧公式，即 a[(<preI)] = sum_j preC[j]*a[(<preI)-1-j]
		//      对于新公式，i 之前的每个公式增加了 d/preD * (a[(<preI)] - sum_j preC[j]*a[(<preI)-1-j]) = d/preD * 0 = 0，所以也符合新公式
		delta := d * pow(preD, mod-2) % mod
		coef[bias] = (coef[bias] + delta) % mod
		for j, c := range preC {
			coef[bias+1+j] = (coef[bias+1+j] - delta*c) % mod
		}

		if sz > oldSz {
			preC = oldCoef
			preI, preD = i, d
		}
	}

	// 去掉不必要的 0
	for len(coef) > 0 && coef[len(coef)-1] == 0 {
		coef = coef[:len(coef)-1]
	}

	// 把负数调整为非负数
	// 比如后面计算递推式第 n 项，这可以保证不会产生负数（但那样的话，可以最后输出时再调整，所以下面的循环其实没必要）
	for i, c := range coef {
		coef[i] = (c + mod) % mod
	}

	return
}

// 已知数列的前 m 项，猜测一个符合最短线性递推式的第 n 项
// https://www.luogu.com.cn/problem/P5487
func guessNth(a []int, n int) int {
	coef := berlekampMassey(a)
	slices.Reverse(coef) // 注意 kitamasa 入参的顺序
	nth := kitamasa(coef, a, n)
	return nth
}

//

func newIdentityMatrix(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) pow(n int) matrix {
	res := newIdentityMatrix(len(a))
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

// 比如 n*n 的国际象棋的马，从 (sx,sy) 走 k 步到 (tx,ty)，需要多少步
// 这里可以先 O(n^2) 预处理走一步的转移，构建矩阵 a
// 然后用一个 [1 * (n^2)] 的矩阵初始矩阵乘 a^k
// 得到一个 [1 * (n^2)] 的结果矩阵 res
// res[0][tx*n+ty] 就是答案
func (a matrix) solve(n, sx, sy, tx, ty, k int) int {
	b := matrix{make([]int, n*n)}
	b[0][sx*n+sy] = 1
	res := b.mul(a.pow(k))
	return res[0][tx*n+ty]
}

//

func readMatrix(in io.Reader, n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
		//a[i] = make([]int, m, m+1) // 方便高斯消元
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	return a
}

func copyMatrix(a matrix) matrix {
	b := make(matrix, len(a))
	for i, row := range a {
		b[i] = slices.Clone(row)
	}
	return b
}

// 顺时针转 90°
func rotateMatrix(a matrix) matrix {
	b := make(matrix, len(a[0]))
	for j := range b {
		b[j] = make([]int, len(a))
		for i, row := range a {
			b[j][len(a)-1-i] = row[j]
		}
	}
	return b
}

func (a matrix) add(b matrix) matrix {
	c := newMatrix(len(a), len(a[0]))
	for i, row := range a {
		for j, aij := range row {
			c[i][j] = aij + b[i][j] // % mod
		}
	}
	return c
}

func (a matrix) sub(b matrix) matrix {
	c := newMatrix(len(a), len(a[0]))
	for i, row := range a {
		for j, aij := range row {
			c[i][j] = aij - b[i][j] // % mod) + mod) % mod
		}
	}
	return c
}

func (a matrix) swapRows(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a matrix) swapCols(i, j int) {
	for k := range a {
		a[k][i], a[k][j] = a[k][j], a[k][i]
	}
}

func (a matrix) mulRow(i, k int) {
	for j := range a[i] {
		a[i][j] *= k // % mod
	}
}

func (a matrix) trace() (sum int) {
	for i, row := range a {
		sum += row[i]
	}
	return
}

// NxN 矩阵求逆
// 模板题 https://www.luogu.com.cn/problem/P4783
func (matrix) inv(A matrix) matrix {
	// 增广一个单位矩阵
	n := len(A)
	m := 2 * n
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a {
			a[i][j] = A[i][j] // or read
		}
		a[i][n+i] = 1
	}

	for i := range a {
		for j := i; j < n; j++ {
			if a[j][i] != 0 {
				a[i], a[j] = a[j], a[i]
				break
			}
		}
		if a[i][i] == 0 {
			// 矩阵不是满秩的
			return nil
		}
		inv := pow(a[i][i], mod-2)
		for j := i; j < m; j++ {
			a[i][j] = a[i][j] * inv % mod
		}
		for j := range a {
			if j != i {
				inv := a[j][i]
				for k := i; k < m; k++ {
					a[j][k] = (a[j][k] - inv*a[i][k]%mod + mod) % mod
				}
			}
		}
	}

	// 结果保存在 a 右侧
	res := make(matrix, n)
	for i, row := range a {
		res[i] = row[n:]
	}
	return res
}

// 高斯消元 Gaussian elimination O(n^3)   列主元消去法
// 求解 Ax=B，A 为方阵，返回解（无解或有无穷多组解）
// https://en.wikipedia.org/wiki/Gaussian_elimination
// https://en.wikipedia.org/wiki/Pivot_element#Partial_and_complete_pivoting
// https://oi-wiki.org/math/gauss/
// 总结 https://cloud.tencent.com/developer/article/1087352
// https://cp-algorithms.com/linear_algebra/determinant-gauss.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GaussianElimination.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GaussJordanElimination.java.html
// 模板题 https://www.luogu.com.cn/problem/P3389 https://www.luogu.com.cn/problem/P2455
//       https://codeforces.com/problemset/problem/21/B
// https://www.luogu.com.cn/problem/P4035
// https://www.luogu.com.cn/problem/P6030 与 SCC 结合
//
// 三对角矩阵算法 托马斯算法
// https://en.wikipedia.org/wiki/Tridiagonal_matrix_algorithm
// https://codeforces.com/problemset/problem/24/D 2400
func gaussJordanElimination(A matrix, B []int) (sol []float64, infSol bool) {
	const eps = 1e-8
	n := len(A)
	// 构造增广矩阵 (or read)
	a := make([][]float64, n)
	for i, row := range A {
		a[i] = make([]float64, n+1)
		for j, v := range row {
			a[i][j] = float64(v)
		}
		a[i][n] = float64(B[i])
	}
	row := 0
	for col := 0; col < n; col++ {
		// 列主元消去法：减小误差，把正在处理的未知数系数的绝对值最大的式子换到第 row 行
		pivot := row
		for i := row; i < n; i++ {
			if math.Abs(a[i][col]) > math.Abs(a[pivot][col]) {
				pivot = i
			}
		}
		// 这一列全为 0，表明无解或有无穷多解，具体是哪一种需要消元完成后才知道
		if math.Abs(a[pivot][col]) < eps {
			continue
		}
		a[row], a[pivot] = a[pivot], a[row]
		// 初等行变换：把正在处理的未知数的系数变为 1
		for j := col + 1; j <= n; j++ {
			a[row][j] /= a[row][col]
		}
		// 消元，构造简化行梯阵式
		for i := range a {
			if i != row {
				// 用当前行对其余行消元：从第 i 个式子中消去第 col 个未知数
				for j := col + 1; j <= n; j++ {
					a[i][j] -= a[i][col] * a[row][j]
				}
			}
		}
		row++
	}
	if row < n {
		for _, r := range a[row:] {
			if math.Abs(r[n]) > eps {
				return nil, false
			}
		}
		return nil, true
	}
	res := make([]float64, n)
	for i, r := range a {
		res[i] = r[n]
	}
	return res, false
}

// EXTRA: 求行列式（对结果模 mod）
// https://en.wikipedia.org/wiki/Determinant
// 参考 https://www.luogu.com.cn/blog/Stormy-Rey/calculate-det
// https://www.luogu.com.cn/problem/P7112
func (a matrix) determinant(mod int) int {
	n := len(a)
	res, sign := 1, 1
	for i := range a {
		for j := i + 1; j < n; j++ {
			for a[i][i] != 0 {
				div := a[j][i] / a[i][i]
				for k := i; k < n; k++ {
					a[j][k] = (a[j][k] - a[i][k]*div%mod + mod) % mod
				}
				a[i], a[j], sign = a[j], a[i], -sign
			}
			a[i], a[j], sign = a[j], a[i], -sign
		}
	}
	for i, r := range a {
		res = res * r[i] % mod
	}
	res = (res*sign + mod) % mod
	return res
}

// 求矩阵的特征多项式
// 哈密尔顿–凯莱定理 Cayley–Hamilton theorem
// 特征多项式是零化多项式
// https://en.wikipedia.org/wiki/Cayley%E2%80%93Hamilton_theorem
// todo https://www.cnblogs.com/ywwyww/p/8522541.html
//  https://www.luogu.com.cn/problem/P7776
//  https://www.luogu.com.cn/problem/P10775

// 线性基（异或空间的极大线性无关子集）
// 可以用来解决「子序列异或和」相关问题
// https://oi-wiki.org/math/basis/
// https://en.wikipedia.org/wiki/Basis_(linear_algebra)
// https://usaco.guide/adv/xor-basis
// 【推荐】https://www.luogu.com.cn/blog/Marser/solution-p3812
// 线性基学习笔记 https://oi.men.ci/linear-basis-notes/
// XOR basis without linear algebra https://codeforces.com/blog/entry/100066
// https://www.luogu.com.cn/blog/i207M/xian-xing-ji-xue-xi-bi-ji-xie-ti-bao-gao
// 讲解+题单 https://www.cnblogs.com/UntitledCpp/p/13912602.html
// https://www.luogu.com.cn/blog/Troverld/xian-xing-ji-xue-xi-bi-ji
// todo 讲到了线性基的删除操作 https://blog.csdn.net/a_forever_dream/article/details/83654397
// 线性基求交 https://www.cnblogs.com/BakaCirno/p/11298102.html
// https://zhuanlan.zhihu.com/p/139074556
//
// 模板题 https://www.luogu.com.cn/problem/P3812 https://loj.ac/p/113
// 题单 https://www.luogu.com.cn/training/11251
// https://codeforces.com/problemset/problem/845/G 2300 异或最短路/最长路
// - https://www.luogu.com.cn/problem/P4151
// https://codeforces.com/problemset/problem/1101/G 2300
// https://codeforces.com/problemset/problem/662/A 2400 博弈
// https://codeforces.com/problemset/problem/959/F 2400
// https://codeforces.com/problemset/problem/1163/E 2400
// https://codeforces.com/problemset/problem/1902/F 2400 LCA
// - https://www.luogu.com.cn/problem/P3292 [SCOI2016] 幸运数字
// https://codeforces.com/problemset/problem/1100/F 2500
// https://codeforces.com/problemset/problem/1427/E 2500 构造
// https://codeforces.com/problemset/problem/1778/E 2500
// https://codeforces.com/problemset/problem/724/G 2600 图上线性基
// https://codeforces.com/problemset/problem/251/D 2700 输出具体方案
// - https://atcoder.jp/contests/abc141/tasks/abc141_f 简单版本
// https://codeforces.com/problemset/problem/19/E 2900 图上线性基
// https://codeforces.com/problemset/problem/587/E 2900
// - https://www.luogu.com.cn/problem/P5607
// https://codeforces.com/problemset/problem/1299/D 3000
// https://codeforces.com/problemset/problem/1336/E2 3500
// https://codeforces.com/gym/102331/problem/E
// https://atcoder.jp/contests/agc045/tasks/agc045_a
// https://atcoder.jp/contests/cf16-exhibition-final/tasks/cf16_exhibition_final_h
// https://www.luogu.com.cn/problem/P3857
// https://loj.ac/p/2978
// - https://codeforces.com/problemset/problem/895/C
type xorBasis struct {
	b []int // 核心就这一个

	rightMost     []int
	rightMostZero int

	num int
	or  int

	canBeZero bool  // 见 minXor 和 kthXor
	basis     []int // 见 initOnce
}

func newXorBasis(a []int) *xorBasis {
	b := &xorBasis{b: make([]int, 64)}  // or 32
	b.rightMost = make([]int, len(b.b)) // 注意这里是 0
	b.rightMostZero = -1                // 注意这里是 -1
	for _, v := range a {
		b.insert(v)
	}
	return b
}

// 尝试插入 v，看能否找到一个新的线性无关基
// 针对稀疏二进制的写法 https://leetcode.cn/problems/partition-array-for-maximum-xor-and-and/solution/shi-zi-bian-xing-xian-xing-ji-pythonjava-3e80/
func (b *xorBasis) insert(v int) bool {
	b.or |= v
	// 从高到低遍历，保证计算 maxXor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i == 0 { // 由于大于 i 的位都被我们异或成了 0，所以 v>>i 的结果只能是 0 或 1
			continue
		}
		if b.b[i] == 0 { // x 和之前的基是线性无关的
			b.b[i] = v // 新增一个基，最高位为 i
			b.num++
			return true
		}
		v ^= b.b[i] // 保证每个基的二进制长度互不相同
	}
	// 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
	b.canBeZero = true // 说明存在非空集合，异或和为 0
	return false
}

// EXTRA: 从高到低，对于二进制长度相同的基，选更靠右的
// https://atcoder.jp/contests/abc223/tasks/abc223_h
// https://codeforces.com/problemset/problem/1902/F 2400
// https://codeforces.com/problemset/problem/1100/F 2500
// https://codeforces.com/problemset/problem/1778/E 2500
func (b *xorBasis) insertRightMost(idx, v int) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.rightMost[i] = idx
			b.num++
			return true
		}
		// 替换掉之前的基，尽量保证基的下标都是最新的
		// 替换后，可能插入新的基，也可能淘汰掉旧的基
		if idx > b.rightMost[i] {
			idx, b.rightMost[i] = b.rightMost[i], idx
			v, b.b[i] = b.b[i], v
		}
		v ^= b.b[i]
	}
	b.canBeZero = true // 没有找到，但这说明了可以选一些数使得异或和为 0
	b.rightMostZero = max(b.rightMostZero, idx)
	return false
}

// v 能否被线性基表出
func (b *xorBasis) decompose(v int) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i == 0 {
			continue
		}
		// 配合 insertRightMost
		// b.b[i] == 0 || b.rightMost[i] < lowerIndex
		if b.b[i] == 0 {
			return false
		}
		v ^= b.b[i]
	}
	return true
}

// 返回能被线性基表出的最大值
// 如果线性基为空，返回 0
// https://www.luogu.com.cn/problem/P3812
// https://loj.ac/p/113
// https://leetcode.cn/problems/partition-array-for-maximum-xor-and-and/solutions/3734850/shi-zi-bian-xing-xian-xing-ji-pythonjava-3e80/
func (b *xorBasis) maxXor() (res int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		res = max(res, res^b.b[i])
	}
	return
}

func (b *xorBasis) maxXorWithVal(val int) int {
	res := val
	for i := len(b.b) - 1; i >= 0; i-- {
		res = max(res, res^b.b[i])
	}
	return res
}

// 配合 insertRightMost
func (b *xorBasis) maxXorWithLowerIndex(lowerIndex int) (res int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if res>>i&1 == 0 && b.rightMost[i] >= lowerIndex {
			res = max(res, res^b.b[i])
		}
	}
	return
}

// 考虑插入的过程，因为每一次跳转操作，x 的二进制最高位必定单调降低，所以不可能插入两个二进制最高位相同的数。
// 而此时，线性基中最小值异或上其他数，必定会增大。
// 所以，直接输出线性基中的最小值即可。
func (b *xorBasis) minXor() int {
	if b.canBeZero {
		return 0
	}
	for i := 0; ; i++ {
		if b.b[i] > 0 {
			return b.b[i]
		}
	}
}

// 保证每列只有一个 1
func (b *xorBasis) initOnce() {
	if b.basis != nil {
		return
	}
	tmp := slices.Clone(b.b)
	for i := range tmp {
		if tmp[i] == 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if tmp[i]>>j&1 > 0 {
				tmp[i] ^= tmp[j]
			}
		}
		b.basis = append(b.basis, tmp[i])
	}
}

// 线性基能表出的所有不同元素中的第 k 小值（不允许空集）
// k 从 1 开始
// https://loj.ac/p/114 http://acm.hdu.edu.cn/showproblem.php?pid=3949
func (b *xorBasis) kthXor(k int) (xor int) {
	b.initOnce()     // 只会初始化一次
	if b.canBeZero { // 0 是最小的
		k-- // 占用了一个数
	}
	if k >= 1<<len(b.basis) { // 非空子集有 2^len(b.basis) - 1 个
		return -1
	}
	for i, v := range b.basis {
		if k>>i&1 > 0 {
			xor ^= v
		}
	}
	return
}

// todo https://www.luogu.com.cn/problem/P4869
func (b *xorBasis) rank(xor int) (k int) {
	panic("todo")
}

// https://codeforces.com/problemset/problem/1902/F
func (b *xorBasis) merge(other *xorBasis) {
	for i := len(other.b) - 1; i >= 0; i-- {
		v := other.b[i]
		if v > 0 {
			b.insert(v)
		}
	}
}

/* 矩阵树定理 基尔霍夫定理 Kirchhoff‘s theorem
https://oi-wiki.org/graph/matrix-tree/
https://en.wikipedia.org/wiki/Kirchhoff%27s_theorem

https://atcoder.jp/contests/jsc2021/tasks/jsc2021_g
https://atcoder.jp/contests/abc253/tasks/abc253_h
https://atcoder.jp/contests/abc323/tasks/abc323_g
*/

// 线性规划（单纯形法）  LP, linear programming (simplex method)
// https://en.wikipedia.org/wiki/Mathematical_optimization
// https://en.wikipedia.org/wiki/Linear_programming
// https://en.wikipedia.org/wiki/Integer_programming
// https://en.wikipedia.org/wiki/Simplex_algorithm
// todo https://oi-wiki.org/math/simplex/
//      https://zhuanlan.zhihu.com/p/31644892
//  https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/LinearProgramming.java.html
//  https://web.mit.edu/lpsolve/doc/Python.htm
//
// todo https://uoj.ac/problem/179
//  https://codeforces.com/problemset/problem/1430/G https://codeforces.com/blog/entry/83614?#comment-709868
//  https://codeforces.com/problemset/problem/375/E
//  NOI08 志愿者招募 https://www.luogu.com.cn/problem/P3980
//       整数线性规划与全幺模矩阵 https://www.acwing.com/file_system/file/content/whole/index/content/2197334/
