## 一、寻找子问题

方便把递归翻译成递推，**从右往左**思考。

举几个例子：

- 如果 $i=n-1$ 时，我们选择召唤火龙，且对手召唤了地精，那么我们获得 $1$ 分。问题变成从 $n-2$ 到 $0$，在我们得到 $1$ 分，对手得到 $0$ 分且 $i=n-2$ 无法召唤火龙的前提下，战胜对手的不同出招序列的数量。
- 如果 $i=n-1$ 时，我们选择召唤地精，且对手召唤了火龙，那么对手获得 $1$ 分。问题变成从 $n-2$ 到 $0$，在我们得到 $0$ 分，对手得到 $1$ 分且 $i=n-2$ 无法召唤地精的前提下，战胜对手的不同出招序列的数量。
- 如果 $i=n-1$ 时，我们选择召唤火龙，且对手召唤了火龙，那么双方都不会获得分数。问题变成从 $n-2$ 到 $0$，在我们得到 $0$ 分，对手得到 $0$ 分且 $i=n-2$ 无法召唤火龙的前提下，战胜对手的不同出招序列的数量。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

## 二、状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前讨论到第 $i$ 回合（$i$ 是 $s$ 的下标）。此时对手召唤的的生物为 $s[i]$。
- $j$：我们与对手的得分之差。注意无需分别维护我们与对手的得分。
- $\textit{ban}$：前一回合（$i+1$）我们召唤的生物。这也表示当前回合（$i$）我们无法召唤的生物。

因此，定义状态为 $\textit{dfs}(i,j,\textit{ban})$，表示从 $i$ 到 $0$，在我们与对手的得分之差为 $j$ 且第 $i$ 回合我们无法召唤的生物为 $\textit{ban}$ 的前提下，战胜对手的不同出招序列的数量。

接下来，思考如何从一个状态转移到另一个状态。

把 $\texttt{F},\texttt{W},\texttt{E}$ 分别记作 $0,1,2$。

枚举当前回合召唤的生物为 $k=0,1,2$ 且 $k\ne \textit{ban}$：

- 计算 $\textit{score} = (k-s[i]+3)\bmod 3$，如果 $\textit{score}=2$，则改成 $\textit{score}=-1$。
- 问题变成从 $i-1$ 到 $0$，在我们与对手的得分之差为 $j+\textit{score}$ 且第 $i-1$ 回合我们无法召唤的生物为 $k$ 的前提下，战胜对手的不同出招序列的数量，即 $\textit{dfs}(i-1,j+\textit{score},k)$。

累加 $\textit{dfs}(i-1,j+\textit{score},k)$ 得到 $\textit{dfs}(i,j,\textit{ban})$。

**递归边界**：

- 如果 $-j > i$，即使后面全胜，也无法战胜对手，返回 $0$。
- 如果 $j > i+1$，即使后面全败，也一定能战胜对手。由于剩余 $i+1$ 个回合，每个回合在两个可以召唤的生物中随意选择，所以方案数为 $2^{i+1}$，返回 $2^{i+1}$。

**递归入口**：$\textit{dfs}(n-1,0,-1)$，也就是答案。一开始得分之差为 $0$，没有召唤生物的限制。

也可以把 $\textit{dfs}(n-1,0,0)$ 作为递归入口，代码中通过判断 `i == n - 1` 避免讨论 `k == ban` 的情况。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j,\textit{ban})$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

```py [sol-Python3]
class Solution:
    def countWinningSequences(self, s: str) -> int:
        MOD = 1_000_000_007
        mp = {c: i for i, c in enumerate("FWE")}
        s = [mp[c] for c in s]

        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, ban: int) -> int:
            if -j > i:  # 必败
                return 0
            if j > i + 1:  # 必胜
                return pow(2, i + 1, MOD)
            res = 0
            for k in range(3):  # 枚举当前召唤的生物
                if k == ban:  # 不能连续两个回合召唤相同的生物
                    continue
                score = (k - s[i]) % 3
                if score == 2:
                    score = -1
                res += dfs(i - 1, j + score, k)
            return res % MOD
        return dfs(len(s) - 1, 0, -1)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int[] MAP = new int[128];
    private static final int[] POW2 = new int[500];

    static {
        MAP['F'] = 0;
        MAP['W'] = 1;
        MAP['E'] = 2;

        POW2[0] = 1;
        for (int i = 1; i < POW2.length; i++) {
            POW2[i] = POW2[i - 1] * 2 % MOD;
        }
    }

    public int countWinningSequences(String s) {
        int n = s.length();
        int[][][] memo = new int[n][n * 2 + 1][3];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1); // -1 表示没有计算过
            }
        }
        return dfs(n - 1, 0, 0, s.toCharArray(), memo);
    }

    private int dfs(int i, int j, int ban, char[] s, int[][][] memo) {
        if (-j > i) { // 必败
            return 0;
        }
        if (j > i + 1) { // 必胜
            return POW2[i + 1];
        }
        int n = s.length;
        if (memo[i][j + n][ban] != -1) { // 之前计算过
            return memo[i][j + n][ban];
        }
        int res = 0;
        for (int k = 0; k < 3; k++) { // 枚举当前召唤的生物
            // 判断 i == n - 1，避免讨论 k == ban 的情况
            if (i == n - 1 || k != ban) { // 不能连续两个回合召唤相同的生物
                int score = (k - MAP[s[i]] + 3) % 3;
                if (score == 2) {
                    score = -1;
                }
                res = (res + dfs(i - 1, j + score, k, s, memo)) % MOD;
            }
        }
        return memo[i][j + n][ban] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countWinningSequences(string s) {
        const int MOD = 1'000'000'007;
        int mp[128];
        mp['F'] = 0;
        mp['W'] = 1;
        mp['E'] = 2;

        int n = s.length();
        vector<int> pow2((n + 1) / 2);
        pow2[0] = 1;
        for (int i = 1; i < pow2.size(); i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }
        
        vector<vector<array<int, 3>>> memo(n, vector<array<int, 3>>(n * 2 + 1, {-1, -1, -1}));
        auto dfs = [&](auto&& dfs, int i, int j, int ban) -> int {
            if (-j > i) { // 必败
                return 0;
            }
            if (j > i + 1) { // 必胜
                return pow2[i + 1];
            }
            int& res = memo[i][j + n][ban]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            res = 0;
            for (int k = 0; k < 3; k++) { // 枚举当前召唤的生物
                // 判断 i == n - 1，避免讨论 k == ban 的情况
                if (i == n - 1 || k != ban) { // 不能连续两个回合召唤相同的生物
                    int score = (k - mp[s[i]] + 3) % 3;
                    if (score == 2) {
                        score = -1;
                    }
                    res = (res + dfs(dfs, i - 1, j + score, k)) % MOD;
                }
            }
            return res;
        };
        return dfs(dfs, n - 1, 0, 0); // ban=0,1,2 都可以
    }
};
```

```go [sol-Go]
func countWinningSequences(s string) int {
	const mod = 1_000_000_007
	n := len(s)
	pow2 := make([]int, (n+1)/2)
	pow2[0] = 1
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	mp := [...]int{'F': 0, 'W': 1, 'E': 2}
	memo := make([][][3]int, n)
	for i := range memo {
		memo[i] = make([][3]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = [3]int{-1, -1, -1} // -1 表示没有计算过
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, ban int) (res int) {
		if -j > i { // 必败
			return
		}
		if j > i+1 { // 必胜
			return pow2[i+1]
		}
		p := &memo[i][j+n][ban]
		if *p != -1 { // 之前计算过
			return *p
		}
		for k := 0; k < 3; k++ { // 枚举当前召唤的生物
			// 判断 i == n-1，避免讨论 k == ban 的情况
			if i == n-1 || k != ban { // 不能连续两个回合召唤相同的生物
				score := (k - mp[s[i]] + 3) % 3
				if score == 2 {
					score = -1
				}
				res += dfs(i-1, j+score, k)
			}
		}
		res %= mod
		*p = res // 记忆化
		return
	}
	return dfs(n-1, 0, 0) // ban=0,1,2 都可以
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2K^2)$，其中 $n$ 为 $s$ 的长度，$K=3$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2K)$，单个状态的计算时间为 $\mathcal{O}(K)$，所以总的时间复杂度为 $\mathcal{O}(n^2K^2)$。
- 空间复杂度：$\mathcal{O}(n^2K)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j+n][\textit{ban}]$ 的定义和 $\textit{dfs}(i,j,\textit{ban})$ 的定义是一样的，都表示从 $i$ 到 $0$，在我们与对手的得分之差为 $j$ 且第 $i$ 回合我们无法召唤的生物为 $\textit{ban}$ 的前提下，战胜对手的不同出招序列的数量。这里 $+1$ 和 $+n$ 是为了避免出现负数。

枚举当前回合召唤的生物为 $k=0,1,2$ 且 $k\ne \textit{ban}$：

- 计算 $\textit{score} = (k-s[i]+3)\bmod 3$，如果 $\textit{score}=2$，则改成 $\textit{score}=-1$。
- 问题变成从 $i-1$ 到 $0$，在我们与对手的得分之差为 $j+\textit{score}$ 且第 $i-1$ 回合我们无法召唤的生物为 $k$ 的前提下，战胜对手的不同出招序列的数量，即 $f[i][j+\textit{score}+n][k]$。

累加 $f[i][j+\textit{score}+n][k]$ 得到 $f[i+1][j+n][\textit{ban}]$。

循环时，$j$ 的范围为 $[-i, n-i-1]$：

- $j\ge -i$，这来自递归边界，因为 $-j > i$ 即 $j < -i$ 时状态值一定是 $0$，无需计算。
- $j\le n-i-1$，讨论记忆化搜索，全胜时有 $i+j=n-1$，即 $j=n-i-1$，这是 $j$ 的上限。计算 $j> n-i-1$ 的状态是没有意义的，因为用不到。

初始值：

- $f[0][j][\textit{ban}]=0$，其中 $j\le n$。
- $f[0][j][\textit{ban}]=1$，其中 $j > n$。
- $f[i+1][j+n][\textit{ban}]=2^{i+1}$，其中 $j >i+1$。

答案为 $f[n][n][0]$，翻译自递归入口 $\textit{dfs}(n-1,0,0)$。

```py [sol-Python3]
class Solution:
    def countWinningSequences(self, s: str) -> int:
        MOD = 1_000_000_007
        n = len(s)
        f = [[[0] * 3 for _ in range(n * 2 + 1)] for _ in range(n + 1)]
        for j in range(n + 1, n * 2 + 1):
            f[0][j] = [1] * 3

        mp = {'F': 0, 'W': 1, 'E': 2}
        pow2 = 1
        for i, c in enumerate(s):
            x = mp[c]
            pow2 = pow2 * 2 % MOD
            for j in range(-i, n - i):
                for ban in range(3):
                    if j > i + 1:
                        f[i + 1][j + n][ban] = pow2
                        continue
                    res = 0
                    for k in range(3):
                        if i == n - 1 or k != ban:
                            score = (k - x) % 3
                            if score == 2:
                                score = -1
                            res += f[i][j + score + n][k]
                    f[i + 1][j + n][ban] = res % MOD
        return f[n][n][0]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int[] MAP = new int[128];

    static {
        MAP['F'] = 0;
        MAP['W'] = 1;
        MAP['E'] = 2;
    }

    public int countWinningSequences(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[][][] f = new int[n + 1][n * 2 + 1][3];
        for (int j = n + 1; j <= n * 2; j++) {
            Arrays.fill(f[0][j], 1);
        }
        int pow2 = 1;
        for (int i = 0; i < n; i++) {
            int x = MAP[s[i]];
            pow2 = pow2 * 2 % MOD;
            for (int j = -i; j < n - i; j++) {
                for (int ban = 0; ban < 3; ban++) {
                    if (j > i + 1) {
                        f[i + 1][j + n][ban] = pow2;
                        continue;
                    }
                    int res = 0;
                    for (int k = 0; k < 3; k++) {
                        if (i == n - 1 || k != ban) {
                            int score = (k - x + 3) % 3;
                            if (score == 2) {
                                score = -1;
                            }
                            res = (res + f[i][j + score + n][k]) % MOD;
                        }
                    }
                    f[i + 1][j + n][ban] = res;
                }
            }
        }
        return f[n][n][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countWinningSequences(string s) {
        const int MOD = 1'000'000'007;
        int mp[128];
        mp['F'] = 0;
        mp['W'] = 1;
        mp['E'] = 2;

        int n = s.size();
        vector<vector<array<int, 3>>> f(n + 1, vector<array<int, 3>>(n * 2 + 1));
        for (int j = n + 1; j <= n * 2; j++) {
            f[0][j] = {1, 1, 1};
        }

        int pow2 = 1;
        for (int i = 0; i < n; i++) {
            int x = mp[s[i]];
            pow2 = pow2 * 2 % MOD;
            for (int j = -i; j < n - i; j++) {
                for (int ban = 0; ban < 3; ban++) {
                    if (j > i + 1) {
                        f[i + 1][j + n][ban] = pow2;
                        continue;
                    }
                    int& res = f[i + 1][j + n][ban]; // 注意这里是引用
                    for (int k = 0; k < 3; k++) {
                        if (i == n - 1 || k != ban) {
                            int score = (k - x + 3) % 3;
                            if (score == 2) {
                                score = -1;
                            }
                            res = (res + f[i][j + score + n][k]) % MOD;
                        }
                    }
                }
            }
        }
        return f[n][n][0];
    }
};
```

```go [sol-Go]
func countWinningSequences(s string) int {
	const mod = 1_000_000_007
	mp := [...]int{'F': 0, 'W': 1, 'E': 2}
	n := len(s)
	f := make([][][3]int, n+1)
	for i := range f {
		f[i] = make([][3]int, n*2+1)
	}
	for j := n + 1; j <= n*2; j++ {
		f[0][j] = [3]int{1, 1, 1}
	}
	pow2 := 1
	for i, c := range s {
		pow2 = pow2 * 2 % mod
		for j := -i; j < n-i; j++ {
			for pre := 0; pre < 3; pre++ {
				if j > i+1 {
					f[i+1][j+n][pre] = pow2
					continue
				}
				res := 0
				for cur := 0; cur < 3; cur++ {
					if i == n-1 || cur != pre {
						score := (cur - mp[c] + 3) % 3
						if score == 2 {
							score = -1
						}
						res += f[i][j+score+n][cur]
					}
				}
				f[i+1][j+n][pre] = res % mod
			}
		}
	}
	return f[n][n][0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2K^2)$，其中 $n$ 为 $s$ 的长度，$K=3$。
- 空间复杂度：$\mathcal{O}(n^2K)$。

注：利用滚动数组，可以把空间复杂度优化到 $\mathcal{O}(nK)$。

## 相似题目

见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**五、状态机 DP**」和「**§7.5 多维 DP**」。

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
