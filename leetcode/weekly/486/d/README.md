从高到低，依次考虑答案二进制的第 $i$ 位填 $0$ 还是填 $1$。

在示例 1 中，我们要计算第 $4$ 小的二进制数，恰好包含 $2$ 个 $1$。

如果答案是 $0\texttt{\_\_\_}$，那么我们要在剩下的 $3$ 个位置上填 $2$ 个 $1$，这一共有 $\binom 3 2 = 3$ 种方案，说明前 $3$ 小的二进制数都是 $0\texttt{\_\_\_}$。但 $n=4 > 3$，所以第 $4$ 小的二进制数一定是 $1\texttt{\_\_\_}$。

填入 $1$ 后，问题变成：

- 计算第 $4-3=1$ 小的二进制数，恰好包含 $2-1=1$ 个 $1$。

这是一个规模更小的子问题，用同样的方法解决。最终我们得到答案 $1001$。

一般地，枚举 $i=49,48,\ldots,0$，计算在 $i$ 个位置上填 $k$ 个 $1$ 的方案数 $c = \binom i k$，分类讨论：

- 如果 $n\le c$，那么答案二进制的 $i$ 位填 $0$。
- 如果 $n > c$，说明 $n$ 比较大，答案二进制的 $i$ 位填 $1$。然后把 $n$ 减少 $c$，$k$ 减少 $1$，解决这个规模更小的子问题。

代码实现时，可以预处理 $50$ 以内的组合数（题目保证答案小于 $2^{50}$），原理见 [118. 杨辉三角](https://leetcode.cn/problems/pascals-triangle/)，[我的题解](https://leetcode.cn/problems/pascals-triangle/solutions/2784222/jian-dan-ti-jian-dan-zuo-pythonjavaccgoj-z596/)。

[本题视频讲解](https://www.bilibili.com/video/BV1W2zQBnE3g/?t=17m52s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def nthSmallest(self, n: int, k: int) -> int:
        ans = 0
        for i in range(49, -1, -1):
            c = comb(i, k)  # 第 i 位填 0 的方案数
            if n > c:  # n 比较大，第 i 位必须填 1
                n -= c
                ans |= 1 << i
                k -= 1  # 维护剩余的 1 的个数
                if k == 0:  # 填完了 1，ans 剩余低位都是 0
                    return ans
```

```py [sol-Python3 预处理]
# 预处理组合数
MX = 50
comb = [[0] * (MX + 1) for _ in range(MX)]
for i in range(MX):
    comb[i][0] = 1
    for j in range(1, i + 1):
        comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j]

class Solution:
    def nthSmallest(self, n: int, k: int) -> int:
        ans = 0
        for i in range(MX - 1, -1, -1):
            c = comb[i][k]  # 第 i 位填 0 的方案数
            if n > c:  # n 比较大，第 i 位必须填 1
                n -= c
                ans |= 1 << i
                k -= 1  # 维护剩余的 1 的个数
                if k == 0:  # 填完了 1，ans 剩余低位都是 0
                    return ans
```

```java [sol-Java]
class Solution {
    private static final int MX = 50;
    private static final long[][] comb = new long[MX][MX + 1];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理组合数
        for (int i = 0; i < MX; i++) {
            comb[i][0] = 1;
            for (int j = 1; j <= i; j++) {
                comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
            }
        }
    }

    public long nthSmallest(long n, int k) {
        long ans = 0;
        for (int i = MX - 1; k > 0; i--) {
            long c = comb[i][k]; // 第 i 位填 0 的方案数
            if (n > c) { // n 比较大，第 i 位必须填 1
                n -= c;
                ans |= 1L << i;
                k--; // 维护剩余的 1 的个数
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int MX = 50;
long long comb[MX][MX + 1];

auto init = [] {
    // 预处理组合数
    for (int i = 0; i < MX; i++) {
        comb[i][0] = 1;
        for (int j = 1; j <= i; j++) {
            comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
        }
    }
    return 0;
}();

class Solution {
public:
    long long nthSmallest(long long n, int k) {
        long long ans = 0;
        for (int i = MX - 1; k > 0; i--) {
            long long c = comb[i][k]; // 第 i 位填 0 的方案数
            if (n > c) { // n 比较大，第 i 位必须填 1
                n -= c;
                ans |= 1LL << i;
                k--; // 维护剩余的 1 的个数
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
const mx = 50
var comb [mx][mx + 1]int64

func init() {
	// 预处理组合数
	for i := range comb {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}
}

func nthSmallest(n int64, k int) (ans int64) {
	for i := mx - 1; k > 0; i-- {
		c := comb[i][k] // 第 i 位填 0 的方案数
		if n > c { // n 比较大，第 i 位必须填 1
			n -= c
			ans |= 1 << i
			k-- // 维护剩余的 1 的个数
		}
	}
	return
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(W)$，其中 $W=50$ 是答案二进制长度的最大值。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面位运算题单的「**五、试填法**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
