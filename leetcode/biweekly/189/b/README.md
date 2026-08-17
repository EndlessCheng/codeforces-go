## 方法一：暴力枚举

枚举 $s$ 左旋了 $\textit{rot} = 0,1,2,\ldots,n-1$ 次。

$s$ 左旋 $\textit{rot}$ 次后的字符串，是 $s+s$ 的一个子串，左端点下标为 $\textit{rot}$，右端点下标为 $\textit{rot}+n-1$。

我们需要判断该子串能否通过递增操作变成回文串，即 $s+s$ 下标 $\textit{rot}+i$ 的字母要等于下标 $\textit{rot}+n-1-i$ 的字母。

对于字母 $x$ 和 $y$（$x\le y$），我们可以：

- 把 $x$ 增大到 $y$，操作 $y-x$ 次。
- 把 $y$ 增大绕回到 $x$，操作 $26-y+x$ 次。

两种情况取最小值。

> **注**：把 $x$ 和 $y$ 都操作一定不优：把 $x$ 和 $y$ 都少操作一次，两个字母最终仍然相同。

[本题视频讲解](https://www.bilibili.com/video/BV1Q9bD6gEi1/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str) -> int:
        n = len(s)
        ans = inf
        for rot in range(n):
            op = rot
            for i in range(n // 2):
                d = abs(ord(s[(rot + i) % n]) - ord(s[(rot - 1 - i) % n]))
                op += min(d, 26 - d)  # 注：这里可以加个剪枝，如果 op >= ans 则 break
            ans = min(ans, op)
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = Integer.MAX_VALUE;
        for (int rot = 0; rot < n; rot++) {
            int op = rot;
            for (int i = 0; i < n / 2; i++) {
                int d = Math.abs(s[(rot + i) % n] - s[(rot + n - 1 - i) % n]);
                op += Math.min(d, 26 - d); // 注：这里可以加个剪枝，如果 op >= ans 则 break
            }
            ans = Math.min(ans, op);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s) {
        int n = s.size();
        int ans = INT_MAX;
        for (int rot = 0; rot < n; rot++) {
            int op = rot;
            for (int i = 0; i < n / 2; i++) {
                int d = abs(s[(rot + i) % n] - s[(rot + n - 1 - i) % n]);
                op += min(d, 26 - d); // 注：这里可以加个剪枝，如果 op >= ans 则 break
            }
            ans = min(ans, op);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(s string) int {
	n := len(s)
	ans := math.MaxInt
	for rot := range n {
		op := rot
		for i := range n / 2 {
			d := abs(int(s[(rot+i)%n]) - int(s[(rot+n-1-i)%n]))
			op += min(d, 26-d) // 注：这里可以加个剪枝，如果 op >= ans 则 break
		}
		ans = min(ans, op)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：循环自卷积

定义 $D(x,y)$ 为使字母 $x$ 和 $y$ 相等的最少操作次数。

设左旋了 $R$ 次。在方法一中，我们计算了递增操作的总次数

$$
\begin{aligned}
S_R &= \sum_{i=0}^{\lfloor n/2 \rfloor - 1} D(s[(R+i)\bmod n],s[(R+n-1-i)\bmod n])         \\
    &= \dfrac{1}{2} \sum_{i=0}^{n-1} D(s[(R+i)\bmod n],s[(R+n-1-i)\bmod n])           \\
\end{aligned}
$$

如何更快地计算上式？

考虑两个下标之和模 $n$ 的余数

$$
(R+i) + (R+n-1-i) \equiv 2R-1 \pmod n
$$

当 $R$ 固定时，$2R-1$ 是一个**定值**。

回想一下卷积的式子，如果我们能把 $S_R$ 表示成某种**循环卷积**，就能更快地计算答案。

> **注**：给定长为 $n$ 的数组 $a$，那么 $a$ 的循环自卷积为
> 
> $$
> \textit{c}[r] = \sum_{\substack{0\le i,j< n\\ i+j\equiv r\pmod n}} a[i]\cdot a[j]
> $$
> 
> 如果 $a$ 中只有 $0$ 和 $1$，那么计算循环自卷积可以快速回答例如「$a$ 左旋若干次后，有多少个镜像位置不同」的问题。

把字母视作一个在 $[0,25]$ 中的整数。本质上说，$D(x,y)$ 是在长为 $26$ 的环上，点 $x$ 和点 $y$ 的最短距离。

例如字母 $\texttt{b}$ 和 $\texttt{e}$，对应整数 $1$ 和 $4$，最短距离为 $3$。

想象把这个环均匀切开，得到两个半圆弧（半圆弧的端点在整点上）。在什么情况下，$x=1$ 和 $y=4$ 属于不同的半圆弧？

- 如果其中一个半圆孤是 $[2,14]$，那么 $x$ 和 $y$ 就分开了。
- 如果其中一个半圆孤是 $[3,15]$，那么 $x$ 和 $y$ 就分开了。
- 如果其中一个半圆孤是 $[4,16]$，那么 $x$ 和 $y$ 就分开了。
- 其余情况下，$x$ 和 $y$ 属于同一个半圆弧。

所以本质有 $3$ 种切法，可以把 $x=1$ 和 $y=4$ 分开，这恰好是 $x=1$ 和 $y=4$ 的最短距离。

一般地，定义特征函数

$$
I_k(x) =
\begin{cases} 
1, & x\in [k,k+12]     \\
0, & x\notin [k,k+12]     \\
\end{cases}
$$

整数 $x$ 和 $y$ 最短距离，等于有多少个不同的 $k$ 使得 $I_k(x)\ne I_k(y)$，即

$$
D(x,y) = \sum_{k=0}^{12} [I_k(x)\ne I_k(y)]
$$

> 注：记号 $[p\ne q]$ 表示当 $p\ne q$ 时结果为 $1$，否则为 $0$。相当于把 $\texttt{bool}$ 值转成 $\texttt{int}$ 值。对于本题来说，当 $p=0,q=1$ 或者 $p=1,q=0$ 时，记号 $[p\ne q]$ 才是 $1$。

设数组 $a_k = [I_k(s[0]), I_k(s[1]), \ldots, I_k(s[n-1])]$。

代入前文的和式，交换求和顺序，得

$$
\begin{aligned}
S_R &= \dfrac{1}{2} \sum_{i=0}^{n-1} \sum_{k=0}^{12} [a_k[(R+i)\bmod n]\ne a_k[(R+n-1-i)\bmod n)]]         \\
    &= \dfrac{1}{2} \sum_{k=0}^{12} \sum_{i=0}^{n-1} [a_k[(R+i)\bmod n]\ne a_k[(R+n-1-i)\bmod n)]]         \\
\end{aligned}
$$

举个例子，当 $R=0$ 时，内层的和式为

$$
[a_k[0]\ne a_k[n-1]] + [a_k[1]\ne a_k[n-2]] + \cdots + [a_k[n-1]\ne a_k[0]]
$$

设 $a_k$ 的循环自卷积为 $c_k$，其中

$$
\textit{c}_k[r] = \sum_{\substack{0\le i,j< n\\ i+j\equiv r\pmod n}} a_k[i]\cdot a_k[j]
$$

令 $r = (2R-1)\bmod n = -1 \bmod n = n-1$，我们有

$$
\textit{c}_k[n-1] = a_k[0]\cdot a_k[n-1] + a_k[1]\cdot a_k[n-2] + \cdots + a_k[n-1]\cdot a[0]
$$

根据定义，如果 $a_k[i] = a_k[j] = 1$，那么 $a_k[i]\cdot a_k[j] = 1$，否则 $a_k[i]\cdot a_k[j] = 0$。 

设 $a_k$ 中有 $\textit{cnt}_k$ 个 $1$。

$a_k[i]\ne a_k[n-1-i]$，有两种情况：

- $a_k[i] = 1$ 且 $a_k[n-1-i]=0$。用 $a_k[i] = 1$ 的个数（$\textit{cnt}_k$）减去 $a_k[i] = 1$ 且 $a_k[n-1-i]=1$ 的个数（$\textit{c}_k[n-1]$），就是 $a_k[i] = 1$ 且 $a_k[n-1-i]=0$ 的个数，即 $\textit{cnt}_k - \textit{c}_k[n-1]$。
- $a_k[i] = 0$ 且 $a_k[n-1-i]=1$。根据对称性，个数同样为 $\textit{cnt}_k - \textit{c}_k[n-1]$。

所以有

$$
\begin{aligned}
    & [a_k[0]\ne a_k[n-1]] + [a_k[1]\ne a_k[n-2]] + \cdots + [a_k[n-1]\ne a_k[0]]      \\
={} & 2(\textit{cnt}_k - \textit{c}_k[n-1])        \\
\end{aligned}
$$

一般地，令 $r = (2R-1)\bmod n$，那么有

$$
\begin{aligned}
S_R &= \dfrac{1}{2} \sum_{k=0}^{12} 2(\textit{cnt}_k - \textit{c}_k[r])         \\
    &= \sum_{k=0}^{12} \textit{cnt}_k - \sum_{k=0}^{12}\textit{c}_k[r]           \\
\end{aligned}
$$

其中和式 $\sum\limits_{k=0}^{12} \textit{cnt}_k$ 即 $a_0,a_1,\ldots,a_{12}$ 中的 $1$ 的总个数，可以用一个变量 $\textit{total}$ 表示。

设 $\textit{convSum}[r] = \sum\limits_{k=0}^{12}\textit{c}_k[r]$，那么有

$$
S_R = \textit{total} - \textit{convSum}[(2R-1)\bmod n]
$$

枚举 $R=0,1,2,\ldots,n-1$，答案为

$$
\begin{aligned}
    & \min_{R=0}^{n-1} R + S_R      \\
={} & \textit{total} + \min_{R=0}^{n-1} R - \textit{convSum}[(2R-1)\bmod n]        \\
\end{aligned}
$$

```py [sol-Python3]
import numpy as np

# 返回 a 的循环自卷积
def self_cyclic_conv(a: list[int]) -> np.ndarray:
    return np.rint(np.fft.ifft(np.fft.fft(a) ** 2).real)

class Solution:
    def minOperations(self, s: str) -> int:
        s = [ord(c) - ord('a') for c in s]
        n = len(s)
        conv_sum = np.zeros(n, dtype=np.float64)
        a = [0] * n
        total = 0

        for k in range(13):
            for i, ch in enumerate(s):
                if k <= ch < k + 13:
                    a[i] = 1
                    total += 1
                else:
                    a[i] = 0
            c = self_cyclic_conv(a)
            conv_sum += c  # 对每个 i 执行 conv_sum[i] += c[i]

        return total + min(rot - int(conv_sum[(rot * 2 - 1) % n]) for rot in range(n))
```

```go [sol-Go]
type fft struct {
	n        int
	omega    []complex128
	omegaInv []complex128
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

// 计算 a 的自卷积
func selfPolyConvFFT(a []int) []int {
	n := len(a)
	limit := 1 << bits.Len(uint(n*2-1))
	A := make([]complex128, limit)
	for i, v := range a {
		A[i] = complex(float64(v), 0)
	}

	t := newFFT(limit)
	t.dft(A)
	for i, x := range A {
		A[i] *= x
	}
	t.idft(A)

	conv := make([]int, n*2-1)
	for i := range conv {
		conv[i] = int(math.Round(real(A[i])))
	}
	return conv
}

// 计算 a 的循环自卷积
func selfCyclicConvFFT(a []int) []int {
	n := len(a)
	conv := selfPolyConvFFT(a)
	for k := range n - 1 {
		conv[k] += conv[n+k]
	}
	return conv[:n]
}

func minOperations(s string) int {
	n := len(s)
	convSum := make([]int, n)
	a := make([]int, n)
	total := 0
	for k := range 13 {
		for i := range n {
			x := int(s[i] - 'a')
			if k <= x && x < k+13 {
				a[i] = 1
				total++
			} else {
				a[i] = 0
			}
		}
		c := selfCyclicConvFFT(a)
		for i, v := range c {
			convSum[i] += v
		}
	}

	ans := math.MaxInt
	for rot := range n {
		c := (rot*2 - 1 + n) % n
		ans = min(ans, rot-convSum[c])
	}
	return ans + total
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(|\Sigma|n\log n)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
