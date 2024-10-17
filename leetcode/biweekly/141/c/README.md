## 一、寻找子问题

思考方式同 [1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)。

考虑从右往左匹配子序列，分类讨论：

- 不选 $\textit{source}[n-1]$，问题变成 $\textit{source}$ 的前 $n-1$ 个字母和 $\textit{pattern}$ 的前 $m$ 个字母的子问题。
- 如果 $\textit{source}[n-1]=\textit{pattern}[m-1]$，那么匹配（都选），问题变成 $\textit{source}$ 的前 $n-1$ 个字母和 $\textit{pattern}$ 的前 $m-1$ 个字母的子问题。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

> 注 1：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。
>
> 注 2：动态规划有「**选或不选**」和「**枚举选哪个**」两种基本思考方式。子序列相邻无关一般是「选或不选」，子序列相邻相关（例如 LIS 问题）一般是「枚举选哪个」。本题属于子序列相邻无关问题，用「**选或不选**」解决。

## 二、状态定义与状态转移方程

根据上面的讨论，定义状态为 $\textit{dfs}(i,j)$，表示要使 $\textit{pattern}[0]$ 到 $\textit{pattern}[j]$ 是 $\textit{source}[0]$ 到 $\textit{source}[i]$ 的子序列，最多可以进行多少次删除操作。

分类讨论：

- 不选 $\textit{source}[i]$，问题变成要使 $\textit{pattern}[0]$ 到 $\textit{pattern}[j]$ 是 $\textit{source}[0]$ 到 $\textit{source}[i-1]$ 的子序列，最多可以进行多少次删除操作，即 $\textit{dfs}(i-1,j)$。如果 $i$ 在 $\textit{targetIndices}$ 中，那么删除次数加一。
- 如果 $\textit{source}[i]=\textit{pattern}[j]$，那么匹配（都选），问题变成要使 $\textit{pattern}[0]$ 到 $\textit{pattern}[j-1]$ 是 $\textit{source}[0]$ 到 $\textit{source}[i-1]$ 的子序列，最多可以进行多少次删除操作，即 $\textit{dfs}(i-1,j-1)$。

这两种情况取最大值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) =
\begin{cases}
\textit{dfs}(i-1,j) + [i\in \textit{targetIndices}], & j=-1     \\
\max(\textit{dfs}(i-1,j) + [i\in \textit{targetIndices}], \textit{dfs}(i-1,j-1)), & j \ge 0     \\
\end{cases}
$$

其中 $[P]$ 表示：如果 $P$ 成立，那么 $[P]=1$，否则 $[P]=0$。

**递归边界**：

- 如果 $i<j$，那么 $\textit{dfs}(i,j)=-\infty$。用 $-\infty$ 表示不合法的状态，从而保证 $\max$ 不会取到不合法的状态。
- 否则，$\textit{dfs}(-1,-1) = 0$，子序列匹配完毕。

**递归入口**：$\textit{dfs}(n-1,m-1)$，也就是答案。

代码实现时，可以把 $\textit{targetIndices}$ 转成哈希集合或者数组，从而快速判断 $i\in \textit{targetIndices}$ 是否成立。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

详细原理见视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1iR2zYaESG/)（第三题），欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxRemovals(self, source: str, pattern: str, targetIndices: List[int]) -> int:
        targetIndices = set(targetIndices)

        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i < j:
                return -inf
            if i < 0:
                return 0
            res = dfs(i - 1, j) + (i in targetIndices)
            if j >= 0 and source[i] == pattern[j]:
                res = max(res, dfs(i - 1, j - 1))
            return res

        ans = dfs(len(source) - 1, len(pattern) - 1)
        dfs.cache_clear()  # 防止爆内存
        return ans
```

```java [sol-Java]
class Solution {
    public int maxRemovals(String source, String pattern, int[] targetIndices) {
        Set<Integer> targetSet = new HashSet<>();
        for (int index : targetIndices) {
            targetSet.add(index);
        }
        int n = source.length();
        int m = pattern.length();
        int[][] memo = new int[n][m + 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, m - 1, source.toCharArray(), pattern.toCharArray(), targetSet, memo);
    }

    private int dfs(int i, int j, char[] source, char[] pattern, Set<Integer> targetSet, int[][] memo) {
        if (i < j) {
            return Integer.MIN_VALUE;
        }
        if (i < 0) {
            return 0;
        }
        // j+1 避免数组越界
        if (memo[i][j + 1] != -1) { // 之前计算过
            return memo[i][j + 1];
        }
        int res = dfs(i - 1, j, source, pattern, targetSet, memo) + (targetSet.contains(i) ? 1 : 0);
        if (j >= 0 && source[i] == pattern[j]) {
            res = Math.max(res, dfs(i - 1, j - 1, source, pattern, targetSet, memo));
        }
        return memo[i][j + 1] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxRemovals(string source, string pattern, vector<int>& targetIndices) {
        set<int> st(targetIndices.begin(), targetIndices.end());
        int n = source.length(), m = pattern.length();
        vector<vector<int>> memo(n, vector<int>(m + 1, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j) -> int {
            if (i < j) {
                return INT_MIN;
            }
            if (i < 0) {
                return 0;
            }
            int& res = memo[i][j + 1]; // 注意这里是引用，+1 避免数组越界
            if (res != -1) { // 之前计算过
                return res;
            }
            res = dfs(dfs, i - 1, j) + st.count(i);
            if (j >= 0 && source[i] == pattern[j]) {
                res = max(res, dfs(dfs, i - 1, j - 1));
            }
            return res;
        };
        return dfs(dfs, n - 1, m - 1);
    }
};
```

```go [sol-Go]
func maxRemovals(source, pattern string, targetIndices []int) int {
	targetSet := map[int]int{}
	for _, idx := range targetIndices {
		targetSet[idx] = 1
	}
	n, m := len(source), len(pattern)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < j {
			return math.MinInt
		}
		if i < 0 {
			return 0
		}
		p := &memo[i][j+1] // +1 避免数组越界
		if *p != -1 { // 之前计算过
			return *p
		}
		res := dfs(i-1, j) + targetSet[i]
		if j >= 0 && source[i] == pattern[j] {
			res = max(res, dfs(i-1, j-1))
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, m-1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 是 $\textit{pattern}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(nm)$。
- 空间复杂度：$\mathcal{O}(nm)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j+1]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示要使 $\textit{pattern}[0]$ 到 $\textit{pattern}[j]$ 是 $\textit{source}[0]$ 到 $\textit{source}[i]$ 的子序列，最多可以进行多少次删除操作。这里 $+1$ 是为了把 $\textit{dfs}(-1,-1)$ 这个状态也翻译过来，这样我们可以把 $f[0][0]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j+1] =
\begin{cases}
f[i][j+1] + [i\in \textit{targetIndices}], & j=-1     \\
\max(f[i][j+1] + [i\in \textit{targetIndices}], f[i][j]), & j \ge 0     \\
\end{cases}
$$

> 问：为什么 $\textit{source}[i]$ 的下标不用变？为什么不是 $\textit{source}[i+1]$？
>
> 答：如果写成 $\textit{source}[i+1]$，那么当 $i=n-1$ 时，$\textit{source}[i+1]$ 会下标越界，这显然是错误的。

初始值：

- 先把 $f$ 数组全部初始化成 $-\infty$。
- 然后置 $f[0][0]=0$，翻译自递归边界 $\textit{dfs}(-1,-1)=0$。

答案为 $f[n][m]$，翻译自递归入口 $\textit{dfs}(n-1,m-1)$。

```py [sol-Python3]
class Solution:
    def maxRemovals(self, source: str, pattern: str, targetIndices: List[int]) -> int:
        targetIndices = set(targetIndices)
        n, m = len(source), len(pattern)
        f = [[-inf] * (m + 1) for _ in range(n + 1)]
        f[0][0] = 0
        for i, x in enumerate(source):
            is_del = 1 if i in targetIndices else 0
            f[i + 1][0] = f[i][0] + is_del
            for j in range(min(i + 1, m)):
                res = f[i][j + 1] + is_del
                if x == pattern[j]:
                    res = max(res, f[i][j])
                f[i + 1][j + 1] = res
        return f[n][m]
```

```java [sol-Java]
class Solution {
    public int maxRemovals(String source, String pattern, int[] targetIndices) {
        Set<Integer> targetSet = new HashSet<>();
        for (int i : targetIndices) {
            targetSet.add(i);
        }

        char[] s = source.toCharArray();
        char[] p = pattern.toCharArray();
        int n = s.length;
        int m = p.length;

        int[][] f = new int[n + 1][m + 1];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        f[0][0] = 0;

        for (int i = 0; i < n; i++) {
            int isDel = targetSet.contains(i) ? 1 : 0;
            f[i + 1][0] = f[i][0] + isDel;
            for (int j = 0; j < Math.min(i + 1, m); j++) {
                int res = f[i][j + 1] + isDel;
                if (s[i] == p[j]) {
                    res = Math.max(res, f[i][j]);
                }
                f[i + 1][j + 1] = res;
            }
        }
        return f[n][m];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxRemovals(string source, string pattern, vector<int>& targetIndices) {
        unordered_set<int> st(targetIndices.begin(), targetIndices.end());
        int n = source.length(), m = pattern.length();
        vector<vector<int>> f(n + 1, vector<int>(m + 1, INT_MIN));
        f[0][0] = 0;
        for (int i = 0; i < n; i++) {
            int is_del = st.count(i);
            f[i + 1][0] = f[i][0] + is_del;
            for (int j = 0; j < min(i + 1, m); j++) {
                f[i + 1][j + 1] = f[i][j + 1] + is_del;
                if (source[i] == pattern[j]) {
                    f[i + 1][j + 1] = max(f[i + 1][j + 1], f[i][j]);
                }
            }
        }
        return f[n][m];
    }
};
```

```go [sol-Go]
func maxRemovals(source, pattern string, targetIndices []int) int {
	targetSet := map[int]int{}
	for _, i := range targetIndices {
		targetSet[i] = 1
	}

	n, m := len(source), len(pattern)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0

	for i := range source {
		isDel := targetSet[i]
		f[i+1][0] = f[i][0] + isDel
		for j := 0; j < min(i+1, m); j++ {
			res := f[i][j+1] + isDel
			if source[i] == pattern[j] {
				res = max(res, f[i][j])
			}
			f[i+1][j+1] = res
		}
	}
	return f[n][m]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 是 $\textit{pattern}$ 的长度。
- 空间复杂度：$\mathcal{O}(nm)$。

## 五、空间优化

### 1)

无需把 $\textit{targetIndices}$ 转成哈希集合或者数组。

由于 $\textit{targetIndices}$ 是有序数组，可以用**双指针**遍历。

### 2)

观察上面的状态转移方程，在计算 $f[i+1]$ 时，只会用到 $f[i]$，不会用到比 $i$ 更早的状态。

因此可以像 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/) 那样，去掉第一个维度，把 $f[i+1]$ 和 $f[i]$ 保存到**同一个数组**中。注意和 0-1 背包一样，$j$ 要**倒序**枚举。

状态转移方程改为

$$
f[j+1] =
\begin{cases}
f[j+1] + [i\in \textit{targetIndices}], & j=-1     \\
\max(f[j+1] + [i\in \textit{targetIndices}], f[j]), & j \ge 0     \\
\end{cases}
$$

初始值 $f[0]=0$，其余为 $-\infty$。

答案为 $f[m]$。

```py [sol-Python3]
class Solution:
    def maxRemovals(self, source: str, pattern: str, targetIndices: List[int]) -> int:
        m = len(pattern)
        f = [0] + [-inf] * m
        k = 0
        for i, x in enumerate(source):
            if k < len(targetIndices) and targetIndices[k] < i:
                k += 1
            is_del = 1 if k < len(targetIndices) and targetIndices[k] == i else 0
            for j in range(min(i, m - 1), -1, -1):
                f[j + 1] += is_del
                if x == pattern[j]:
                    f[j + 1] = max(f[j + 1], f[j])
            f[0] += is_del
        return f[m]
```

```py [sol-Python3 手动 max]
class Solution:
    def maxRemovals(self, source: str, pattern: str, targetIndices: List[int]) -> int:
        m = len(pattern)
        f = [0] + [-inf] * m
        k = 0
        for i, x in enumerate(source):
            if k < len(targetIndices) and targetIndices[k] < i:
                k += 1
            is_del = 1 if k < len(targetIndices) and targetIndices[k] == i else 0
            for j in range(min(i, m - 1), -1, -1):
                f[j + 1] += is_del
                if x == pattern[j] and f[j] > f[j + 1]:
                    f[j + 1] = f[j]
            f[0] += is_del
        return f[m]
```

```java [sol-Java]
class Solution {
    public int maxRemovals(String source, String pattern, int[] targetIndices) {
        char[] s = source.toCharArray();
        char[] p = pattern.toCharArray();

        int m = p.length;
        int[] f = new int[m + 1];
        Arrays.fill(f, Integer.MIN_VALUE);
        f[0] = 0;

        int k = 0;
        for (int i = 0; i < s.length; i++) {
            if (k < targetIndices.length && targetIndices[k] < i) {
                k++;
            }
            int is_del = k < targetIndices.length && targetIndices[k] == i ? 1 : 0;
            for (int j = Math.min(i, m - 1); j >= 0; j--) {
                f[j + 1] += is_del;
                if (s[i] == p[j]) {
                    f[j + 1] = Math.max(f[j + 1], f[j]);
                }
            }
            f[0] += is_del;
        }
        return f[m];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxRemovals(string source, string pattern, vector<int>& targetIndices) {
        int m = pattern.length();
        vector<int> f(m + 1, INT_MIN);
        f[0] = 0;
        int k = 0;
        for (int i = 0; i < source.length(); i++) {
            if (k < targetIndices.size() && targetIndices[k] < i) {
                k++;
            }
            int is_del = k < targetIndices.size() && targetIndices[k] == i;
            for (int j = min(i, m - 1); j >= 0; j--) {
                f[j + 1] += is_del;
                if (source[i] == pattern[j]) {
                    f[j + 1] = max(f[j + 1], f[j]);
                }
            }
            f[0] += is_del;
        }
        return f[m];
    }
};
```

```go [sol-Go]
func maxRemovals(source, pattern string, targetIndices []int) int {
	m := len(pattern)
	f := make([]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = math.MinInt
	}
	k := 0
	for i := range source {
		if k < len(targetIndices) && targetIndices[k] < i {
			k++
		}
		isDel := 0
		if k < len(targetIndices) && targetIndices[k] == i {
			isDel = 1
		}
		for j := min(i, m-1); j >= 0; j-- {
			f[j+1] += isDel
			if source[i] == pattern[j] {
				f[j+1] = max(f[j+1], f[j])
			}
		}
		f[0] += isDel
	}
	return f[m]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 是 $\textit{pattern}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

## 相似题目

见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**§4.1 最长公共子序列（LCS）**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
