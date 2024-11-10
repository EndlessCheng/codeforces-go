## 题意

定义 $f(x)$ 为 $x$ 的二进制表示中的 $1$ 的个数。

定义 $f^*(x)$ 为使 $f(f(\cdots f(x))) = 1$ 的最少嵌套（迭代）次数。也就是不断地把 $x$ 更新为 $f(x)$，最少要更新多少次，才能使 $x$ 变成 $1$。

例如 $f^*(6) = 2$，因为 $f(f(6)) = f(2) = 1$。

计算 $[1,s)$ 中有多少个数 $x$ 满足 $f^*(x) \le k$。

## 思路

根据定义，我们有

$$
f^*(x) = f^*(f(x)) + 1
$$

设 $s$ 的长度为 $n$。从小到大递推（写一个线性 DP），即可算出 $[1,n-1]$ 的所有 $f^*(x)$ 值。注意题目要求数字小于 $s$，所以不可能有 $n$ 个 $1$。

对于满足 $f^*(x) \le k-1$ 的所有 $x$，我们需要计算，$[1,s)$ 中有多少个二进制数，**恰好**有 $x$ 个 $1$？

这些恰好有 $x$ 个 $1$ 的二进制数，满足 $f^*(x) \le k$。

这可以用**数位 DP** 解决。原理请看 [数位 DP 通用模板](https://www.bilibili.com/video/BV1rS4y1s721/?t=20m05s)。

⚠**注意**：本题需要严格小于 $s$，这可以用 $\textit{isLimit}$ 判断：递归到 $i=n$ 时，若仍有 $\textit{isLimit}=\texttt{true}$，则返回 $0$。此外，我们只关心 $1$ 的个数，是否有前导零无影响，所以无需 $\textit{isNum}$ 参数。

代码实现时，可以定义 $f^*(1) = 1$，上文中 $f^*(x) \le k-1$ 可以简化为 $f^*(x) \le k$。

记得取模。

关于取模的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1JVmBYvEnD/?t=22m35s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countKReducibleNumbers(self, s: str, k: int) -> int:
        MOD = 1_000_000_007
        n = len(s)

        @cache
        def dfs(i: int, left1: int, is_limit: bool) -> int:
            if i == n:
                return 0 if is_limit or left1 else 1
            up = int(s[i]) if is_limit else 1
            res = 0
            for d in range(min(up, left1) + 1):
                res += dfs(i + 1, left1 - d, is_limit and d == up)
            return res % MOD

        ans = 0
        f = [0] * n
        for i in range(1, n):
            f[i] = f[i.bit_count()] + 1
            if f[i] <= k:
                # 计算有多少个二进制数恰好有 i 个 1
                ans += dfs(0, i, True)
        dfs.cache_clear()  # 防止爆内存
        return ans % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countKReducibleNumbers(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[][] memo = new int[n][n];
        for (int[] row : memo) {
            Arrays.fill(row, -1);
        }

        long ans = 0;
        int[] f = new int[n];
        for (int i = 1; i < n; i++) {
            f[i] = f[Integer.bitCount(i)] + 1;
            if (f[i] <= k) {
                // 计算有多少个二进制数恰好有 i 个 1
                ans += dfs(0, i, true, s, memo);
            }
        }
        return (int) (ans % MOD);
    }

    private int dfs(int i, int left1, boolean isLimit, char[] s, int[][] memo) {
        if (i == s.length) {
            return !isLimit && left1 == 0 ? 1 : 0;
        }
        if (!isLimit && memo[i][left1] != -1) {
            return memo[i][left1];
        }
        int up = isLimit ? s[i] - '0' : 1;
        int res = 0;
        for (int d = 0; d <= Math.min(up, left1); d++) {
            res = (res + dfs(i + 1, left1 - d, isLimit && d == up, s, memo)) % MOD;
        }
        if (!isLimit) {
            memo[i][left1] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countKReducibleNumbers(string s, int k) {
        const int MOD = 1'000'000'007;
        int n = s.length();
        vector<vector<int>> memo(n, vector<int>(n, -1));
        auto dfs = [&](auto& dfs, int i, int left1, bool is_limit) -> int {
            if (i == n) {
                return !is_limit && left1 == 0;
            }
            if (!is_limit && memo[i][left1] != -1) {
                return memo[i][left1];
            }
            int up = is_limit ? s[i] - '0' : 1;
            int res = 0;
            for (int d = 0; d <= min(up, left1); d++) {
                res = (res + dfs(dfs, i + 1, left1 - d, is_limit && d == up)) % MOD;
            }
            if (!is_limit) {
                memo[i][left1] = res;
            }
            return res;
        };

        long long ans = 0;
        vector<int> f(n);
        for (int i = 1; i < n; i++) {
            f[i] = f[__builtin_popcount(i)] + 1;
            if (f[i] <= k) {
                // 计算有多少个二进制数恰好有 i 个 1
                ans += dfs(dfs, 0, i, true);
            }
        }
        return ans % MOD;
    }
};
```

```go [sol-Go]
func countKReducibleNumbers(s string, k int) (ans int) {
    const mod = 1_000_000_007
    n := len(s)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(int, int, bool) int
    dfs = func(i, left1 int, isLimit bool) (res int) {
        if i == n {
            if !isLimit && left1 == 0 {
                return 1
            }
            return
        }
        if !isLimit {
            p := &memo[i][left1]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        up := 1
        if isLimit {
            up = int(s[i] - '0')
        }
        for d := 0; d <= min(up, left1); d++ {
            res += dfs(i+1, left1-d, isLimit && d == up)
        }
        return res % mod
    }

    f := make([]int, n)
    for i := 1; i < n; i++ {
        f[i] = f[bits.OnesCount(uint(i))] + 1
        if f[i] <= k {
            // 计算有多少个二进制数恰好有 i 个 1
            ans += dfs(0, i, true)
        }
    }
    return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

更多相似题目，见下面动态规划题单中的「**十、数位 DP**」。

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
