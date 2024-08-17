## 前言

本题属于 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的状压 DP「排列型 ② 相邻相关」。

如果你从未做过状压 DP，推荐先做一道「排列型 ① 相邻无关」的题目，例如 [526. 优美的排列](https://leetcode.cn/problems/beautiful-arrangement/)。

由于本题与 526 题类似，所以下面的内容，我会接着 [526 题解](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/) 继续讲。

## 一、记忆化搜索

下面会用到一些集合论的术语和符号，请看 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)。

相比 526（题解的第二种状态定义），本题需要额外知道上一个选的数的下标是多少，因此要多一个参数。

定义 $\textit{dfs}(S,i)$ 表示在可以选的下标集合为 $S$，上一个选的数的下标是 $i$ 时，可以构造出多少个特别排列。

枚举当前要选的数的下标 $j$，那么接下来要解决的问题是，在可以选的下标集合为 $S\setminus \{j\}$，上一个选的数的下标是 $j$ 时，可以构造出多少个特别排列。

累加这些方案数，得

$$
\textit{dfs}(S,i) = \sum_{j\in S} \textit{dfs}(S\setminus \{j\},j)
$$

其中 $j$ 满足 $\textit{nums}[j]\bmod \textit{nums}[i]=0$ 或 $\textit{nums}[i]\bmod \textit{nums}[j]=0$。

**递归边界**：$\textit{dfs}(\varnothing,i) = 1$，表示找到了一个特别排列。

**递归入口**：$\textit{dfs}(U\setminus \{i\},i)$，其中全集 $U=\{0,1,2,\cdots,n-1\}$。

枚举特别排列的第一个数的下标 $i$，累加所有 $\textit{dfs}(U\setminus \{i\},i)$，即为答案。

```py [sol-Python3]
class Solution:
    def specialPerm(self, nums: List[int]) -> int:
        @cache
        def dfs(s: int, i: int) -> int:
            if s == 0:
                return 1  # 找到一个特别排列
            res = 0
            pre = nums[i]
            for j, x in enumerate(nums):
                if s >> j & 1 and (pre % x == 0 or x % pre == 0):
                    res += dfs(s ^ (1 << j), j)
            return res

        n = len(nums)
        u = (1 << n) - 1
        return sum(dfs(u ^ (1 << i), i) for i in range(n)) % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int specialPerm(int[] nums) {
        int n = nums.length;
        int u = (1 << n) - 1;
        long[][] memo = new long[u][n];
        for (long[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        long ans = 0;
        for (int i = 0; i < n; i++) {
            ans += dfs(u ^ (1 << i), i, nums, memo);
        }
        return (int) (ans % 1_000_000_007);
    }

    private long dfs(int s, int i, int[] nums, long[][] memo) {
        if (s == 0) {
            return 1; // 找到一个特别排列
        }
        if (memo[s][i] != -1) { // 之前计算过
            return memo[s][i];
        }
        long res = 0;
        for (int j = 0; j < nums.length; j++) {
            if ((s >> j & 1) > 0 && (nums[i] % nums[j] == 0 || nums[j] % nums[i] == 0)) {
                res += dfs(s ^ (1 << j), j, nums, memo);
            }
        }
        return memo[s][i] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int specialPerm(vector<int>& nums) {
        int n = nums.size(), u = (1 << n) - 1;
        vector<vector<long long>> memo(u, vector<long long>(n, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int s, int i) -> long long {
            if (s == 0) {
                return 1; // 找到一个特别排列
            }
            auto& res = memo[s][i]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = 0;
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) && (nums[i] % nums[j] == 0 || nums[j] % nums[i] == 0)) {
                    res += dfs(dfs, s ^ (1 << j), j);
                }
            }
            return res;
        };
        long long ans = 0;
        for (int i = 0; i < n; i++) {
            ans += dfs(dfs, u ^ (1 << i), i);
        }
        return ans % 1'000'000'007;
    }
};
```

```go [sol-Go]
func specialPerm(nums []int) (ans int) {
    n := len(nums)
    u := 1<<n - 1
    memo := make([][]int, u)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    var dfs func(int, int) int
    dfs = func(s, i int) (res int) {
        if s == 0 {
            return 1 // 找到一个特别排列
        }
        p := &memo[s][i]
        if *p != -1 { // 之前计算过
            return *p
        }
        for j, x := range nums {
            if s>>j&1 > 0 && (nums[i]%x == 0 || x%nums[i] == 0) {
                res += dfs(s^(1<<j), j)
            }
        }
        *p = res // 记忆化
        return
    }
    for i := range nums {
        ans += dfs(u^(1<<i), i)
    }
    return ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^22^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，因此时间复杂度为 $\mathcal{O}(n^22^n)$。
- 空间复杂度：$\mathcal{O}(n2^n)$。保存多少状态，就需要多少空间。

## 二、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[S][i]$ 的定义和 $\textit{dfs}(S,i)$ 的定义是一样的，都表示在可以选的下标集合为 $S$，上一个选的数的下标是 $i$ 时，可以构造出多少个特别排列。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[S][i] =\sum_{j\in S} f[S\setminus \{j\}][j]
$$

初始值 $f[\varnothing][i]=1$，翻译自递归边界 $\textit{dfs}(\varnothing,i)=1$。

答案为 $f[U\setminus \{i\}][i]$ 之和，翻译自递归入口 $\textit{dfs}(U\setminus \{i\},i)$。

> 注：在随机数据下，由于相邻元素无法整除，记忆化搜索有很多状态无法访问到，记忆化搜索比递推快。

```py [sol-Python3]
class Solution:
    def specialPerm(self, nums: List[int]) -> int:
        n = len(nums)
        u = (1 << n) - 1
        f = [[0] * n for _ in range(u)]
        f[0] = [1] * n
        for s in range(1, u):
            for i, pre in enumerate(nums):
                if s >> i & 1:
                    continue
                for j, x in enumerate(nums):
                    if s >> j & 1 and (pre % x == 0 or x % pre == 0):
                        f[s][i] += f[s ^ (1 << j)][j]
        return sum(f[u ^ (1 << i)][i] for i in range(n)) % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int specialPerm(int[] nums) {
        int n = nums.length;
        int u = (1 << n) - 1;
        long[][] f = new long[u][n];
        Arrays.fill(f[0], 1L);
        for (int s = 1; s < u; s++) {
            for (int i = 0; i < n; i++) {
                if ((s >> i & 1) != 0) {
                    continue;
                }
                for (int j = 0; j < n; j++) {
                    if ((s >> j & 1) != 0 && (nums[i] % nums[j] == 0 || nums[j] % nums[i] == 0)) {
                        f[s][i] += f[s ^ (1 << j)][j];
                    }
                }
            }
        }
        long ans = 0;
        for (int i = 0; i < n; i++) {
            ans += f[u ^ (1 << i)][i];
        }
        return (int) (ans % 1_000_000_007);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int specialPerm(vector<int>& nums) {
        int n = nums.size(), u = (1 << n) - 1;
        vector<vector<long long>> f(u, vector<long long>(n));
        ranges::fill(f[0], 1LL);
        for (int s = 1; s < u; s++) {
            for (int i = 0; i < n; i++) {
                if (s >> i & 1) {
                    continue;
                }
                for (int j = 0; j < n; j++) {
                    if ((s >> j & 1) && (nums[i] % nums[j] == 0 || nums[j] % nums[i] == 0)) {
                        f[s][i] += f[s ^ (1 << j)][j];
                    }
                }
            }
        }
        long long ans = 0;
        for (int i = 0; i < n; i++) {
            ans += f[u ^ (1 << i)][i];
        }
        return ans % 1'000'000'007;
    }
};
```

```go [sol-Go]
func specialPerm2(nums []int) (ans int) {
    n := len(nums)
    u := 1<<n - 1
    f := make([][]int, u)
    for i := range f {
        f[i] = make([]int, n)
    }
    for i := range nums {
        f[0][i] = 1
    }
    for s := 1; s < u; s++ {
        for i, pre := range nums {
            if s>>i&1 != 0 {
                continue
            }
            for j, x := range nums {
                if s>>j&1 != 0 && (pre%x == 0 || x%pre == 0) {
                    f[s][i] += f[s^(1<<j)][j]
                }
            }
        }
    }
    for i := range nums {
        ans += f[u^(1<<i)][i]
    }
    return ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^22^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n2^n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
