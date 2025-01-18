## 分析

对于一个栈（数组），我们只能移除其前缀。注意题目说数组 $\textit{piles}[i]$ 从左到右表示**栈顶到栈底**。

对每个栈求前缀和 $s$，其中 $s[w]$ 表示一个体积为 $w+1$，价值为 $s[w]$ 的物品。

问题转化成：

- 从 $n$ 个物品组中选物品，每组至多选一个物品（可以不选），要求体积总和至多为 $k$，求物品价值总和的最大值。

⚠**注意**：对于本题来说，由于元素值都是非负数，且一定可以选 $k$ 个硬币，所以「至多」和「恰好」计算出来的结果是一样的。为方便写代码这里用至多。

## 一、记忆化搜索

类似 0-1 背包，定义 $\textit{dfs}(i,j)$ 表示从 $\textit{piles}[0]$ 到 $\textit{piles}[i]$ 中，选体积之和至多为 $j$ 的物品时，物品价值之和的最大值。

枚举第 $i$ 组的所有物品（枚举前缀和），设当前物品体积为 $w$，价值为 $v$，那么问题变成从前 $i-1$ 个物品组中，选体积之和至多为 $j-w$ 的物品时，物品价值之和的最大值，即 $\textit{dfs}(i-1,j-w)$，加上 $v$ 得到 $\textit{dfs}(i,j)$。

所有情况取最大值，得

$$
\textit{dfs}(i,j) = \max_{(v,w)}\textit{dfs}(i-1,j-w) + v
$$

如果该组不选物品，则上式中的 $v=w=0$。

**递归边界**：$\textit{dfs}(-1,j)=0$。

**递归入口**：$\textit{dfs}(n-1,k)$，这是原问题，也是答案。

关于记忆化搜索的原理，请看 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

本题由于元素值都是正数，所以可以把 $\textit{memo}$ 数组初始化成 $0$。

```py [sol-Python3]
class Solution:
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i < 0:
                return 0
            # 不选这一组中的任何物品
            res = dfs(i - 1, j)
            # 枚举选哪个
            for w, v in enumerate(accumulate(piles[i][:j]), 1):
                res = max(res, dfs(i - 1, j - w) + v)
            return res
        return dfs(len(piles) - 1, k)
```

```java [sol-Java]
class Solution {
    public int maxValueOfCoins(List<List<Integer>> piles, int k) {
        int n = piles.size();
        int[][] memo = new int[n][k + 1];
        return dfs(n - 1, k, piles, memo);
    }

    private int dfs(int i, int j, List<List<Integer>> piles, int[][] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i][j] != 0) { // 之前计算过
            return memo[i][j];
        }
        // 不选这一组中的任何物品
        int res = dfs(i - 1, j, piles, memo);
        // 枚举选哪个
        int v = 0;
        for (int w = 0; w < Math.min(j, piles.get(i).size()); w++) {
            v += piles.get(i).get(w);
            // w 从 0 开始，物品体积为 w+1
            res = Math.max(res, dfs(i - 1, j - w - 1, piles, memo) + v);
        }
        return memo[i][j] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValueOfCoins(vector<vector<int>>& piles, int k) {
        int n = size(piles);
        vector memo(n, vector<int>(k + 1));
        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            if (i < 0) {
                return 0;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            // 不选这一组中的任何物品
            res = dfs(i - 1, j);
            // 枚举选哪个
            int v = 0;
            for (int w = 0; w < min(j, (int) piles[i].size()); w++) {
                v += piles[i][w];
                // w 从 0 开始，物品体积为 w+1
                res = max(res, dfs(i - 1, j - w - 1) + v);
            }
            return res;
        };
        return dfs(n - 1, k);
    }
};
```

```go [sol-Go]
func maxValueOfCoins(piles [][]int, k int) int {
    n := len(piles)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, k+1)
    }
    var dfs func(int, int) int
    dfs = func(i, j int) (res int) {
        if i < 0 {
            return
        }
        p := &memo[i][j]
        if *p != 0 { // 之前计算过
            return *p
        }
        defer func() { *p = res }() // 记忆化

        // 不选这一组中的任何物品
        res = dfs(i-1, j)
        // 枚举选哪个
        v := 0
        for w := range min(j, len(piles[i])) {
            v += piles[i][w]
            // w 从 0 开始，物品体积为 w+1
            res = max(res, dfs(i-1, j-w-1)+v)
        }
        return
    }
    return dfs(n-1, k)
}
```

## 二、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示从 $\textit{piles}[0]$ 到 $\textit{piles}[i]$ 中，选体积之和至多为 $j$ 的物品时，物品价值之和的最大值。这里 $+1$ 是为了把 $\textit{dfs}(-1,j)$ 这个状态也翻译过来，这样我们可以把 $f[0][j]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j] = \max_{(v,w)}f[i][j-w] + v
$$

如果该组不选物品，则上式中的 $v=w=0$。

初始值 $f[0][j]=0$，翻译自递归边界 $\textit{dfs}(-1,j)=0$。

答案为 $f[n][k]$，翻译自递归入口 $\textit{dfs}(n-1,k)$。

```py [sol-Python3]
class Solution:
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        f = [[0] * (k + 1) for _ in range(len(piles) + 1)]
        for i, pile in enumerate(piles):
            for j in range(k + 1):
                # 不选这一组中的任何物品
                f[i + 1][j] = f[i][j]
                # 枚举选哪个
                for w, v in enumerate(accumulate(pile[:j]), 1):
                    f[i + 1][j] = max(f[i + 1][j], f[i][j - w] + v)
        return f[-1][k]
```

```java [sol-Java]
class Solution {
    public int maxValueOfCoins(List<List<Integer>> piles, int k) {
        int[][] f = new int[piles.size() + 1][k + 1];
        for (int i = 0; i < piles.size(); i++) {
            List<Integer> pile = piles.get(i);
            for (int j = 0; j <= k; j++) {
                // 不选这一组中的任何物品
                f[i + 1][j] = f[i][j];
                // 枚举选哪个
                int v = 0;
                for (int w = 0; w < Math.min(j, pile.size()); w++) {
                    v += pile.get(w);
                    // w 从 0 开始，物品体积为 w+1
                    f[i + 1][j] = Math.max(f[i + 1][j], f[i][j - w - 1] + v);
                }
            }
        }
        return f[piles.size()][k];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValueOfCoins(vector<vector<int>>& piles, int k) {
        vector f(piles.size() + 1, vector<int>(k + 1));
        for (int i = 0; i < piles.size(); i++) {
            auto& pile = piles[i];
            for (int j = 0; j <= k; j++) {
                // 不选这一组中的任何物品
                f[i + 1][j] = f[i][j];
                // 枚举选哪个
                int v = 0;
                for (int w = 0; w < min(j, (int) pile.size()); w++) {
                    v += pile[w];
                    // w 从 0 开始，物品体积为 w+1
                    f[i + 1][j] = max(f[i + 1][j], f[i][j - w - 1] + v);
                }
            }
        }
        return f.back()[k];
    }
};
```

```go [sol-Go]
func maxValueOfCoins(piles [][]int, k int) int {
    f := make([][]int, len(piles)+1)
    for i := range f {
        f[i] = make([]int, k+1)
    }
    for i, pile := range piles {
        for j := range k + 1 {
            // 不选这一组中的任何物品
            f[i+1][j] = f[i][j]
            // 枚举选哪个
            v := 0
            for w := range min(j, len(pile)) {
                v += pile[w]
                // w 从 0 开始，物品体积为 w+1
                f[i+1][j] = max(f[i+1][j], f[i][j-w-1]+v)
            }
        }
    }
    return f[len(piles)][k]
}
```

## 三、空间优化

和 0-1 背包一样，把 $f$ 数组的第一维去掉，第二维倒序枚举。原理请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

优化后 $f[i + 1][j] = f[i][j]$ 变成 $f[j] = f[j]$，可以省略。

此外，循环次数也可以优化。想一想，如果 $\textit{piles}[0]$ 和 $\textit{piles}[1]$ 的大小之和只有 $5$，而 $k=100$，那么是否需要从 $j=100$ 开始倒序枚举呢？

```py [sol-Python3]
class Solution:
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        f = [0] * (k + 1)
        sum_n = 0
        for pile in piles:
            n = len(pile)
            for i in range(1, n):
                pile[i] += pile[i - 1]  # 提前计算 pile 的前缀和
            sum_n = min(sum_n + n, k)
            for j in range(sum_n, 0, -1):  # 优化：j 从前 i 个栈的大小之和开始枚举
                # w 从 0 开始，物品体积为 w+1
                f[j] = max(f[j], max(f[j - w - 1] + pile[w] for w in range(min(n, j))))
        return f[k]
```

```java [sol-Java]
class Solution {
    public int maxValueOfCoins(List<List<Integer>> piles, int k) {
        int[] f = new int[k + 1];
        int sumN = 0;
        for (List<Integer> Pile : piles) {
            Integer[] pile = Pile.toArray(Integer[]::new); // 转成数组处理更快更方便
            int n = pile.length;
            for (int i = 1; i < n; i++) {
                pile[i] += pile[i - 1]; // 提前计算 pile 的前缀和
            }
            sumN = Math.min(sumN + n, k);
            for (int j = sumN; j > 0; j--) { // 优化：j 从前 i 个栈的大小之和开始枚举
                for (int w = 0; w < Math.min(n, j); w++) {
                    f[j] = Math.max(f[j], f[j - w - 1] + pile[w]); // w 从 0 开始，物品体积为 w+1
                }
            }
        }
        return f[k];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValueOfCoins(vector<vector<int>>& piles, int k) {
        vector<int> f(k + 1);
        int sum_n = 0;
        for (auto& pile : piles) {
            partial_sum(pile.begin(), pile.end(), pile.begin()); // 提前计算 pile 的前缀和
            int n = pile.size();
            sum_n = min(sum_n + n, k);
            for (int j = sum_n; j; j--) { // 优化：j 从前 i 个栈的大小之和开始枚举
                for (int w = 0; w < min(n, j); w++) {
                    f[j] = max(f[j], f[j - w - 1] + pile[w]); // w 从 0 开始，物品体积为 w+1
                }
            }
        }
        return f[k];
    }
};
```

```go [sol-Go]
func maxValueOfCoins(piles [][]int, k int) int {
    f := make([]int, k+1)
    sumN := 0
    for _, pile := range piles {
        n := len(pile)
        for i := 1; i < n; i++ {
            pile[i] += pile[i-1] // 提前计算 pile 的前缀和
        }
        sumN = min(sumN+n, k)
        for j := sumN; j > 0; j-- { // 优化：j 从前 i 个栈的大小之和开始枚举
            for w, v := range pile[:min(n, j)] {
                f[j] = max(f[j], f[j-w-1]+v) // w 从 0 开始，物品体积为 w+1
            }
        }
    }
    return f[k]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(Lk)$。将外层循环与最内层循环合并，即为每个栈的大小之和，记作 $\textit{L}$。再算上中间这层 $\mathcal{O}(k)$ 的循环，时间复杂度为 $\mathcal{O}(Lk)$。
- 空间复杂度：$\mathcal{O}(k)$。

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
