## 分析

下文把 $\textit{arr}$ 简记为 $a$。

题目说，每个 $a_i$ 都可以被 $a_{i-1}$ 整除，即 $\dfrac{a_i}{a_{i-1}}$ 是整数。

比如 $a=[1,1,2,4,4,8,8,8]$ 是符合题目要求的。看上去，$a$ 中有很多重复元素，或者说，不同元素个数并不会很多。

想一想，如果 $n=10^4$，$\textit{maxValue}=8$，那么 $a$ 中至多有多少个**不同**的元素？

如果 $a_{i-1}\ne a_i$，那么 $a_i$ 至少是 $2\cdot a_{i-1}$。假设 $a_0=1$，至多乘三次 $2$，得到 $8$，就不能再变大了，所以 $a$ 中至多有 $4$ 个不同的元素，即 $1,2,4,8$。

这个例子说明，即使 $n$ 很大，$a$ 中也至多有 $\left\lfloor\log_2 \textit{maxValue}\right\rfloor + 1$ 个不同的元素。

这启发我们重点考虑 $a_{i-1}\ne a_i$ 的情况，也就是 $\dfrac{a_i}{a_{i-1}} > 1$ 的情况。

## 商分

类似 [差分](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，我们来计算 $a$ 的「商分」，即相邻两数的商。

定义 $q_0 = a_0$，$q_i = \dfrac{a_i}{a_{i-1}}\ (i\ge 1)$。

例如：

- $a=[2,2,4,8,8]$ 的商分数组为 $q=[2,1,2,2,1]$。
- $a=[1,4,4,8,8]$ 的商分数组为 $q=[1,4,1,2,1]$。

不同的 $a$，唯一对应着不同的 $q$。计算理想数组 $a$ 的个数，可以转化成计算商分数组 $q$ 的个数。

根据 $q$ 的定义，所有 $q_i$ 的乘积等于

$$
\prod_{i=0}^{n-1}q_i = a_0\cdot \dfrac{a_1}{a_0} \cdot \dfrac{a_2}{a_1} \cdots \dfrac{a_{n-1}}{a_{n-2}} = a_{n-1}
$$

现在假设 $a_{n-1}=8$，也就是所有 $q_i$ 的乘积等于 $8$，有多少个不同的 $q$？

## 放球问题

这个问题等价于，有 $n$ 个位置，把 $3$ 个 $2$ 分配到 $n$ 个位置的方案数。注：分配到同一个位置就乘起来，比如 $2$ 个 $2$ 分配到同一个位置就是 $4$。没分配 $2$ 的位置是 $1$。

这等价于如下放球问题：

- 把 $k=3$ 个无区别的小球放到 $n$ 个有区别的盒子中，允许盒子为空，一个盒子也可以放多个小球，有多少种不同的放法？

![lc2338-c.png](https://pic.leetcode.cn/1744685351-GGrXfu-lc2338-c.png)

## 思路

枚举 $a_{n-1}=1,2,3,\ldots,\textit{maxValue}$，根据上图最后的公式，计算方案数，加到答案中。

## 如何分解质因子

#### 方法一：枚举（适用于本题）

计算 $x$ 每个质因子的个数。从 $i=2$ 开始枚举，如果 $x$ 能被 $i$ 整除，就不断地除 $i$，直到 $x$ 不能被 $i$ 整除为止，统计除 $i$ 的次数，即为 $x$ 中的质因子 $i$ 的出现次数。

什么时候停止枚举呢？如果 $i^2 > x$，继续向后枚举是不会出现 $x$ 被 $i$ 整除的情况的。这可以用**反证法**证明：假设存在 $i$，满足 $i^2>x$ 且 $x$ 能被 $i$ 整除，那么 $x$ 也能被 $\dfrac{x}{i}$ 整除，注意到 $\dfrac{x}{i}<i$，但我们已经处理完小于 $i$ 的质因子了，不会出现 $x$ 仍可以被一个小于 $i$ 的质因子整除的情况，矛盾。所以当 $i^2 > x$ 时可以停止枚举。

循环结束后，如果 $x>1$，说明还有一个质因子为 $x$。

#### 方法二：预处理 LPF（适用于更大的数据范围）

利用埃氏筛或者欧拉筛，用质数 $p$ 标记 $p$ 的倍数（跳过已经标记的数），我们可以预处理每个数 $x$ 的最小质因子 $\text{LPF}[x]$。不断地更新 $x$ 为 $\dfrac{x}{\text{LPF}[x]}$，直到 $x=1$，在这个过程中统计每个质因子的出现次数。

## 如何计算组合数

#### 方法一：递推（适用于本题）

对于从 $n$ 个物品中选择 $k$ 个物品的方案数 $C(n,k)$，可以用「选或不选」来思考，对于第 $n$ 个物品：

- 不选：问题变成从 $n-1$ 个物品中选择 $k$ 个物品的方案数 $C(n-1,k)$。
- 选：问题变成从 $n-1$ 个物品中选择 $k-1$ 个物品的方案数 $C(n-1,k-1)$。

所以 $C(n,k) = C(n-1,k) + C(n-1,k-1)$。

初始值：$C(n,0) = 1$。

对于本题，由于 $2^{13} < 10^4 < 2^{14}$，我们可以预处理 $n\le 10^4 + 13-1$ 和 $k \le 13$ 的组合数。

#### 方法二：预处理阶乘及其逆元（适用于更大的数据范围）

见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
MOD = 1_000_000_007
MAX_N = 10_000
MAX_E = 13

# EXP[x] 为 x 分解质因数后，每个质因数的指数
EXP = [[] for _ in range(MAX_N + 1)]
for x in range(2, len(EXP)):
    t = x
    i = 2
    while i * i <= t:
        e = 0
        while t % i == 0:
            e += 1
            t //= i
        if e:
            EXP[x].append(e)
        i += 1
    if t > 1:
        EXP[x].append(1)

# 预处理组合数
C = [[0] * (MAX_E + 1) for _ in range(MAX_N + MAX_E)]
for i in range(len(C)):
    C[i][0] = 1
    for j in range(1, min(i, MAX_E) + 1):
        C[i][j] = (C[i - 1][j] + C[i - 1][j - 1]) % MOD

class Solution:
    def idealArrays(self, n: int, maxValue: int) -> int:
        ans = 0
        for x in range(1, maxValue + 1):
            res = 1
            for e in EXP[x]:
                res = res * C[n + e - 1][e] % MOD
            ans += res
        return ans % MOD
```

```py [sol-Python3 库函数]
MOD = 1_000_000_007
MAX_N = 10_000

# EXP[x] 为 x 分解质因数后，每个质因数的指数
EXP = [[] for _ in range(MAX_N + 1)]
for x in range(2, len(EXP)):
    t = x
    i = 2
    while i * i <= t:
        e = 0
        while t % i == 0:
            e += 1
            t //= i
        if e:
            EXP[x].append(e)
        i += 1
    if t > 1:
        EXP[x].append(1)

class Solution:
    def idealArrays(self, n: int, maxValue: int) -> int:
        ans = 0
        for x in range(1, maxValue + 1):
            res = 1
            for e in EXP[x]:
                res = res * comb(n + e - 1, e) % MOD
            ans += res
        return ans % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MAX_N = 10_000;
    private static final int MAX_E = 13;

    private static final List<Integer>[] EXP = new ArrayList[MAX_N + 1];
    private static final int[][] C = new int[MAX_N + MAX_E][MAX_E + 1];

    private static boolean done = false;

    private void init() {
        // 这样写比 static block 更快
        if (done) {
            return;
        }
        done = true;

        // EXP[x] 为 x 分解质因数后，每个质因数的指数
        for (int x = 1; x < EXP.length; x++) {
            EXP[x] = new ArrayList<>();
            int t = x;
            for (int i = 2; i * i <= t; i++) {
                int e = 0;
                for (; t % i == 0; t /= i) {
                    e++;
                }
                if (e > 0) {
                    EXP[x].add(e);
                }
            }
            if (t > 1) {
                EXP[x].add(1);
            }
        }

        // 预处理组合数
        for (int i = 0; i < MAX_N + MAX_E; i++) {
            C[i][0] = 1;
            for (int j = 1; j <= Math.min(i, MAX_E); j++) {
                C[i][j] = (C[i - 1][j] + C[i - 1][j - 1]) % MOD;
            }
        }
    }

    public int idealArrays(int n, int maxValue) {
        init();
        long ans = 0;
        for (int x = 1; x <= maxValue; x++) {
            long mul = 1;
            for (int e : EXP[x]) {
                mul = mul * C[n + e - 1][e] % MOD;
            }
            ans += mul;
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MAX_N = 10'000;
const int MAX_E = 13;

vector<int> EXP[MAX_N + 1]; 
int C[MAX_N + MAX_E][MAX_E + 1];

int init = []() {
    // EXP[x] 为 x 分解质因数后，每个质因数的指数
    for (int x = 2; x <= MAX_N; x++) {
        int t = x;
        for (int i = 2; i * i <= t; i++) {
            int e = 0;
            for (; t % i == 0; t /= i) {
                e++;
            }
            if (e) {
                EXP[x].push_back(e);
            }
        }
        if (t > 1) {
            EXP[x].push_back(1);
        }
    }

    // 预处理组合数
    for (int i = 0; i < MAX_N + MAX_E; i++) {
        C[i][0] = 1;
        for (int j = 1; j <= min(i, MAX_E); j++) {
            C[i][j] = (C[i - 1][j] + C[i - 1][j - 1]) % MOD;
        }
    }
    return 0;
}();

class Solution {
public:
    int idealArrays(int n, int maxValue) {
        long long ans = 0;
        for (int x = 1; x <= maxValue; x++) {
            long long res = 1;
            for (int e : EXP[x]) {
                res = res * C[n + e - 1][e] % MOD;
            }
            ans += res;
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007
const maxN = 10_000
const maxE = 13

var exp [maxN + 1][]int
var c [maxN + maxE][maxE + 1]int

func init() {
	// exp[x] 为 x 分解质因数后，每个质因数的指数
	for x := 2; x <= maxN; x++ {
		t := x
		for i := 2; i*i <= t; i++ {
			e := 0
			for ; t%i == 0; t /= i {
				e++
			}
			if e > 0 {
				exp[x] = append(exp[x], e)
			}
		}
		if t > 1 {
			exp[x] = append(exp[x], 1)
		}
	}

	// 预处理组合数
	for i := range c {
		c[i][0] = 1
		for j := 1; j <= min(i, maxE); j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}
}

func idealArrays(n, maxValue int) (ans int) {
	for x := 1; x <= maxValue; x++ {
		res := 1
		for _, e := range exp[x] {
			res = res * c[n+e-1][e] % mod
		}
		ans += res
	}
	return ans % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(m\log\log m)$，其中 $m$ 表示 $\textit{maxValue}$。循环次数等同于 $[1,m]$ 中的每个数的不同质因子个数之和。由于每个质数 $p$ 的贡献不超过 $\dfrac{m}{p}$，累加得 $m\displaystyle\sum\limits_{p\le m}\dfrac{1}{p} = \mathcal{O}(m\log\log m)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[1735. 生成乘积数组的方案数](https://leetcode.cn/problems/count-ways-to-make-array-with-product/)

更多相似题目，见下面数学题单的「**§2.3 放球问题**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
