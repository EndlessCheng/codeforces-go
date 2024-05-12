## 分析

为方便描述，将 $\textit{nums}$ 数组简记为 $a$，将下标 $[0,n-1]$ 的排列 $\textit{perm}$ 简记为 $p$。

例如 $p=[0,2,1]$，对于分数的计算式子，要相减的两个数的对应关系是：

- $p_0$ 对应 $a_{p_1}$。
- $p_1$ 对应 $a_{p_2}$。
- $p_2$ 对应 $a_{p_0}$。

把 $p$ 循环左移一位，得到 $p'=[2,1,0]$，我们在计算得分时，要相减的两个数的对应关系是：

- $p'_0$ 对应 $a_{p'_1}$，即原来的 $p_1$ 对应 $a_{p_2}$。
- $p'_1$ 对应 $a_{p'_2}$，即原来的 $p_2$ 对应 $a_{p_0}$。
- $p'_2$ 对应 $a_{p'_0}$，即原来的 $p_0$ 对应 $a_{p_1}$。

可以发现，在循环左移后，对应关系是不变的。换句话说，**循环左移后，分数是不变的**。

**推论**：题目要计算的字典序最小的排列 $p$，一定满足 $p_0 = 0$。

**证明**：反证法，假设 $p_0 > 0$，我们可以将 $p$ 不断循环左移，直到 $p_0 =0$。由于循环左移后分数不变，我们得到了一个字典序更小的答案，矛盾，故原命题成立。

## 一、寻找子问题

假设 $n=6$。

回顾枚举 [46. 全排列](https://leetcode.cn/problems/permutations/) 的思路：

- $p_0 = 0$。
- 枚举 $p_1$ 选哪个，它不能是之前选过的数 $0$，可以是 $1,2,3,4,5$，假设 $p_1 = 2$。
- 枚举 $p_2$ 选哪个，它不能是之前选过的数 $0,2$，可以是 $1,3,4,5$，假设 $p_2 = 4$。
- 枚举 $p_3$ 选哪个，它不能是之前选过的数 $0,2,4$，可以是 $1,3,5$，假设 $p_3 = 1$。
- 枚举 $p_4$ 选哪个，它不能是之前选过的数 $0,1,2,4$，可以是 $3,5$，假设 $p_4 = 5$。
- 枚举 $p_5$ 选哪个，它不能是之前选过的数 $0,1,2,4,5$，只能是 $3$，所以 $p_5 = 3$。
- 生成的排列为 $p = [0,2,4,1,5,3]$。

在这个过程中，我们需要知道：

- 哪些数不能选。这可以用一个集合 $S$ 存储选过的数。枚举 $p_i$ 的值，然后把 $p_i$ 加入集合 $S$。注意剩下能选的数变少了，要解决的问题规模也变小了。

除此以外，为了计算分数：

- 对于 $p_i$，我们还需要知道上一个数 $p_{i-1}$ 是多少。设 $j = p_{i-1}$。

在上面的例子中，对于 $p_4$ 而言，之前选过的数的集合 $S=\{0,1,2,4\}$，上一个选的数是 $j=1$，那么枚举：

- $p_4=3$，把 $|1-a_3|$ 加入分数，接下来要解决的问题是：在 $S=\{0,1,2,3,4\},\ j=3$ 的情况下，剩余数字能得到的最低分数。
- $p_4=5$，把 $|1-a_5|$ 加入分数，接下来要解决的问题是：在 $S=\{0,1,2,4,5\},\ j=5$ 的情况下，剩余数字能得到的最低分数。

由于这些都是**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「枚举选哪个」。

请注意上述过程中，会产生**重复子问题**，例如：

- 目前生成的排列是 $p = [0,2,4,1,\text{\_},\text{\_}]$，现在递归到倒数第二个位置，那么 $S=\{0,1,2,4\},\ j=1$。
- 目前生成的排列是 $p = [0,4,2,1,\text{\_},\text{\_}]$，现在递归到倒数第二个位置，那么 $S=\{0,1,2,4\},\ j=1$。

这样的重复子问题，是本题可以用 DP 优化的关键。换句话说，前面的排列具体长啥样，我们并不关心，**我们记录的是无序的集合，不是有序的列表**。

## 二、状态定义与状态转移方程

按照上面的讨论，定义 $\textit{dfs}(S,j)$ 表示在之前选过的数的集合为 $S$，上一个选的数是 $j$ 的情况下，剩余数字能得到的最低分数。

考虑当前数字选什么：

- 枚举 $k=1,2,\cdots,n-1$ 且 $k\notin S$。注意 $0$ 一定在 $S$ 中，无需枚举。
- 要解决的问题变成：在之前选过的数的集合为 $S \cup \{k\}$，上一个选的数是 $k$ 的情况下，剩余数字能得到的最低分数。这个分数加上 $|j-a_k|$，更新 $\textit{dfs}(S,j)$ 的最小值。

即

$$
\textit{dfs}(S,j) = \min_{k} \textit{dfs}(S \cup \{k\}, k) + |j - a_k|
$$

递归边界：$\textit{dfs}(U,j)=|j - a_0|$，其中 $U=\{0,1,2,\cdots n-1\}$。

递归入口：$\textit{dfs}(\{0\}, 0)$，也就是答案。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

由于有重复子问题，整个递归过程中会有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(S,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

原理请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含如何把记忆化搜索 1:1 翻译成递推的技巧。

代码实现时，用位运算实现集合操作，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

如何输出方案，请看我之前的一篇题解[【输出具体方案】从递归到递推，教你一步步思考动态规划！](https://leetcode.cn/problems/shortest-common-supersequence/solution/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-auy8z/)，以及本题的 [视频讲解](https://www.bilibili.com/video/BV1bx4y1i7rP/) 第四题。注意我们是从左往右构造 $p$，从小到大枚举 $k$ 的，这正好符合字典序的定义，所以按照上述计算规则，答案的字典序自然是最小的。

```py [sol-Python3]
class Solution:
    def findPermutation(self, a: List[int]) -> List[int]:
        n = len(a)
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(s: int, j: int) -> int:
            if s == (1 << n) - 1:
                # 所有位置都填完了，最后一个位置是下标 j
                return abs(j - a[0])
            res = inf
            # 枚举当前位置填下标 k
            for k in range(1, n):
                if s >> k & 1 == 0:  # k 之前没填过
                    res = min(res, dfs(s | 1 << k, k) + abs(j - a[k]))
            return res

        ans = []
        # 原理见上面贴的题解链接
        def make_ans(s: int, j: int) -> None:
            ans.append(j)
            if s == (1 << n) - 1:
                return
            final_res = dfs(s, j)
            for k in range(1, n):
                if s >> k & 1 == 0 and dfs(s | 1 << k, k) + abs(j - a[k]) == final_res:
                    make_ans(s | 1 << k, k)
                    break
        make_ans(1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] findPermutation(int[] a) {
        int n = a.length;
        int[][] memo = new int[1 << n][n];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        int[] ans = new int[n];
        makeAns(1, 0, a, memo, ans, 0);
        return ans;
    }

    private int dfs(int s, int j, int[] a, int[][] memo) {
        if (s == (1 << a.length) - 1) {
            return Math.abs(j - a[0]);
        }
        if (memo[s][j] != -1) { // 之前计算过
            return memo[s][j];
        }
        int res = Integer.MAX_VALUE;
        for (int k = 1; k < a.length; k++) {
            if ((s >> k & 1) == 0) { // k 之前没填过
                res = Math.min(res, dfs(s | 1 << k, k, a, memo) + Math.abs(j - a[k]));
            }
        }
        memo[s][j] = res; // 记忆化
        return res;
    }

    private void makeAns(int s, int j, int[] a, int[][] memo, int[] ans, int i) {
        ans[i] = j;
        if (s == (1 << a.length) - 1) {
            return;
        }
        int finalRes = dfs(s, j, a, memo);
        for (int k = 1; k < a.length; k++) {
            if ((s >> k & 1) == 0 && dfs(s | 1 << k, k, a, memo) + Math.abs(j - a[k]) == finalRes) {
                makeAns(s | 1 << k, k, a, memo, ans, i + 1);
                break;
            }
        }
    }
}   
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findPermutation(vector<int>& a) {
        int n = a.size();
        vector<vector<int>> memo(1 << n, vector<int>(n, -1)); // -1 表示没有计算过
        function<int(int, int)> dfs = [&](int s, int j) -> int {
            if (s == (1 << n) - 1) {
                return abs(j - a[0]);
            }
            int& res = memo[s][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = INT_MAX;
            for (int k = 1; k < n; k++) {
                if ((s >> k & 1) == 0) { // k 之前没填过
                    res = min(res, dfs(s | 1 << k, k) + abs(j - a[k]));
                }
            }
            return res;
        };

        vector<int> ans;
        function<void(int, int)> make_ans = [&](int s, int j) -> void {
            ans.push_back(j);
            if (s == (1 << n) - 1) {
                return;
            }
            int final_res = dfs(s, j);
            for (int k = 1; k < n; k++) {
                if ((s >> k & 1) == 0 && dfs(s | 1 << k, k) + abs(j - a[k]) == final_res) {
                    make_ans(s | 1 << k, k);
                    break;
                }
            }
        };
        make_ans(1, 0);
        return ans;
    }
};
```

```go [sol-Go]
func findPermutation(a []int) []int {
	n := len(a)
	memo := make([][]int, 1<<n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(s, j int) int {
		if s == 1<<n-1 {
			return abs(j - a[0])
		}
		p := &memo[s][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		for k := 1; k < n; k++ {
			if s>>k&1 == 0 { // k 之前没填过
				res = min(res, dfs(s|1<<k, k)+abs(j-a[k]))
			}
		}
		*p = res // 记忆化
		return res
	}

	ans := make([]int, 0, n)
	var makeAns func(int, int)
	makeAns = func(s, j int) {
		ans = append(ans, j)
		if s == 1<<n-1 {
			return
		}
		finalRes := dfs(s, j)
		for k := 1; k < n; k++ {
			if s>>k&1 == 0 && dfs(s|1<<k, k)+abs(j-a[k]) == finalRes {
				makeAns(s|1<<k, k)
				break
			}
		}
	}
	makeAns(1, 0)
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2 2^n)$，其中 $n$ 为 $\textit{a}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2 2^n)$。
- 空间复杂度：$\mathcal{O}(n2^n)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[S][j]$ 的定义和 $\textit{dfs}(S,j)$ 的定义是一样的，都表示在之前选过的数的集合为 $S$，上一个选的数是 $j$ 的情况下，剩余数字能得到的最低分数。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[S][j] = \min_{k} f[S \cup \{k\}][k] + |j - a_k|
$$

其中 $k=1,2,\cdots,n-1$ 且 $k\notin S$。

初始值 $f[U][j]=|j - a_0|$，翻译自递归边界 $\textit{dfs}(U,j)=|j - a_0|$。

答案为 $f[\{0\}][0]$，翻译自递归入口 $\textit{dfs}(\{0\}, 0)$。

为了输出方案，我们可以记录每个状态下填的数字，即 $\textit{g}[U][j] = k$。然后从 $j=0$ 开始，不断迭代 $j = \textit{g}[S][j]$ 来生成 $p$。

#### 答疑

**问**：如何思考循环顺序？什么时候要正序枚举，什么时候要倒序枚举？

**答**：这里有一个通用的做法：盯着状态转移方程，想一想，要计算 $f[S][j]$，必须先把 $f[S \cup \{k\}][\cdot]$ 算出来，那么只有 $S$ 从大到小枚举才能做到。

对于 $j$ 来说，由于在计算 $f[S][j]$ 的时候，$f[S \cup \{k\}][\cdot ]$ 已经全部计算完毕，所以 $j$ 无论是正序还是倒序枚举都可以。

```py [sol-Python3]
class Solution:
    def findPermutation(self, a: List[int]) -> List[int]:
        n = len(a)
        f = [[inf] * n for _ in range(1 << n)]
        g = [[-1] * n for _ in range(1 << n)]
        for j in range(n):
            f[-1][j] = abs(j - a[0])
        for s in range((1 << n) - 3, 0, -2):  # 注意偶数不含 0，是无效状态
            for j in range(n):
                if s >> j & 1 == 0:  # 无效状态，因为 j 一定在 s 中
                    continue
                for k in range(1, n):
                    if s >> k & 1:  # k 之前填过
                        continue
                    v = f[s | 1 << k][k] + abs(j - a[k])
                    if v < f[s][j]:
                        f[s][j] = v
                        g[s][j] = k  # 记录该状态下填了哪个数

        ans = []
        s = j = 0
        while j >= 0:
            ans.append(j)
            s |= 1 << j
            j = g[s][j]
        return ans
```

```java [sol-Java]
class Solution {
    public int[] findPermutation(int[] a) {
        int n = a.length;
        int[][] f = new int[1 << n][n];
        int[][] g = new int[1 << n][n];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MAX_VALUE);
        }
        for (int j = 0; j < n; j++) {
            f[(1 << n) - 1][j] = Math.abs(j - a[0]);
        }
        Arrays.fill(g[(1 << n) - 1], -1);
        for (int s = (1 << n) - 3; s > 0; s -= 2) { // 注意偶数不含 0，是无效状态
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0) { // 无效状态，因为 j 一定在 s 中
                    continue;
                }
                for (int k = 1; k < n; k++) {
                    if ((s >> k & 1) > 0) { // k 之前填过
                        continue;
                    }
                    int v = f[s | 1 << k][k] + Math.abs(j - a[k]);
                    if (v < f[s][j]) {
                        f[s][j] = v;
                        g[s][j] = k; // 记录该状态下填了哪个数
                    }
                }
            }
        }

        int[] ans = new int[n];
        int s = 0, i = 0, j = 0;
        while (j >= 0) {
            ans[i++] = j;
            s |= 1 << j;
            j = g[s][j];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findPermutation(vector<int>& a) {
        int n = a.size();
        vector<vector<int>> f(1 << n, vector<int>(n, INT_MAX));
        vector<vector<int>> g(1 << n, vector<int>(n, -1));
        for (int j = 0; j < n; j++) {
            f[(1 << n) - 1][j] = abs(j - a[0]);
        }
        for (int s = (1 << n) - 3; s > 0; s -= 2) { // 注意偶数不含 0，是无效状态
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0) { // 无效状态，因为 j 一定在 s 中
                    continue;
                }
                for (int k = 1; k < n; k++) {
                    if (s >> k & 1) { // k 之前填过
                        continue;
                    }
                    int v = f[s | 1 << k][k] + abs(j - a[k]);
                    if (v < f[s][j]) {
                        f[s][j] = v;
                        g[s][j] = k; // 记录该状态下填了哪个数
                    }
                }
            }
        }

        vector<int> ans;
        int s = 0, j = 0;
        while (j >= 0) {
            ans.push_back(j);
            s |= 1 << j;
            j = g[s][j];
        }
        return ans;
    }
};
```

```go [sol-Go]
func findPermutation(a []int) []int {
	n := len(a)
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	g := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
		g[i] = make([]int, n)
	}
	for j := range f[u] {
		f[u][j] = abs(j - a[0])
		g[u][j] = -1
	}
	for s := u - 2; s > 0; s -= 2 { // 注意偶数不含 0，是无效状态
		for j := 0; j < n; j++ {
			if s>>j&1 == 0 { // 无效状态，因为 j 一定在 s 中
				continue
			}
			for k := 1; k < n; k++ {
				if s>>k&1 > 0 { // k 之前填过
					continue
				}
				v := f[s|1<<k][k] + abs(j-a[k])
				if v < f[s][j] {
					f[s][j] = v
					g[s][j] = k // 记录该状态下填了哪个数
				}
			}
		}
	}

	ans := make([]int, 0, n)
	for s, j := 0, 0; j >= 0; {
		ans = append(ans, j)
		s |= 1 << j
		j = g[s][j]
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```go [sol-Go 更快的写法]
func findPermutation(a []int) []int {
	n := len(a)
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	g := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
		g[i] = make([]int, n)
	}
	for j := range f[u] {
		f[u][j] = abs(j - a[0])
		g[u][j] = -1
	}
	for s := u - 2; s > 0; s -= 2 { // 注意偶数不含 0，是无效状态
		for _s := uint(s); _s > 0; _s &= _s - 1 {
			j := bits.TrailingZeros(_s)
			for cus, lb := u^s, 0; cus > 0; cus ^= lb {
				lb = cus & -cus
				k := bits.TrailingZeros(uint(lb))
				v := f[s|lb][k] + abs(j-a[k])
				if v < f[s][j] {
					f[s][j] = v
					g[s][j] = k
				}
			}
		}
	}

	ans := make([]int, 0, n)
	for s, j := 0, 0; j >= 0; {
		ans = append(ans, j)
		s |= 1 << j
		j = g[s][j]
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2 2^n)$，其中 $n$ 为 $\textit{a}$ 的长度。
- 空间复杂度：$\mathcal{O}(n2^n)$。

## 相似题目

见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「§9.2 排列型 ② 相邻相关」。如果觉得本题太难，推荐先从较为容易「§9.1 排列型 ① 相邻无关」开始。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
