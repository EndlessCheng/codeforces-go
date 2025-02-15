package copypasta

import (
	. "fmt"
	"io"
	"math"
	"slices"
)

// 3B1B 线性代数的本质 https://www.bilibili.com/video/BV1ys411472E

/* 矩阵加速
https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97#%E7%B7%9A%E6%80%A7%E4%BB%A3%E6%95%B8%E8%A7%A3%E6%B3%95
https://zhuanlan.zhihu.com/p/56444434
https://codeforces.com/blog/entry/80195 Matrix Exponentiation video + training contest
浅谈矩阵乘法在算法竞赛中的应用 https://zhuanlan.zhihu.com/p/631804105
F2 矩阵 有可能是可逆的，和或的 01 矩阵 似乎是肯定不可逆的，逆矩阵有时候也有一定的应用场景
除了直接的矩阵乘法，矩阵加法有时候也有用，有时候可以通过分块矩阵 或者逆矩阵 把连加表达成矩阵求幂
https://atcoder.jp/contests/abc299/tasks/abc299_h
这个开关灯问题 也涉及F2矩阵的逆矩阵（或高斯消元） https://github.com/tdzl2003/leetcode_live/blob/master/poj/1222_1753_3279.md
F2 矩阵 int64 to int64 的散列（可逆意味着一一映射，意味着无冲突） https://github.com/tdzl2003/leetcode_live/blob/master/other/int64_hash.md

三对角矩阵算法（托马斯算法）https://en.wikipedia.org/wiki/Tridiagonal_matrix_algorithm
https://codeforces.com/contest/24/problem/D

哈密尔顿–凯莱定理 Cayley–Hamilton theorem
特征多项式是零化多项式
https://en.wikipedia.org/wiki/Cayley%E2%80%93Hamilton_theorem

浅谈范德蒙德(Vandermonde)方阵的逆矩阵与拉格朗日(Lagrange)插值的关系以及快速傅里叶变换(FFT)中IDFT的原理 https://www.cnblogs.com/gzy-cjoier/p/9741950.html

模板题 https://www.luogu.com.cn/problem/P1939 https://ac.nowcoder.com/acm/contest/6357/A
https://codeforces.com/problemset/problem/1182/E
https://atcoder.jp/contests/abc232/tasks/abc232_e
有向图中长度为 k 的路径数 https://atcoder.jp/contests/dp/tasks/dp_r
TR 的数列 https://blog.csdn.net/zyz_bz/article/details/88993616
挑战 P202 一维方块染色 http://poj.org/problem?id=3734
3xM 的格子，其中有一些障碍物，求从第二行最左走到第二行最右的方案数，每次可以向右/右上/右下走一步 https://codeforces.com/problemset/problem/954/F
https://codeforces.com/problemset/problem/166/E

min max 矩阵快速幂
https://atcoder.jp/contests/abc236/tasks/abc236_g

& xor 矩阵快速幂
https://atcoder.jp/contests/abc009/tasks/abc009_4

todo poj 2345 3532 3526
*/

// 一些题目：https://oi-wiki.org/math/matrix/

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

/*
矩阵快速幂优化 DP
视频讲解：https://www.bilibili.com/video/BV1hn1MYhEtC/?t=21m27s
文字讲解：https://leetcode.cn/problems/student-attendance-record-ii/solutions/2885136/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-a8kj/
m 项递推式，以及包含常数项的情况见《挑战》P201
https://codeforces.com/problemset/problem/450/B 1300 也可以找规律
https://www.luogu.com.cn/problem/P10310
https://ac.nowcoder.com/acm/contest/9247/A
https://codeforces.com/problemset/problem/1117/D a(n) = a(n-1) + a(n-m)

https://www.luogu.com.cn/problem/P9777
已知 f(1) = x + 1/x = k，计算 f(n) = x^n + 1/x^n
由于 f(n) * f(1) = f(n+1) + f(n-1)
所以 f(n+1) = k*f(n) - f(n-1)，矩阵快速幂解决
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

// 一般是状态机 DP
// 操作 k 次
func solveDP(k int) (ans int) {
	const size = 26 // 第二维度的大小

	// DP 初始值（递归边界）
	// 一般是一个全为 1 的列向量，对应初始值 f[0][j]=1 或者递归边界 dfs(0,j)=1
	f0 := newMatrix(size, 1)
	for i := range f0 {
		f0[i][0] = 1
	}

	// 递推式中的 f[i][j] += f[i-1][k] * 2，提取系数得 m[j][k] = 2
	m := newMatrix(size, size)
	for i := range m {
		m[i][(i+1)%size] = 1 // 举例 f[i][j] = f[i][j+1] + f[i][j+2]
		m[i][(i+2)%size] = 1
	}

	// fk 和 f0 一样，都是长为 size 的列向量
	fk := m.powMul(k, f0)

	// 现在 fk[i][0] 就是 f[k][i] 或者 dfs(k,i)
	// 特别地，fk[0][0] 就是 f[k][0] 或者 dfs(k,0)
	for _, row := range fk {
		ans += row[0] // 举例 ans = sum(f[k])
	}
	ans %= mod

	return
}

// -----------------------------------------------------------------------------

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

// a(n) = p*a(n-1) + q*a(n-2)
// a(n-1) = a(n-1)
// 转成矩阵乘法
// 注意：数列从 0 开始，若题目从 1 开始则输入的 n 为 n-1
func calcFibonacci(p, q, a0, a1, n int) int {
	const mod = 1_000_000_007 // 998244353
	//n--
	if n == 0 {
		return (a0%mod + mod) % mod
	}
	if n == 1 {
		return (a1%mod + mod) % mod
	}
	m := matrix{
		{p, q},
		{1, 0},
	}.pow(n - 1)
	return ((m[0][0]*a1+m[0][1]*a0)%mod + mod) % mod
	//return m[0][0]
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
// 与 SCC 结合 https://www.luogu.com.cn/problem/P6030
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
				// 用当前行对其余行进行消元：从第 i 个式子中消去第 col 个未知数
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
// todo https://www.cnblogs.com/ywwyww/p/8522541.html
//  https://www.luogu.com.cn/problem/P7776
//  Berlekamp–Massey 算法 https://www.luogu.com.cn/problem/P5487

// 线性基（异或空间的极大线性无关子集）
// 可以用来解决「子序列异或和」相关问题
// https://oi-wiki.org/math/basis/
// https://en.wikipedia.org/wiki/Basis_(linear_algebra)
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
// 图上线性基
// https://www.luogu.com.cn/problem/P4151
// https://codeforces.com/problemset/problem/724/G 2600
// 类似思想 https://codeforces.com/problemset/problem/19/E 2900
//
// 模板题 https://loj.ac/p/113 https://www.luogu.com.cn/problem/P3812
// 题单 https://www.luogu.com.cn/training/11251
// https://codeforces.com/problemset/problem/959/F
// https://atcoder.jp/contests/abc141/tasks/abc141_f
// todo 构造 https://codeforces.com/problemset/problem/1427/E
//  https://codeforces.com/problemset/problem/1101/G
//  https://codeforces.com/problemset/problem/895/C   
//  - 加强版 https://loj.ac/p/2978
//  异或最短路/最长路 https://codeforces.com/problemset/problem/845/G https://www.luogu.com.cn/problem/P4151
//  https://www.luogu.com.cn/problem/P3857
type xorBasis struct {
	b []int // 核心就这一个

	num int
	or  int

	canBeZero bool  // 见 minXor 和 kthXor
	basis     []int // 见 initOnce

	rightMost     []int
	rightMostZero int
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
func (b *xorBasis) insert(v int) bool {
	b.or |= v
	// 从高到低遍历，方便计算下面的 maxXor 和 minXor
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 { // 线性无关
			b.b[i] = v
			b.num++
			return true
		}
		v ^= b.b[i]
	}
	b.canBeZero = true // 没有找到，但这说明了可以选一些数使得异或和为 0
	return false
}

// EXTRA: 如果遇到线性相关的基，保留位置最靠右的
// https://atcoder.jp/contests/abc223/tasks/abc223_h
// https://codeforces.com/problemset/problem/1902/F 2400
// https://codeforces.com/problemset/problem/1100/F 2500
// https://codeforces.com/problemset/problem/1778/E 2500
func (b *xorBasis) insertRightMost(idx, v int) bool {
	// 从高到低遍历，方便计算下面的 maxXor 和 minXor
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 { // 线性无关
			b.b[i] = v
			b.rightMost[i] = idx
			b.num++
			return true
		}
		if idx >= b.rightMost[i] { // 注意 b.rightMost[i] 的初始值为 0
			idx, b.rightMost[i] = b.rightMost[i], idx // 换个旧的 idx
			v, b.b[i] = b.b[i], v                     // 继续插入之前的基
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
		if v>>i&1 == 0 {
			continue
		}
		// b.b[i] == 0 || b.rightMost[i] < lowerIndex
		if b.b[i] == 0 {
			return false
		}
		v ^= b.b[i]
	}
	return true
}

// https://www.luogu.com.cn/problem/P3812
// https://loj.ac/p/113
func (b *xorBasis) maxXor() (xor int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if xor^b.b[i] > xor {
			xor ^= b.b[i]
		}
	}
	return
}

func (b *xorBasis) maxXorWithVal(val int) int {
	xor := val
	for i := len(b.b) - 1; i >= 0; i-- {
		if xor^b.b[i] > xor {
			xor ^= b.b[i]
		}
	}
	return xor
}

func (b *xorBasis) maxXorWithLowerIndex(lowerIndex int) (xor int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if xor>>i&1 == 0 && b.rightMost[i] >= lowerIndex && xor^b.b[i] > xor {
			xor ^= b.b[i]
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

func (b *xorBasis) initOnce() {
	if b.basis != nil {
		return
	}
	tmp := append([]int{}, b.b...)
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

// 线性基能表出的所有不同元素中的第 k 小值（不允许空）
// k 从 1 开始
// https://loj.ac/p/114 http://acm.hdu.edu.cn/showproblem.php?pid=3949
func (b *xorBasis) kthXor(k int) (xor int) {
	b.initOnce()
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
		x := other.b[i]
		if x > 0 {
			b.insert(x)
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
//
// todo https://uoj.ac/problem/179
//  https://codeforces.com/problemset/problem/1430/G https://codeforces.com/blog/entry/83614?#comment-709868
//  https://codeforces.com/problemset/problem/375/E
//  NOI08 志愿者招募 https://www.luogu.com.cn/problem/P3980
//       整数线性规划与全幺模矩阵 https://www.acwing.com/file_system/file/content/whole/index/content/2197334/
