**点评**：本题没有用到高级数论知识，更多的是一些关于离线、枚举右维护左的技巧。解决本题可以加深你对 [HH 的项链](https://www.luogu.com.cn/problem/P1972) 这题的理解。

**前置题目**：[P1972 [SDOI2009] HH 的项链](https://www.luogu.com.cn/problem/P1972)。

**前置知识**：无平方因子核（core），见 [我的题解](https://leetcode.cn/problems/sum-of-perfect-square-ancestors/solution/ping-fang-sheng-yu-he-mei-ju-you-wei-hu-bfyxy/)。

$f(b)$ 等于把 $b$ 中两个 core 变成相等的最小操作次数。

比如其中两个 core 为 $x=2\cdot 3\cdot 5$，$y=3\cdot 7$，要想变成一样的，$x$ 中的 $2$ 和 $5$ 需要去掉（或者等价地把 $y$ 乘以 $2$ 再乘以 $5$），$y$ 中的 $7$ 需要去掉（或者等价地把 $x$ 乘以 $7$）。

所以这两个 $x$ 和 $y$ 变成一样的最小操作次数，等于 $x$ 和 $y$ 的质因子集合的对称差的大小。对称差就是不在交集中的元素。

枚举右，维护左。什么是左？考虑枚举 $x=\text{core}(a_i)$ 的质因子集合的子集 $S$ 作为交集，那么质因子集合包含 $S$ 的，就是我们要找的左。

在本题范围下，一个数的质因子个数 $\omega\le 7$，即使枚举质因子集合的所有子集，也只有 $2^7=128$ 个。

从左到右遍历 $a$，设 $x=\text{core}(a_i)$，枚举 $x$ 的质因子集合的子集 $S$，我们需要知道：

- 设左边的 $y=\text{core}(a_j)$，对于质因子集合包含 $S$ 的 $y$，下标 $j$ 最大是多少？

注意这里还需要按照 $y$ 的质因子集合的大小分别计算：

- 有的 $y$ 质因子集合小，但是 $j$ 的下标也小，超出询问范围。
- 有的 $y$ 质因子集合大，但是 $j$ 的下标也大，在询问范围中。

如何抉择？$\omega\le 7$，枚举。

定义 $\textit{maxI}_{m,\omega}$ 表示子集 $S$ 的元素积等于 $m$，且原集合的大小等于 $\omega$ 的最大下标。
例如 $x=\text{core}(a_i)$ 的 $\omega=3$，其中一个子集的元素积等于 $6$，那么记录 $\textit{maxI}_{6,3}=i$。

最后，按照右端点离线回答询问，如何找到答案？

技巧：保存操作次数分别为 $0,1,2,\ldots,14$ 时，最大的 $j$。那么从小到大枚举操作次数，第一个满足「$j$ 大于等于询问左端点」的操作次数就是答案。

定义 $\textit{opToI}_{\textit{op}}$ 表示当操作次数等于 $\textit{op}$ 时，最大的 $j$。

$\textit{op}$ 等于对称差的大小，这等于两个集合的大小，减去交集的大小乘以 $2$。交集的大小就是子集 $S$ 的大小（记作 $\textit{common}$）。

枚举「左」集合的大小 $\omega_2=\textit{common},\textit{common}+1,\ldots,7$，计算操作次数 $ \textit{op}=\omega+\omega_2-\textit{common}\cdot 2$，用 $\textit{maxI}_{m,\omega_2}$ 更新 $\textit{opToI}_{\textit{op}}$ 的最大值。

AC 代码（Golang）：

```go
package main
import("bufio";."fmt";"math/bits";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const mx = 5032108
	// 预处理每个数的最小质因子，用于质因子分解
	lpf := [mx]int{}
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	// 离线询问，按照右端点分组
	type query struct{ l, i int }
	qs := make([][]query, n)
	for i := range q {
		var l, r int
		Fscan(in, &l, &r)
		qs[r-1] = append(qs[r-1], query{l, i})
	}

	ans := make([]int, q)
	const mxW = 7
	opToI := make([]int, mxW*2+1)
	maxI := [mx][mxW + 1]int{}
	muls := [1 << mxW]int{1}

	for i, v := range a {
		// 计算 core(a[i]) 的所有质因子集合（计算元素积），作为交集
		mul := muls[:1]
		w := 0 // core(a[i]) 的质因子个数
		for v > 1 {
			p := lpf[v]
			e := 1
			for v /= p; v%p == 0; v /= p {
				e ^= 1
			}
			if e > 0 {
				w++
				for _, m := range mul {
					mul = append(mul, m*p)
				}
			}
		}

		// 枚举交集
		for mask, m := range mul {
			common := bits.OnesCount8(uint8(mask)) // 交集大小
			// 枚举左边的 core(a[j]) 的质因子个数 w2
			for w2 := common; w2 <= mxW; w2++ {
				op := w + w2 - common*2 // 操作次数等于对称差的大小
				// 维护当操作次数等于 op 时，最大的 j 的下标
				opToI[op] = max(opToI[op], maxI[m][w2])
			}
			maxI[m][w] = i + 1
		}

		for _, p := range qs[i] {
			for op, j := range opToI {
				if j >= p.l { // 左边的 a[j] 在询问范围内
					ans[p.i] = op
					break
				}
			}
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}
```

时间复杂度：$\mathcal{O}(U\log\log U + n\omega 2^{\omega}+q\omega)$。其中 $U=\max(a)\le 5032108$，$\omega\le 7$ 是 $a_i$ 的不同质因子个数。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
