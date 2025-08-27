## 分析

$a$ 除以 $4$ 下取整等价于 $a$ 右移 $2$ 位。所以把一个长为 $m$ 的二进制数变成 $0$，需要操作 $\left\lceil\dfrac{m}{2}\right\rceil$ 次。

例如 $l=1$，$r=5$，如果每次操作只把一个数右移 $2$ 位，那么 $1$ 到 $5$ 每个数的操作次数分别为

$$
\textit{ops}=[1,1,1,2,2]
$$

本题一次可以操作两个数，问题等价于：

- 每次操作把 $\textit{ops}$ 中的两个数都减去 $1$，问：要让 $\textit{ops}$ 中没有正数，至少要操作多少次？

**分析**：设 $\textit{tot}=\sum \textit{ops}[i]$，$\textit{mx}=\max(\textit{ops})$。假如每次可以把 $\textit{tot}$ 减少 $2$，那么把 $\textit{tot}$ 减少到 $\le 0$，至少要操作 $\left\lceil\dfrac{\textit{tot}}{2}\right\rceil$ 次。但如果 $\textit{mx}$ 很大，操作次数就等于 $\textit{mx}$（每次操作选 $\textit{mx}$ 和另一个数）。

**定理**：最小操作次数为

$$
\max\left(\left\lceil\dfrac{\textit{tot}}{2}\right\rceil, \textit{mx}\right)
$$

[证明+具体构造方案](https://leetcode.cn/problems/reorganize-string/solutions/2779462/tan-xin-gou-zao-pai-xu-bu-pai-xu-liang-c-h9jg/)

本题由于 $\textit{nums}$ 中的数字是连续整数，相邻数字的操作次数至多相差 $1$。

例如两个数的情况 $\textit{ops}=[x-1,x]$，得到 $\textit{mx}=x$，$\textit{tot} = 2x-1$。由于 $\left\lceil\dfrac{2x-1}{2}\right\rceil = x$，所以 $\textit{mx} \le \left\lceil\dfrac{\textit{tot}}{2}\right\rceil$ 成立。其他情况元素更多，$\textit{tot}$ 更大，$\textit{mx} \le \left\lceil\dfrac{\textit{tot}}{2}\right\rceil$ 更加成立。

所以本题

$$
\textit{mx} \le \left\lceil\dfrac{\textit{tot}}{2}\right\rceil
$$

恒成立。

所以

$$
\max\left(\left\lceil\dfrac{\textit{tot}}{2}\right\rceil, \textit{mx}\right) = \left\lceil\dfrac{\textit{tot}}{2}\right\rceil
$$

算出了 $\textit{tot}$，就算出了答案。

## 公式

定义 $f(n)$ 为 $[1,n]$ 中的单个数的操作次数之和。

设 $n$ 的二进制长度为 $m$，那么：

- 对于长为 $i$ 的二进制数（其中 $1\le i\le m-1$），最小是 $2^{i-1}$，最大是 $2^i-1$，共有 $2^{i-1}$ 个，每个需要操作 $\left\lceil\dfrac{i}{2}\right\rceil$ 次。
- 对于长为 $m$ 的二进制数，最小是 $2^{m-1}$，最大是 $n$，共有 $n+1-2^{m-1}$ 个，每个需要操作 $\left\lceil\dfrac{m}{2}\right\rceil$ 次。

累加得

$$
f(n) = \sum_{i=1}^{m-1} \left\lceil\dfrac{i}{2}\right\rceil 2^{i-1} + \left\lceil\dfrac{m}{2}\right\rceil(n+1-2^{m-1})
$$

$[l,r]$ 中的单个数的操作次数之和为 

$$
\textit{tot} = f(r) - f(l-1)
$$

每次操作至多两个数，操作次数为

$$
\left\lceil\dfrac{f(r) - f(l-1)}{2}\right\rceil = \left\lfloor\dfrac{ f(r) - f(l-1) + 1}{2}\right\rfloor
$$

[本题视频讲解](https://www.bilibili.com/video/BV12eXYYVE5H/?t=22m46s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
# 返回 [1,n] 的单个元素的操作次数之和
def f(n: int) -> int:
    m = n.bit_length()
    res = sum((i + 1) // 2 << (i - 1) for i in range(1, m))
    return res + (m + 1) // 2 * (n + 1 - (1 << m >> 1))

class Solution:
    def minOperations(self, queries: List[List[int]]) -> int:
        return sum((f(r) - f(l - 1) + 1) // 2 for l, r in queries)
```

```java [sol-Java]
class Solution {
    public long minOperations(int[][] queries) {
        long ans = 0;
        for (int[] q : queries) {
            ans += (f(q[1]) - f(q[0] - 1) + 1) / 2;
        }
        return ans;
    }

    // 返回 [1,n] 的单个元素的操作次数之和
    private long f(int n) {
        int m = 32 - Integer.numberOfLeadingZeros(n);
        long res = 0;
        for (int i = 1; i < m; i++) {
            res += (long) (i + 1) / 2 << (i - 1);
        }
        return res + (long) (m + 1) / 2 * (n + 1 - (1 << m >> 1));
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回 [1,n] 的单个元素的操作次数之和
    long long f(int n) {
        int m = bit_width((uint32_t) n);
        long long res = 0;
        for (int i = 1; i < m; i++) {
            res += 1LL * (i + 1) / 2 << (i - 1);
        }
        return res + 1LL * (m + 1) / 2 * (n + 1 - (1 << m >> 1));
    }

public:
    long long minOperations(vector<vector<int>>& queries) {
        long long ans = 0;
        for (auto& q : queries) {
            ans += (f(q[1]) - f(q[0] - 1) + 1) / 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
// 返回 [1,n] 的单个元素的操作次数之和
func f(n int) (res int) {
	m := bits.Len(uint(n))
	for i := 1; i < m; i++ {
		res += (i + 1) / 2 << (i - 1)
	}
	return res + (m+1)/2*(n+1-1<<m>>1)
}

func minOperations(queries [][]int) int64 {
	ans := 0
	for _, q := range queries {
		ans += (f(q[1]) - f(q[0]-1) + 1) / 2
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $r$ 的最大值。每个询问需要 $\mathcal{O}(\log U)$ 的时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化

$\mathcal{O}(1)$ 计算 $f(n)$。

设 $n$ 的二进制长度为 $m$。设 $k$ 为小于 $m$ 的最大偶数。

上面代码的循环，把二进制长为 $1,2$ 的分成一组（每个数操作 $1$ 次），长为 $3,4$ 的分成一组（每个数操作 $2$ 次），长为 $5,6$ 的分成一组（每个数操作 $3$ 次）……依此类推，累加得

$$
(2^2-2^0)\cdot 1 + (2^4-2^2)\cdot 2 + \cdots + (2^k-2^{k-2})\cdot \dfrac{k}{2}
$$

利用错位相减法（详见 [视频讲解](https://www.bilibili.com/video/BV12eXYYVE5H/?t=22m46s)），上式可化简为

$$
k\cdot 2^{k-1} - \dfrac{2^k-1}{3}
$$

对于长为 $[k+1,m]$ 的二进制数，最小是 $2^k$，最大是 $n$，共有 $n+1-2^k$ 个，每个需要操作 $\left\lceil\dfrac{m}{2}\right\rceil$ 次。

相加得

$$
f(n) = k\cdot 2^{k-1} - \dfrac{2^k-1}{3} + \left\lceil\dfrac{m}{2}\right\rceil(n+1-2^k)
$$

代码实现时，如果 $k=0$，没法左移 $k-1=-1$ 位。可以改为先左移 $k$ 位，再右移一位，这样无需特判 $k=0$ 的情况。

```py [sol-Python3]
def f(n: int) -> int:
    if n == 0:
        return 0
    m = n.bit_length()
    k = (m - 1) // 2 * 2
    res = (k << k >> 1) - (1 << k) // 3  # -1 可以省略
    return res + (m + 1) // 2 * (n + 1 - (1 << k))

class Solution:
    def minOperations(self, queries: List[List[int]]) -> int:
        return sum((f(r) - f(l - 1) + 1) // 2 for l, r in queries)
```

```java [sol-Java]
class Solution {
    public long minOperations(int[][] queries) {
        long ans = 0;
        for (int[] q : queries) {
            ans += (f(q[1]) - f(q[0] - 1) + 1) / 2;
        }
        return ans;
    }

    private long f(int n) {
        int m = 32 - Integer.numberOfLeadingZeros(n);
        int k = (m - 1) / 2 * 2;
        long res = ((long) k << k >> 1) - (1 << k) / 3; // -1 可以省略
        return res + (long) (m + 1) / 2 * (n + 1 - (1 << k));
    }
}
```

```cpp [sol-C++]
class Solution {
    long long f(int n) {
        int m = bit_width((uint32_t) n);
        int k = (m - 1) / 2 * 2;
        long long res = (1LL * k << k >> 1) - (1 << k) / 3; // -1 可以省略
        return res + 1LL * (m + 1) / 2 * (n + 1 - (1 << k));
    }

public:
    long long minOperations(vector<vector<int>>& queries) {
        long long ans = 0;
        for (auto& q : queries) {
            ans += (f(q[1]) - f(q[0] - 1) + 1) / 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func f(n int) int {
	m := bits.Len(uint(n))
	k := (m - 1) / 2 * 2
	res := k<<k>>1 - 1<<k/3 // -1 可以省略
	return res + (m+1)/2*(n+1-1<<k)
}

func minOperations(queries [][]int) int64 {
	ans := 0
	for _, q := range queries {
		ans += (f(q[1]) - f(q[0]-1) + 1) / 2
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q)$，其中 $q$ 是 $\textit{queries}$ 的长度。每个询问只需 $\mathcal{O}(1)$ 时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
