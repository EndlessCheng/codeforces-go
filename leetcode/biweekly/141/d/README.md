## 思路

枚举有 $i$ 个非空（有表演者）的节目，原问题可以拆分成三个问题：

1. 从 $x$ 个节目中，选 $i$ 个节目，且**有顺序**的方案数，即排列数 $A_x^i = \dfrac{x!}{(x-i)!}$。有顺序是因为下面计算的第二类斯特林数是无顺序的划分。
2. 把 $n$ 个人划分成 $i$ 个非空集合的方案数（这 $i$ 个集合没有顺序），即第二类斯特林数 $S(n,i)$。下面细说。
3. 评委会打分的方案数。每个节目有 $y$ 种打分方法，根据乘法原理，$i$ 个节目就是 $y^i$ 种打分方法。

三个问题互相独立，三者相乘得

$$
A_x^i\cdot S(n,i)\cdot y^i
$$

$i$ 从 $1$ 枚举到 $\min(n,x)$，答案为

$$
\sum_{i=1}^{\min(n,x)} A_x^i\cdot S(n,i)\cdot y^i
$$

## 第二类斯特林数

定义 $S(i,j)$ 表示把 $i$ 个人划分成 $j$ 个非空集合的方案数（这 $j$ 个集合没有顺序）。

讨论第 $i$ 个人怎么划分：

- 第 $i$ 个人单独形成一个集合，那么问题变成把 $i-1$ 个人划分成 $j-1$ 个非空集合的方案数，即 $S(i-1,j-1)$。
- 第 $i$ 个人放入前面 $j$ 个集合中，有 $j$ 种方法，问题变成把 $i-1$ 个人划分成 $j$ 个非空集合的方案数，二者相乘得 $j\cdot S(i-1,j)$。

根据加法原理，二者相加得

$$
S(i,j) = S(i-1,j-1) + j\cdot S(i-1,j)
$$

初始值 $S(0,0) = 1$，表示 $0$ 个人划分成 $0$ 个集合，也算一种方案。

> 你也可以从递归的角度理解，递归到 $i=0$ 且 $j=0$ 时，表示找到了一个合法划分方案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1iR2zYaESG/) 第四题，欢迎点赞关注~

```py [sol-Python3]
MOD = 1_000_000_007
MX = 1001

s = [[0] * MX for _ in range(MX)]
s[0][0] = 1
for i in range(1, MX):
    for j in range(1, i + 1):
        s[i][j] = (s[i - 1][j - 1] + j * s[i - 1][j]) % MOD

class Solution:
    def numberOfWays(self, n: int, x: int, y: int) -> int:
        ans = 0
        perm = pow_y = 1
        for i in range(1, min(n, x) + 1):
            perm = perm * (x + 1 - i) % MOD
            pow_y = pow_y * y % MOD
            ans += perm * s[n][i] * pow_y
        return ans % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 1001;
    private static final int[][] s = new int[MX][MX];

    static {
        s[0][0] = 1;
        for (int i = 1; i < MX; i++) {
            for (int j = 1; j <= i; j++) {
                s[i][j] = (int) ((s[i - 1][j - 1] + (long) j * s[i - 1][j]) % MOD);
            }
        }
    }

    public int numberOfWays(int n, int x, int y) {
        long ans = 0;
        long perm = 1;
        long powY = 1;
        for (int i = 1; i <= Math.min(n, x); i++) {
            perm = perm * (x + 1 - i) % MOD;
            powY = powY * y % MOD;
            ans = (ans + perm * s[n][i] % MOD * powY) % MOD;
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 1001;
int s[MX][MX];

auto init = [] {
    s[0][0] = 1;
    for (int i = 1; i < MX; i++) {
        for (int j = 1; j <= i; j++) {
            s[i][j] = (s[i - 1][j - 1] + (long long) j * s[i - 1][j]) % MOD;
        }
    }
    return 0;
}();

class Solution {
public:
    int numberOfWays(int n, int x, int y) {
        long long ans = 0, perm = 1, pow_y = 1;
        for (int i = 1; i <= min(n, x); i++) {
            perm = perm * (x + 1 - i) % MOD;
            pow_y = pow_y * y % MOD;
            ans = (ans + perm * s[n][i] % MOD * pow_y) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const mx = 1001

var s [mx][mx]int

func init() {
	s[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 1; j <= i; j++ {
			s[i][j] = (s[i-1][j-1] + j*s[i-1][j]) % mod
		}
	}
}

func numberOfWays(n, x, y int) (ans int) {
	perm, powY := 1, 1
	for i := 1; i <= min(n, x); i++ {
		perm = perm * (x + 1 - i) % mod
		powY = powY * y % mod
		ans = (ans + perm*s[n][i]%mod*powY) % mod
	}
	return
}
```

#### 复杂度分析

预处理的时间和空间均为 $\mathcal{O}(N^2)$，其中 $N=1000$。

对于 $\texttt{numberOfWays}$：

- 时间复杂度：$\mathcal{O}(\min(n,x))$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面数学题单中的「**组合数学**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
