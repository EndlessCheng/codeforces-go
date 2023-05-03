package copypasta

import (
	. "fmt"
	"io"
	"math"
)

// 3B1B 线性代数的本质 https://www.bilibili.com/video/BV1ys411472E

/* 矩阵加速
https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97#%E7%B7%9A%E6%80%A7%E4%BB%A3%E6%95%B8%E8%A7%A3%E6%B3%95
https://zhuanlan.zhihu.com/p/56444434
https://codeforces.com/blog/entry/80195 Matrix Exponentiation video + training contest

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

todo poj 2345 3532 3526
*/

// 一些题目：https://oi-wiki.org/math/matrix/

func readMatrix(in io.Reader, n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, m)
		//a[i] = make([]int64, m, m+1) // 方便高斯消元
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	return a
}

func copyMatrix(a matrix) matrix {
	b := make(matrix, len(a))
	for i, row := range a {
		b[i] = append([]int64(nil), row...)
	}
	return b
}

// 顺时针转 90°
func rotateMatrix(a matrix) matrix {
	b := make(matrix, len(a[0]))
	for j := range b {
		b[j] = make([]int64, len(a))
		for i, row := range a {
			b[j][len(a)-1-i] = row[j]
		}
	}
	return b
}

// 矩阵快速幂

type matrix [][]int64

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, m)
	}
	return a
}

func newIdentityMatrix(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	const mod int64 = 1e9 + 7 // 998244353
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod // 注：此处不能化简
			}
			if c[i][j] < 0 {
				c[i][j] += mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int64) matrix {
	res := newIdentityMatrix(len(a))
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
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
func (a matrix) solve(n, sx, sy, tx, ty int, k int64) int64 {
	b := matrix{make([]int64, n*n)}
	b[0][sx*n+sy] = 1
	res := b.mul(a.pow(k))
	return res[0][tx*n+ty]
}

// a(n) = p*a(n-1) + q*a(n-2)
// 注意：数列从 0 开始，若题目从 1 开始则输入的 n 为 n-1
// https://ac.nowcoder.com/acm/contest/9247/A
// m 项递推式，以及包含常数项的情况见《挑战》P201
// a(n) = a(n-1) + a(n-m) https://codeforces.com/problemset/problem/1117/D
func calcFibonacci(p, q, a0, a1, n int64) int64 {
	const mod int64 = 1e9 + 7 // 998244353
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

//

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

func (a matrix) mulRow(i int, k int64) {
	for j := range a[i] {
		a[i][j] *= k // % mod
	}
}

func (a matrix) trace() (sum int64) {
	for i, row := range a {
		sum += row[i]
	}
	return
}

// NxN 矩阵求逆
// 模板题 https://www.luogu.com.cn/problem/P4783
func (matrix) inv(A matrix) matrix {
	const mod int64 = 1e9 + 7
	pow := func(x int64) (res int64) {
		//x %= mod
		res = 1
		for n := mod - 2; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	// 增广一个单位矩阵
	n := len(A)
	m := 2 * n
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, m)
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
		inv := pow(a[i][i])
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
func gaussJordanElimination(A matrix, B []int64) (sol []float64, infSol bool) {
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
func (a matrix) determinant(mod int64) int64 {
	n := len(a)
	res, sign := int64(1), 1
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
	res = (res*int64(sign) + mod) % mod
	return res
}

// 求矩阵的特征多项式
// todo https://www.cnblogs.com/ywwyww/p/8522541.html

// 线性基（异或空间的极大线性无关子集）
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
// 模板题 https://loj.ac/p/113 https://www.luogu.com.cn/problem/P3812
// 题单 https://www.luogu.com.cn/training/11251
// todo 构造 https://codeforces.com/problemset/problem/1427/E
//  https://codeforces.com/problemset/problem/1101/G
//  https://codeforces.com/problemset/problem/895/C
//  异或最短路/最长路 https://codeforces.com/problemset/problem/845/G https://www.luogu.com.cn/problem/P4151
//  https://www.luogu.com.cn/problem/P3857
//  最右线性基 https://codeforces.com/problemset/problem/1778/E
type xorBasis struct {
	b   []int64
	num uint8

	canBeZero bool
	basis     []int64

	rightMost []int
}

func newXorBasis(a []int64) *xorBasis {
	b := &xorBasis{b: make([]int64, 64)} // 32
	b.rightMost = make([]int, len(b.b))
	for _, v := range a {
		b.insert(v)
	}
	return b
}

// 尝试插入 v，看能否找到一个新的线性无关基
func (b *xorBasis) insert(v int64) {
	// 从高到低遍历，方便计算下面的 maxXor 和 minXor
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 { // 线性无关
			b.b[i] = v
			b.num++
			return
		}
		v ^= b.b[i]
	}
	b.canBeZero = true // 没有找到，但这说明了可以选一些数使得异或和为 0
}

// EXTRA: 如果遇到线性相关的基，保留位置最靠右的
// https://codeforces.com/problemset/problem/1778/E
func (b *xorBasis) insertRightMost(idx int, v int64) {
	// 从高到低遍历，方便计算下面的 maxXor 和 minXor
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 { // 线性无关
			b.b[i] = v
			b.rightMost[i] = idx
			b.num++
			return
		}
		if idx >= b.rightMost[i] { // 注意 b.rightMost[i] 的初始值为 0
			idx, b.rightMost[i] = b.rightMost[i], idx
			v, b.b[i] = b.b[i], v // 继续插入之前的基
		}
		v ^= b.b[i]
	}
	b.canBeZero = true // 没有找到，但这说明了可以选一些数使得异或和为 0
}

// v 能否被线性基表出
func (b *xorBasis) decompose(v int64) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 > 0 {
			if b.b[i] == 0 {
				return false
			}
			v ^= b.b[i]
		}
	}
	return true
}

// https://www.luogu.com.cn/problem/P3812 https://loj.ac/p/113
func (b *xorBasis) maxXor() (xor int64) {
	for i := len(b.b) - 1; i >= 0; i-- {
		//if xor>>i&1 > 0 {
		//	continue
		//}
		if xor^b.b[i] > xor {
			xor ^= b.b[i]
		}
	}
	return
}

// 考虑插入的过程，因为每一次跳转操作，x 的二进制最高位必定单调降低，所以不可能插入两个二进制最高位相同的数。
// 而此时，线性基中最小值异或上其他数，必定会增大。
// 所以，直接输出线性基中的最小值即可。
func (b *xorBasis) minXor() int64 {
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
	tmp := append([]int64{}, b.b...)
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
func (b *xorBasis) kthXor(k int64) (xor int64) {
	b.initOnce()
	if b.canBeZero { // 0 是最小的
		k-- // 占用了一个数
	}
	if k >= int64(1)<<len(b.basis) { // 非空子集有 2^len(b.basis) - 1 个
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
func (b *xorBasis) rank(xor int64) (k int64) {
	panic("todo")
}

func (b *xorBasis) merge(other *xorBasis) {
	for i := len(other.b) - 1; i >= 0; i-- {
		x := other.b[i]
		if x > 0 {
			b.insert(x)
		}
	}
}

// 矩阵树定理 基尔霍夫定理 Kirchhoff‘s theorem
// https://oi-wiki.org/graph/matrix-tree/
// https://en.wikipedia.org/wiki/Kirchhoff%27s_theorem

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
