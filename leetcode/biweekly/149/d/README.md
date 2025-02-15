## 一、寻找子问题

为方便描述，下文把字符串 $\textit{caption}$ 简称为 $s$。

首先，如果 $s$ 的长度小于 $3$，无解，返回空字符串。下文讨论 $s$ 长度至少为 $3$ 的情况。

我们要解决的问题（原问题）是：

- 把 $s$ 的每个连续相同段的长度都变成 $\ge 3$ 的最少操作次数。

为了让答案的字典序最小，从左到右思考。

枚举 $s[0]$ 变成哪个字母。比如变成 $\texttt{a}$，那么：

- 如果 $s[1]$ 也是 $\texttt{a}$，那么问题变成：在 $s[1]=\texttt{a}$ 的前提下，$s[1]$ 到 $s[n-1]$ 的最少操作次数。
- 什么时候可以枚举其他字母？如果 $s[1]$ 和 $s[2]$ 也是 $\texttt{a}$，就可以保证 $s[0]$ 一定在一个长度至少为 $3$ 的连续相同子串中，这样 $s[3]$ 就可以枚举其他字母 $k$ 了，问题变成：在 $s[3]=k$ 的前提下，$s[3]$ 到 $s[n-1]$ 的最少操作次数。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

## 二、状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：剩余子串从 $s[i]$ 到 $s[n-1]$。
- $j$：规定 $s[i]$ 变成字母 $j$。

因此，定义状态为 $\textit{dfs}(i,j)$，表示在 $s[i]=j$ 的前提下，$s[i]$ 到 $s[n-1]$ 的最少操作次数。

分类讨论：

- $s[i+1]$ 也是 $j$，问题变成：在 $s[i+1]=j$ 的前提下，$s[i+1]$ 到 $s[n-1]$ 的最少操作次数，即 $\textit{dfs}(i+1,j)$。
- $s[i+1]$ 和 $s[i+2]$ 也是 $j$，但 $s[i+3]$ 是字母 $k$，问题变成：在 $s[i+3]=k$ 的前提下，$s[i+3]$ 到 $s[n-1]$ 的最少操作次数，即 $\textit{dfs}(i+3,k)$。注意这要求 $s[i+3]$ 到 $s[n-1]$ 的长度至少是 $3$
  ，也就是 $n-(i+3)\ge 3$，即 $i\le n-6$。

这两种情况取最小值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \min\left(\textit{dfs}(i+1,j) + |s[i]-j|,\ \min_{k=0}^{25} \textit{dfs}(i+3,k) + |s[i]-j| + |s[i+1]-j| + |s[i+2]-j| \right)
$$

注意无需判断 $k\ne j$，这不会得到比 $\textit{dfs}(i+1,j)$ 更优的答案。

**递归边界**：$\textit{dfs}(n,j)=0$。

**递归入口**：$\textit{dfs}(0,j)$。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$
到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1eUF6eaERQ/?t=8m14s)。

为方便大家理解，这里先给出计算最少操作次数的代码。输出具体方案在下文的递推中实现。

```py
# 只计算最小操作次数的代码
class Solution:
    def minCostGoodCaption(self, s: str) -> int:
        n = len(s)
        if n < 3:
            return -1

        s = [ord(c) - ord('a') for c in s]
        @cache
        def dfs(i: int, j: int) -> int:
            if i == n:
                return 0
            res = dfs(i + 1, j) + abs(s[i] - j)
            if i <= n - 6:
                mn = min(dfs(i + 3, k) for k in range(26))
                res = min(res, mn + abs(s[i] - j) + abs(s[i + 1] - j) + abs(s[i + 2] - j))
            return res
        return min(dfs(0, j) for j in range(26))
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$
  单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n|\Sigma|)$，单个状态的计算时间为 $\mathcal{O}(|\Sigma|)$，所以总的时间复杂度为 $\mathcal{O}(n|\Sigma|^2)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示在 $s[i]=j$ 的前提下，$s[i]$ 到 $s[n-1]$ 的最少操作次数。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \min\left(f[i+1][j] + |s[i]-j|,\ \min_{k=0}^{25} f[i+3][k] + |s[i]-j| + |s[i+1]-j| + |s[i+2]-j| \right)
$$

初始值 $f[n][j]=0$，翻译自递归边界 $\textit{dfs}(n,j)=0$。

先把上面的记忆化搜索 1:1 翻译过来，然后讨论优化和输出具体方案。

```py
# 只计算最小操作次数的代码
class Solution:
    def minCostGoodCaption(self, s: str) -> int:
        n = len(s)
        if n < 3:
            return -1

        s = [ord(c) - ord('a') for c in s]
        f = [[0] * 26 for _ in range(n + 1)]
        for i in range(n - 1, -1, -1):
            for j in range(26):
                res = f[i + 1][j] + abs(s[i] - j)
                res2 = min(f[i + 3]) + abs(s[i] - j) + abs(s[i + 1] - j) + abs(s[i + 2] - j) if i <= n - 6 else inf
                f[i][j] = min(res, res2)
        return min(f[0])
```

## 五、时间优化 + 输出具体方案

把 $\min\limits_{k=0}^{25} f[i][k]$ 保存到 $\textit{minF}[i]$ 中，于是转移方程优化成

$$
f[i][j] = \min\left(f[i+1][j] + |s[i]-j|,\ \textit{minF}[i+3] + |s[i]-j| + |s[i+1]-j| + |s[i+2]-j| \right)
$$

这样时间复杂度就优化至 $\mathcal{O}(n|\Sigma|)$ 了。

```py
# 只计算最小操作次数的代码
class Solution:
    def minCostGoodCaption(self, s: str) -> int:
        n = len(s)
        if n < 3:
            return -1

        s = [ord(c) - ord('a') for c in s]
        f = [[0] * 26 for _ in range(n + 1)]
        min_f = [0] * n
        for i in range(n - 1, -1, -1):
            for j in range(26):
                res = f[i + 1][j] + abs(s[i] - j)
                res2 = min_f[i + 3] + abs(s[i] - j) + abs(s[i + 1] - j) + abs(s[i + 2] - j) if i <= n - 6 else inf
                f[i][j] = min(res, res2)
            min_f[i] = min(f[i])
        return min_f[0]
```

本题还需要输出具体方案，这可以在递归的过程中，用一个 $\textit{nxt}$ 数组记录每个状态 $f[i][j]$ 的最优决策来自哪。

此外，我们还需要知道 $\textit{minF}[i]$ 等于哪个 $f[i][j]$，这里的 $j$ 是多少。这需要在计算最小值的过程中，把对应的 $j$ 保存到 $\textit{minJ}[i]$ 中。

于是，当上面代码中出现 $\textit{res}=\textit{res}_2$ 的情况时，需要比较 $\textit{minJ}[i+3]$ 和 $j$ 的大小关系，如果前者更小，那么要从 $\textit{res}_2$ 转移过来。

最终代码：

```py [sol-Python3]
class Solution:
    def minCostGoodCaption(self, s: str) -> str:
        n = len(s)
        if n < 3:
            return ""

        s = [ord(c) - ord('a') for c in s]
        f = [[0] * 26 for _ in range(n + 1)]
        min_j = [0] * (n + 1)
        nxt = [[0] * 26 for _ in range(n + 1)]
        for i in range(n - 1, -1, -1):
            mn = inf
            for j in range(26):
                res = f[i + 1][j] + abs(s[i] - j)
                res2 = f[i + 3][min_j[i + 3]] + abs(s[i] - j) + abs(s[i + 1] - j) + abs(s[i + 2] - j) if i <= n - 6 else inf
                if res2 < res or res2 == res and min_j[i + 3] < j:
                    res = res2
                    nxt[i][j] = min_j[i + 3]  # 记录转移来源
                else:
                    nxt[i][j] = j  # 记录转移来源
                f[i][j] = res
                if res < mn:
                    mn = res
                    min_j[i] = j  # 记录最小的 f[i][j] 中的 j 是多少

        ans = [''] * n
        i, j = 0, min_j[0]
        while i < n:
            ans[i] = ascii_lowercase[j]
            k = nxt[i][j]
            if k == j:
                i += 1
            else:
                ans[i + 2] = ans[i + 1] = ans[i]
                i += 3
                j = k
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String minCostGoodCaption(String S) {
        int n = S.length();
        if (n < 3) {
            return "";
        }

        char[] s = S.toCharArray();
        int[][] f = new int[n + 1][26];
        int[] minJ = new int[n + 1];
        int[][] nxt = new int[n + 1][26];
        for (int i = n - 1; i >= 0; i--) {
            int mn = Integer.MAX_VALUE;
            for (int j = 0; j < 26; j++) {
                int res = f[i + 1][j] + Math.abs(s[i] - 'a' - j);
                int res2 = i <= n - 6 ? f[i + 3][minJ[i + 3]] + Math.abs(s[i] - 'a' - j) + Math.abs(s[i + 1] - 'a' - j) + Math.abs(s[i + 2] - 'a' - j) : Integer.MAX_VALUE;
                if (res2 < res || res2 == res && minJ[i + 3] < j) {
                    res = res2;
                    nxt[i][j] = minJ[i + 3]; // 记录转移来源
                } else {
                    nxt[i][j] = j; // 记录转移来源
                }
                f[i][j] = res;
                if (res < mn) {
                    mn = res;
                    minJ[i] = j; // 记录最小的 f[i][j] 中的 j 是多少
                }
            }
        }

        char[] ans = new char[n];
        int i = 0;
        int j = minJ[0];
        while (i < n) {
            ans[i] = (char) ('a' + j);
            int k = nxt[i][j];
            if (k == j) {
                i++;
            } else {
                ans[i + 2] = ans[i + 1] = ans[i];
                i += 3;
                j = k;
            }
        }
        return new String(ans);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string minCostGoodCaption(string s) {
        int n = s.size();
        if (n < 3) {
            return "";
        }

        vector<array<int, 26>> f(n + 1);
        vector<int> min_j(n + 1);
        vector<array<int, 26>> nxt(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            int mn = INT_MAX;
            for (int j = 0; j < 26; j++) {
                int res = f[i + 1][j] + abs(s[i] - 'a' - j);
                int res2 = i <= n - 6 ? f[i + 3][min_j[i + 3]] + abs(s[i] - 'a' - j) + abs(s[i + 1] - 'a' - j) + abs(s[i + 2] - 'a' - j) : INT_MAX;
                if (res2 < res || res2 == res && min_j[i + 3] < j) {
                    res = res2;
                    nxt[i][j] = min_j[i + 3]; // 记录转移来源
                } else {
                    nxt[i][j] = j; // 记录转移来源
                }
                f[i][j] = res;
                if (res < mn) {
                    mn = res;
                    min_j[i] = j; // 记录最小的 f[i][j] 中的 j 是多少
                }
            }
        }

        string ans(n, 0);
        int i = 0, j = min_j[0];
        while (i < n) {
            ans[i] = 'a' + j;
            int k = nxt[i][j];
            if (k == j) {
                i++;
            } else {
                ans[i + 2] = ans[i + 1] = ans[i];
                i += 3;
                j = k;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCostGoodCaption(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([][26]int, n+1)
	minJ := make([]int, n+1)
	nxt := make([][26]int, n+1)
	for i := n - 1; i >= 0; i-- {
		mn := math.MaxInt
		for j := 0; j < 26; j++ {
			res := f[i+1][j] + abs(int(s[i]-'a')-j)
			res2 := math.MaxInt
			if i <= n-6 {
				res2 = f[i+3][minJ[i+3]] + abs(int(s[i]-'a')-j) + abs(int(s[i+1]-'a')-j) + abs(int(s[i+2]-'a')-j)
			}
			if res2 < res || res2 == res && minJ[i+3] < j {
				res = res2
				nxt[i][j] = minJ[i+3] // 记录转移来源
			} else {
				nxt[i][j] = j // 记录转移来源
			}
			f[i][j] = res
			if res < mn {
				mn = res
				minJ[i] = j // 记录最小的 f[i][j] 中的 j 是多少
			}
		}
	}

	ans := make([]byte, n)
	i, j := 0, minJ[0]
	for i < n {
		ans[i] = 'a' + byte(j)
		k := nxt[i][j]
		if k == j {
			i++
		} else {
			ans[i+1] = ans[i]
			ans[i+2] = ans[i]
			i += 3
			j = k
		}
	}
	return string(ans)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。

## 六、另一种思路：中位数贪心 + DP

注意到，对于长度 $\ge 6$ 的子串：

- 长为 $6$ 的子串可以拆分成两个长为 $3$ 的子串，这两个子串是互相独立的，可以分别计算最小修改次数。
- 长为 $7$ 的子串可以拆分成长为 $3$ 和 $4$ 的子串，两个子串分别计算最小修改次数。
- 长为 $8$ 的子串可以拆分成长为 $3$ 和 $5$，或者 $4$ 和 $4$ 的子串，两个子串分别计算最小修改次数。
- ……

所以本题的「基本元素」只有长为 $3,4,5$ 的子串，换句话说，问题相当于：

- 把 $s$ **划分**成若干长为 $3,4,5$ 的子串，把每个子串中的字母都变成相同的，求最小操作次数。

这可以用**划分型 DP** 解决。

定义 $f[i]$ 表示后缀 $s[i]$ 到 $s[n-1]$ 的最小操作次数。

枚举第一个子串的长度 $3,4,5$：

- 长为 $3$ 的子串，设其中字母分别为 $a,b,c$，不妨设 $a\le b\le c$，根据 [中位数贪心及其证明](https://leetcode.cn/problems/5TxKeK/solution/zhuan-huan-zhong-wei-shu-tan-xin-dui-din-7r9b/)，把所有数变成中位数 $b$ 是最优的，最小操作次数为 $(c-b)+(b-a)=c-a$。
- 长为 $4$ 的子串，设其中字母分别为 $a,b,c,d$，不妨设 $a\le b\le c\le d$，同理可得最小操作次数为 $c+d-a-b$，都变成 $b$。注意变成 $b$ 到 $c$ 中的字母操作次数都是最小的，由于本题要求答案字典序最小，所以变成 $b$。
- 长为 $5$ 的子串，设其中字母分别为 $a,b,c,d,e$，不妨设 $a\le b\le c\le d\le e$，同理可得最小操作次数为 $d+e-a-b$，都变成 $c$。

那么有

$$
f[i] = \min(f[i+3] + \textit{cost}_3, f[i+4] + \textit{cost}_4, f[i+5] + \textit{cost}_5)
$$

其中 $\textit{cost}_j$ 对应上面长为 $j$ 的子串的最小操作次数。

初始值 $f[n] = 0,\ f[n-1] = f[n-2] = \infty$。

```py
# 只计算最小操作次数的代码
class Solution:
    def minCostGoodCaption(self, s: str) -> int:
        n = len(s)
        if n < 3:
            return -1

        s = list(map(ord, s))
        f = [0] * (n + 1)
        f[n - 1] = f[n - 2] = inf
        for i in range(n - 3, -1, -1):
            a, _, c = sorted(s[i: i + 3])
            f[i] = f[i + 3] + c - a
            if i + 4 <= n:
                a, b, c, d = sorted(s[i: i + 4])
                f[i] = min(f[i], f[i + 4] + c + d - a - b)
            if i + 5 <= n:
                a, b, _, d, e = sorted(s[i: i + 5])
                f[i] = min(f[i], f[i + 5] + d + e - a - b)
        return f[0]
```

为了输出具体方案，需要额外比较字典序的大小。

为此，额外定义 $t_i$ 表示 $s_i$ 要变成的字母。

仍然枚举第一个子串的长度 $3,4,5$：

- 长为 $3$ 的子串（设字母分别为 $a,b,c$，已排序），算上 $s_{i+3}$ 要变成的字母 $t_{i+3}$，那么前 $6$ 个字母分别为 $b,b,b,t_{i+3},t_{i+3},t_{i+3}$。
- 长为 $4$ 的子串（设字母分别为 $a,b,c,d$，已排序），算上 $s_{i+4}$ 要变成的字母 $t_{i+4}$，那么前 $6$ 个字母分别为 $b,b,b,b,t_{i+4},t_{i+4}$。
- 长为 $5$ 的子串（设字母分别为 $a,b,c,d,e$，已排序），算上 $s_{i+5}$ 要变成的字母 $t_{i+5}$，那么前 $6$ 个字母分别为 $c,c,c,c,c,t_{i+5}$。

取前 $6$ 个字母的最小字典序，即为最终的转移来源。

> **问**：为什么不考虑第 $7$ 个字母？
> 
> **答**：这是因为，如果前 $6$ 个字母都一样，那么可以拆分成两个长为 $3$ 的子串，所以可以去掉前 $3$ 个字母，$f[i]$ 的最优决策就等同于 $f[i+3]$ 的最优决策。对于 $f[i]$ 而言，第 $7$ 个字母的最优值，就是 $f[i+3]$ 的第 $4$ 个字母的最优值，这已经在 $f[i+3]$ 中计算好了。

由于前 $3$ 个字母都是一样的，我们可以只比较第 $3$ 个到第 $6$ 个字母的字典序。 

### 优化前

```py [sol-Python3]
class Solution:
    def minCostGoodCaption(self, s: str) -> str:
        n = len(s)
        if n < 3:
            return ""

        s = list(map(ord, s))
        f = [0] * (n + 1)
        f[n - 1] = f[n - 2] = inf
        t = [0] * (n + 1)
        size = [3] * n

        for i in range(n - 3, -1, -1):
            a, b, c = sorted(s[i: i + 3])
            s3 = t[i + 3]
            res = (f[i + 3] + c - a, b, s3, s3, s3)

            if i + 4 <= n:
                a, b, c, d = sorted(s[i: i + 4])
                s4 = t[i + 4]
                tp = (f[i + 4] + c + d - a - b, b, b, s4, s4)
                if tp < res:
                    res = tp
                    size[i] = 4

            if i + 5 <= n:
                a, b, c, d, e = sorted(s[i: i + 5])
                tp = (f[i + 5] + d + e - a - b, c, c, c, t[i + 5])
                if tp < res:
                    res = tp
                    size[i] = 5

            f[i] = res[0]
            t[i] = res[1]

        ans = []
        i = 0
        while i < n:
            ans.append(chr(t[i]) * size[i])
            i += size[i]
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String minCostGoodCaption(String s) {
        int n = s.length();
        if (n < 3) {
            return "";
        }

        int[] f = new int[n + 1];
        f[n - 1] = f[n - 2] = Integer.MAX_VALUE / 2;
        char[] t = new char[n + 1];
        byte[] size = new byte[n];

        for (int i = n - 3; i >= 0; i--) {
            char[] sub = s.substring(i, i + 3).toCharArray();
            Arrays.sort(sub);
            char a = sub[0], b = sub[1], c = sub[2];
            char s3 = t[i + 3];
            int[] res = {f[i + 3] + (c - a), b, s3, s3, s3};
            size[i] = 3;

            if (i + 4 <= n) {
                char[] sub4 = s.substring(i, i + 4).toCharArray();
                Arrays.sort(sub4);
                char a4 = sub4[0], b4 = sub4[1], c4 = sub4[2], d4 = sub4[3];
                char s4 = t[i + 4];
                int[] tp = {f[i + 4] + (c4 - a4 + d4 - b4), b4, b4, s4, s4};
                if (less(tp, res)) {
                    res = tp;
                    size[i] = 4;
                }
            }

            if (i + 5 <= n) {
                char[] sub5 = s.substring(i, i + 5).toCharArray();
                Arrays.sort(sub5);
                char a5 = sub5[0], b5 = sub5[1], c5 = sub5[2], d5 = sub5[3], e5 = sub5[4];
                int[] tp = {f[i + 5] + (d5 - a5 + e5 - b5), c5, c5, c5, t[i + 5]};
                if (less(tp, res)) {
                    res = tp;
                    size[i] = 5;
                }
            }

            f[i] = res[0];
            t[i] = (char) res[1];
        }

        StringBuilder ans = new StringBuilder(n); // 预分配空间
        for (int i = 0; i < n; i += size[i]) {
            for (int j = 0; j < size[i]; j++) {
                ans.append(t[i]);
            }
        }
        return ans.toString();
    }

    private boolean less(int[] a, int[] b) {
        for (int i = 0; i < a.length; i++) {
            if (a[i] != b[i]) {
                return a[i] < b[i];
            }
        }
        return false;
    }
}
```

```java [sol-Java21]
class Solution {
    public String minCostGoodCaption(String s) {
        int n = s.length();
        if (n < 3) {
            return "";
        }

        int[] f = new int[n + 1];
        f[n - 1] = f[n - 2] = Integer.MAX_VALUE / 2;
        char[] t = new char[n + 1];
        byte[] size = new byte[n];

        for (int i = n - 3; i >= 0; i--) {
            char[] sub = s.substring(i, i + 3).toCharArray();
            Arrays.sort(sub);
            char a = sub[0], b = sub[1], c = sub[2];
            char s3 = t[i + 3];
            int[] res = {f[i + 3] + (c - a), b, s3, s3, s3};
            size[i] = 3;

            if (i + 4 <= n) {
                char[] sub4 = s.substring(i, i + 4).toCharArray();
                Arrays.sort(sub4);
                char a4 = sub4[0], b4 = sub4[1], c4 = sub4[2], d4 = sub4[3];
                char s4 = t[i + 4];
                int[] tp = {f[i + 4] + (c4 - a4 + d4 - b4), b4, b4, s4, s4};
                if (Arrays.compare(tp, res) < 0) {
                    res = tp;
                    size[i] = 4;
                }
            }

            if (i + 5 <= n) {
                char[] sub5 = s.substring(i, i + 5).toCharArray();
                Arrays.sort(sub5);
                char a5 = sub5[0], b5 = sub5[1], c5 = sub5[2], d5 = sub5[3], e5 = sub5[4];
                int[] tp = {f[i + 5] + (d5 - a5 + e5 - b5), c5, c5, c5, t[i + 5]};
                if (Arrays.compare(tp, res) < 0) {
                    res = tp;
                    size[i] = 5;
                }
            }

            f[i] = res[0];
            t[i] = (char) res[1];
        }

        StringBuilder ans = new StringBuilder(n); // 预分配空间
        for (int i = 0; i < n; i += size[i]) {
            ans.repeat(t[i], size[i]);
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string minCostGoodCaption(string s) {
        int n = s.size();
        if (n < 3) {
            return "";
        }

        vector<int> f(n + 1);
        f[n - 1] = f[n - 2] = INT_MAX / 2;
        vector<char> t(n + 1);
        vector<uint8_t> size(n);

        for (int i = n - 3; i >= 0; i--) {
            string sub = s.substr(i, 3);
            ranges::sort(sub);
            char a = sub[0], b = sub[1], c = sub[2];
            char s3 = t[i + 3];
            array<int, 5> res = {f[i + 3] + (c - a), b, s3, s3, s3};
            size[i] = 3;

            if (i + 4 <= n) {
                string sub = s.substr(i, 4);
                ranges::sort(sub);
                char a = sub[0], b = sub[1], c = sub[2], d = sub[3];
                char s4 = t[i + 4];
                array<int, 5> tp = {f[i + 4] + (c - a + d - b), b, b, s4, s4};
                if (tp < res) {
                    res = tp;
                    size[i] = 4;
                }
            }

            if (i + 5 <= n) {
                string sub = s.substr(i, 5);
                ranges::sort(sub);
                char a = sub[0], b = sub[1], c = sub[2], d = sub[3], e = sub[4];
                array<int, 5> tp = {f[i + 5] + (d - a + e - b), c, c, c, t[i + 5]};
                if (tp < res) {
                    res = tp;
                    size[i] = 5;
                }
            }

            f[i] = res[0];
            t[i] = res[1];
        }

        string ans;
        for (int i = 0; i < n; i += size[i]) {
            ans.append(size[i], t[i]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCostGoodCaption(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([]int, n+1)
	f[n-1], f[n-2] = math.MaxInt/2, math.MaxInt/2
	t := make([]byte, n+1)
	size := make([]uint8, n)

	for i := n - 3; i >= 0; i-- {
		sub := []byte(s[i : i+3])
		slices.Sort(sub)
		a, b, c := sub[0], sub[1], sub[2]
		s3 := int(t[i+3])
		res := []int{f[i+3] + int(c-a), int(b), s3, s3, s3}
		size[i] = 3

		if i+4 <= n {
			sub := []byte(s[i : i+4])
			slices.Sort(sub)
			a, b, c, d := sub[0], sub[1], sub[2], sub[3]
			s4 := int(t[i+4])
			tp := []int{f[i+4] + int(c-a+d-b), int(b), int(b), s4, s4}
			if slices.Compare(tp, res) < 0 {
				res = tp
				size[i] = 4
			}
		}

		if i+5 <= n {
			sub := []byte(s[i : i+5])
			slices.Sort(sub)
			a, b, c, d, e := sub[0], sub[1], sub[2], sub[3], sub[4]
			tp := []int{f[i+5] + int(d-a+e-b), int(c), int(c), int(c), int(t[i+5])}
			if slices.Compare(tp, res) < 0 {
				res = tp
				size[i] = 5
			}
		}

		f[i] = res[0]
		t[i] = byte(res[1])
	}

	ans := make([]byte, 0, n)
	for i := 0; i < n; i += int(size[i]) {
		ans = append(ans, bytes.Repeat([]byte{t[i]}, int(size[i]))...)
	}
	return string(ans)
}
```

### 优化

把四个字母压缩到一个 $\texttt{int}$ 中，这样可以快速比较字典序。

```py [sol-Python3]
# 注：对于 Python 来说这可能是负优化
class Solution:
    def minCostGoodCaption(self, s: str) -> str:
        n = len(s)
        if n < 3:
            return ""

        s = list(map(ord, s))
        f = [0] * (n + 1)
        f[n - 1] = f[n - 2] = inf
        t = [0] * (n + 1)
        size = [3] * n

        for i in range(n - 3, -1, -1):
            a, b, c = sorted(s[i: i + 3])
            s3 = t[i + 3]
            res = f[i + 3] + c - a
            mask = b << 24 | s3 << 16 | s3 << 8 | s3  # 4 个字母压缩成一个 int，方便比较字典序

            if i + 4 <= n:
                a, b, c, d = sorted(s[i: i + 4])
                s4 = t[i + 4]
                res4 = f[i + 4] + c + d - a - b
                mask4 = b << 24 | b << 16 | s4 << 8 | s4
                if res4 < res or res4 == res and mask4 < mask:
                    res, mask = res4, mask4
                    size[i] = 4

            if i + 5 <= n:
                a, b, c, d, e = sorted(s[i: i + 5])
                res5 = f[i + 5] + d + e - a - b
                mask5 = c << 24 | c << 16 | c << 8 | t[i + 5]
                if res5 < res or res5 == res and mask5 < mask:
                    res, mask = res5, mask5
                    size[i] = 5

            f[i] = res
            t[i] = mask >> 24

        ans = []
        i = 0
        while i < n:
            ans.append(chr(t[i]) * size[i])
            i += size[i]
        return ''.join(ans)
```

```java [sol-Java]
import static java.nio.charset.StandardCharsets.ISO_8859_1;

class Solution {
    public String minCostGoodCaption(String S) {
        int n = S.length();
        if (n < 3) {
            return "";
        }

        byte[] s = S.getBytes(ISO_8859_1);
        int[] f = new int[n + 1];
        f[n - 1] = f[n - 2] = Integer.MAX_VALUE / 2;
        byte[] t = new byte[n + 1];
        byte[] size = new byte[n];

        for (int i = n - 3; i >= 0; i--) {
            byte[] sub = Arrays.copyOfRange(s, i, i + 3); // 效率更高
            Arrays.sort(sub);
            byte a = sub[0], b = sub[1], c = sub[2];
            byte s3 = t[i + 3];
            int res = f[i + 3] + (c - a);
            int mask = b << 24 | s3 << 16 | s3 << 8 | s3; // 4 个 byte 压缩成一个 int，方便比较字典序
            size[i] = 3;

            if (i + 4 <= n) {
                byte[] sub4 = Arrays.copyOfRange(s, i, i + 4);
                Arrays.sort(sub4);
                byte a4 = sub4[0], b4 = sub4[1], c4 = sub4[2], d4 = sub4[3];
                byte s4 = t[i + 4];
                int res4 = f[i + 4] + (c4 - a4 + d4 - b4);
                int mask4 = b4 << 24 | b4 << 16 | s4 << 8 | s4;
                if (res4 < res || res4 == res && mask4 < mask) {
                    res = res4;
                    mask = mask4;
                    size[i] = 4;
                }
            }

            if (i + 5 <= n) {
                byte[] sub5 = Arrays.copyOfRange(s, i, i + 5);
                Arrays.sort(sub5);
                byte a5 = sub5[0], b5 = sub5[1], c5 = sub5[2], d5 = sub5[3], e5 = sub5[4];
                int res5 = f[i + 5] + (d5 - a5 + e5 - b5);
                int mask5 = c5 << 24 | c5 << 16 | c5 << 8 | t[i + 5];
                if (res5 < res || res5 == res && mask5 < mask) {
                    res = res5;
                    mask = mask5;
                    size[i] = 5;
                }
            }

            f[i] = res;
            t[i] = (byte) (mask >> 24);
        }

        StringBuilder ans = new StringBuilder(n); // 预分配空间
        for (int i = 0; i < n; i += size[i]) {
            // Java21 可以简化成 ans.repeat(t[i], size[i]);
            for (int j = 0; j < size[i]; j++) {
                ans.append((char) t[i]);
            }
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string minCostGoodCaption(string s) {
        int n = s.size();
        if (n < 3) {
            return "";
        }

        vector<int> f(n + 1);
        f[n - 1] = f[n - 2] = INT_MAX / 2;
        vector<char> t(n + 1);
        vector<uint8_t> size(n);

        for (int i = n - 3; i >= 0; i--) {
            string sub = s.substr(i, 3);
            ranges::sort(sub);
            char a = sub[0], b = sub[1], c = sub[2];
            char s3 = t[i + 3];
            int res = f[i + 3] + (c - a);
            int mask = b << 24 | s3 << 16 | s3 << 8 | s3; // 4 个 char 压缩成一个 int，方便比较字典序
            size[i] = 3;

            if (i + 4 <= n) {
                string sub = s.substr(i, 4);
                ranges::sort(sub);
                char a = sub[0], b = sub[1], c = sub[2], d = sub[3];
                char s4 = t[i + 4];
                int res4 = f[i + 4] + (c - a + d - b);
                int mask4 = b << 24 | b << 16 | s4 << 8 | s4;
                if (res4 < res || res4 == res && mask4 < mask) {
                    res = res4;
                    mask = mask4;
                    size[i] = 4;
                }
            }

            if (i + 5 <= n) {
                string sub = s.substr(i, 5);
                ranges::sort(sub);
                char a = sub[0], b = sub[1], c = sub[2], d = sub[3], e = sub[4];
                int res5 = f[i + 5] + (d - a + e - b);
                int mask5 = c << 24 | c << 16 | c << 8 | t[i + 5];
                if (res5 < res || res5 == res && mask5 < mask) {
                    res = res5;
                    mask = mask5;
                    size[i] = 5;
                }
            }

            f[i] = res;
            t[i] = mask >> 24;
        }

        string ans;
        for (int i = 0; i < n; i += size[i]) {
            ans.append(size[i], t[i]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCostGoodCaption(s string) string {
	n := len(s)
	if n < 3 {
		return ""
	}

	f := make([]int, n+1)
	f[n-1], f[n-2] = math.MaxInt/2, math.MaxInt/2
	t := make([]byte, n+1)
	size := make([]uint8, n)

	for i := n - 3; i >= 0; i-- {
		sub := []byte(s[i : i+3])
		slices.Sort(sub)
		a, b, c := sub[0], sub[1], sub[2]
		s3 := int(t[i+3])
		res := f[i+3] + int(c-a)
		mask := int(b)<<24 | s3<<16 | s3<<8 | s3 // 4 个 byte 压缩成一个 int，方便比较字典序
		size[i] = 3

		if i+4 <= n {
			sub := []byte(s[i : i+4])
			slices.Sort(sub)
			a, b, c, d := sub[0], sub[1], sub[2], sub[3]
			s4 := int(t[i+4])
			res4 := f[i+4] + int(c-a+d-b)
			mask4 := int(b)<<24 | int(b)<<16 | s4<<8 | s4
			if res4 < res || res4 == res && mask4 < mask {
				res, mask = res4, mask4
				size[i] = 4
			}
		}

		if i+5 <= n {
			sub := []byte(s[i : i+5])
			slices.Sort(sub)
			a, b, c, d, e := sub[0], sub[1], sub[2], sub[3], sub[4]
			res5 := f[i+5] + int(d-a+e-b)
			mask5 := int(c)<<24 | int(c)<<16 | int(c)<<8 | int(t[i+5])
			if res5 < res || res5 == res && mask5 < mask {
				res, mask = res5, mask5
				size[i] = 5
			}
		}

		f[i] = res
		t[i] = byte(mask >> 24)
	}

	ans := make([]byte, 0, n)
	for i := 0; i < n; i += int(size[i]) {
		ans = append(ans, bytes.Repeat([]byte{t[i]}, int(size[i]))...)
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。这个算法也可以理解成是一个 $\mathcal{O}(nk^2)$ 或者 $\mathcal{O}(nk^2\log k)$ 的算法，其中 $k=3$。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**五、状态机 DP**」「**§6.2 最优划分**」「**§7.5 多维 DP**」和「**专题：输出具体方案**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
