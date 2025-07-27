## 前置知识：动态规划入门

请看视频：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)

## 预处理

预处理 $s$ 的每个【长度至少为 $2$ 的】子串修改成半回文串至少要修改多少次，记到数组 $\textit{modify}$ 中，$\textit{modify}[i][j]$ 对应从 $s[i]$ 到 $s[j]$ 的子串。

对于每个子串，枚举其因子 $d$（注意 $d$ 要小于子串长度）。例如子串 $t$ 的长度为 $6$，$d=2$ 时，我们需要判断

- $t[0]+ t[2]+ t[4]$ 组成的字符串，改成回文串需要修改多少个字母。
- $t[1]+ t[3]+ t[5]$ 组成的字符串，改成回文串需要修改多少个字母。

所有修改次数相加，就是因子为 $d$ 时的修改次数。取所有修改次数的最小值，就是这个子串的最小修改次数。

## 记忆化搜索

按照【划分型 DP】的套路，定义 $\textit{dfs}(i,j)$ 表示把 $s[0]$ 到 $s[j]$ 的字符串分成 $i+1$ 个子字符串，使得每个子字符串变成半回文串需要修改的最少字符数目。

枚举 $s[j]$ 结束的子串在 $s[L]$ 处开始，那么有

$$
\textit{dfs}(i,j) = \min_{L=2i}^{j-1} \textit{dfs}(i-1,L-1) + \textit{modify}[L][j]
$$

注意 $L$ 从 $2i$ 开始（为后面的子串预留空间），到 $j-1$ 结束（因为子串长度至少为 $2$）。

递归边界：$i=0$ 时，只剩下一个子串，返回 $\textit{modify}[0][j]$。

递归入口：$\textit{dfs}(k-1,n-1)$，即为答案。

[视频讲解](https://www.bilibili.com/video/BV12w411B7ia/) 第四题。

```py [sol-Python3]
# 预处理每个数的真因子
MX = 201
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i * 2, MX, i):
        divisors[j].append(i)

def get_modify(s: str) -> int:
    res = n = len(s)
    for d in divisors[n]:
        cnt = 0
        for i0 in range(d):
            i, j = i0, n - d + i0
            while i < j:
                cnt += s[i] != s[j]
                i += d
                j -= d
        res = min(res, cnt)
    return res

class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)
        modify = [[0] * n for _ in range(n - 1)]
        for left in range(n - 1):
            for right in range(left + 1, n):  # 半回文串长度至少为 2
                modify[left][right] = get_modify(s[left: right + 1])

        @cache
        def dfs(i: int, j: int) -> int:
            if i == 0:
                return modify[0][j]
            return min(dfs(i - 1, L - 1) + modify[L][j] for L in range(i * 2, j))
        return dfs(k - 1, n - 1)
```

```java [sol-Java]
class Solution {
    public int minimumChanges(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[][] modify = new int[n - 1][n];
        for (int left = 0; left < n - 1; left++) {
            for (int right = left + 1; right < n; right++) {
                modify[left][right] = getModify(s, left, right - left + 1);
            }
        }

        int[][] memo = new int[k][n];
        for (int i = 0; i < k; i++) {
            Arrays.fill(memo[i], -1); // -1 表示没有计算过
        }
        return dfs(k - 1, n - 1, memo, modify);
    }

    private static final int MX = 201;
    private static final List<Integer>[] divisors = new ArrayList[MX];

    static {
        // 预处理每个数的真因子
        Arrays.setAll(divisors, k -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i * 2; j < MX; j += i) {
                divisors[j].add(i);
            }
        }
    }

    private int getModify(char[] s, int begin, int n) {
        int res = n;
        for (int d : divisors[n]) {
            int cnt = 0;
            for (int i0 = 0; i0 < d; i0++) {
                for (int i = i0, j = n - d + i0; i < j; i += d, j -= d) {
                    if (s[begin + i] != s[begin + j]) {
                        cnt++;
                    }
                }
            }
            res = Math.min(res, cnt);
        }
        return res;
    }

    private int dfs(int i, int j, int[][] memo, int[][] modify) {
        if (i == 0) { // 递归边界
            return modify[0][j];
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        int res = Integer.MAX_VALUE;
        for (int L = i * 2; L < j; L++) {
            res = Math.min(res, dfs(i - 1, L - 1, memo, modify) + modify[L][j]);
        }
        return memo[i][j] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
// 预处理每个数的真因子
const int MX = 201;
vector<int> divisors[MX];

int init = [] {
    for (int i = 1; i < MX; i++) {
        for (int j = i * 2; j < MX; j += i) {
            divisors[j].push_back(i);
        }
    }
    return 0;
}();

class Solution {
    int get_modify(const string& s) {
        int n = s.length();
        int res = n;
        for (int d: divisors[n]) {
            int cnt = 0;
            for (int i0 = 0; i0 < d; i0++) {
                for (int i = i0, j = n - d + i0; i < j; i += d, j -= d) {
                    cnt += s[i] != s[j];
                }
            }
            res = min(res, cnt);
        }
        return res;
    }

public:
    int minimumChanges(string s, int k) {
        int n = s.length();
        vector modify(n - 1, vector<int>(n));
        for (int left = 0; left < n - 1; left++) {
            for (int right = left + 1; right < n; right++) {
                modify[left][right] = get_modify(s.substr(left, right - left + 1));
            }
        }

        vector memo(k, vector<int>(n, n + 1)); // n+1 表示没有计算过
        auto dfs = [&](this auto&& dfs, int i, int j) -> int {
            if (i == 0) {
                return modify[0][j];
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res <= n) { // 之前计算过
                return res;
            }
            for (int L = i * 2; L < j; L++) {
                res = min(res, dfs(i - 1, L - 1) + modify[L][j]);
            }
            return res;
        };
        return dfs(k - 1, n - 1);
    }
};
```

```go [sol-Go]
// 预处理每个数的真因子
const mx = 200
var divisors [mx + 1][]int

func init() {
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func calc(s string) int {
	n := len(s)
	res := n
	for _, d := range divisors[n] {
		cnt := 0
		for i0 := 0; i0 < d; i0++ {
			for i, j := i0, n-d+i0; i < j; i, j = i+d, j-d {
				if s[i] != s[j] {
					cnt++
				}
			}
		}
		res = min(res, cnt)
	}
	return res
}

func minimumChanges(s string, k int) (ans int) {
	n := len(s)
	modify := make([][]int, n-1)
	for l := range modify {
		modify[l] = make([]int, n)
		for r := l + 1; r < n; r++ { // 半回文串长度至少为 2
			modify[l][r] = calc(s[l : r+1])
		}
	}

	memo := make([][]int, k)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return modify[0][j]
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := n
		for L := i * 2; L < j; L++ {
			res = min(res, dfs(i-1, L-1)+modify[L][j])
		}
		*p = res
		return res
	}
	return dfs(k-1, n-1)
}
```

## 1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

做法：

- $\textit{dfs}$ 改成 $f$ 数组；
- 递归改成循环（每个参数都对应一层循环）；
- 递归边界改成 $f$ 数组的初始值。

> 相当于之前是用递归去计算每个状态，现在是（按照某种顺序）枚举并计算每个状态。

具体来说，$f[i][j]$ 的含义和 $\textit{dfs}(i,j)$ 的含义是一样的，都表示把 $s[0]$ 到 $s[j]$ 的字符串分成 $i+1$ 个子字符串，使得每个子字符串变成半回文串需要修改的最少字符数目。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 是一样的：

$$
f[i][j] = \min_{L=2i}^{j-1} f[i-1][L-1] + \textit{modify}[L][j]
$$

初始值 $f[0][j]=\textit{modify}[0][j]$，翻译自 $\textit{dfs}(0,j)=\textit{modify}[0][j]$。

答案为 $f[k-1][n-1]$，翻译自 $\textit{dfs}(k-1,n-1)$。

代码实现时，可以像 0-1 背包那样倒序循环 $j$，就可以只用一个一维 DP 数组了。

```py [sol-Python3]
MX = 201
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i * 2, MX, i):
        divisors[j].append(i)

def get_modify(s: str) -> int:
    res = n = len(s)
    for d in divisors[n]:
        cnt = 0
        for i0 in range(d):
            i, j = i0, n - d + i0
            while i < j:
                cnt += s[i] != s[j]
                i += d
                j -= d
        res = min(res, cnt)
    return res

class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)
        modify = [[0] * n for _ in range(n - 1)]
        for left in range(n - 1):
            for right in range(left + 1, n):
                modify[left][right] = get_modify(s[left: right + 1])

        f = modify[0]
        for i in range(1, k):
            for j in range(n - 1 - (k - 1 - i) * 2, i * 2, -1):  # 左右都要预留空间
                f[j] = min(f[L - 1] + modify[L][j] for L in range(i * 2, j))
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int minimumChanges(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[][] modify = new int[n - 1][n];
        for (int left = 0; left < n - 1; left++) {
            for (int right = left + 1; right < n; right++) {
                modify[left][right] = getModify(s, left, right - left + 1);
            }
        }

        int[] f = modify[0];
        for (int i = 1; i < k; i++) {
            for (int j = n - 1 - (k - 1 - i) * 2; j > i * 2; j--) { // 左右都要预留空间
                f[j] = Integer.MAX_VALUE;
                for (int L = i * 2; L < j; L++) {
                    f[j] = Math.min(f[j], f[L - 1] + modify[L][j]);
                }
            }
        }
        return f[n - 1];
    }

    private static final int MX = 201;
    private static final List<Integer>[] divisors = new ArrayList[MX];

    static {
        Arrays.setAll(divisors, k -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i * 2; j < MX; j += i) {
                divisors[j].add(i);
            }
        }
    }

    private int getModify(char[] s, int begin, int n) {
        int res = n;
        for (int d : divisors[n]) {
            int cnt = 0;
            for (int i0 = 0; i0 < d; i0++) {
                for (int i = i0, j = n - d + i0; i < j; i += d, j -= d) {
                    if (s[begin + i] != s[begin + j]) {
                        cnt++;
                    }
                }
            }
            res = Math.min(res, cnt);
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MX = 201;
vector<int> divisors[MX];

int init = [] {
    for (int i = 1; i < MX; i++) {
        for (int j = i * 2; j < MX; j += i) {
            divisors[j].push_back(i);
        }
    }
    return 0;
}();

class Solution {
    int get_modify(const string &s, int begin, int end) {
        int n = end - begin;
        int res = n;
        for (int d: divisors[n]) {
            int cnt = 0;
            for (int i0 = 0; i0 < d; i0++) {
                for (int i = begin + i0, j = end - d + i0; i < j; i += d, j -= d) {
                    cnt += s[i] != s[j];
                }
            }
            res = min(res, cnt);
        }
        return res;
    }

public:
    int minimumChanges(string s, int k) {
        int n = s.length();
        vector<vector<int>> modify(n - 1, vector<int>(n));
        for (int left = 0; left < n - 1; left++) {
            for (int right = left + 1; right < n; right++) {
                modify[left][right] = get_modify(s, left, right + 1);
            }
        }

        vector<int> f(modify[0]);
        for (int i = 1; i < k; i++) {
            for (int j = n - 1 - (k - 1 - i) * 2; j > i * 2; j--) { // 左右都要预留空间
                f[j] = INT_MAX;
                for (int L = i * 2; L < j; L++) {
                    f[j] = min(f[j], f[L - 1] + modify[L][j]);
                }
            }
        }
        return f[n - 1];
    }
};
```

```go [sol-Go]
const mx = 200
var divisors [mx + 1][]int

func init() {
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func calc(s string) int {
	n := len(s)
	res := n
	for _, d := range divisors[n] {
		cnt := 0
		for i0 := 0; i0 < d; i0++ {
			for i, j := i0, n-d+i0; i < j; i, j = i+d, j-d {
				if s[i] != s[j] {
					cnt++
				}
			}
		}
		res = min(res, cnt)
	}
	return res
}

func minimumChanges(s string, k int) (ans int) {
	n := len(s)
	modify := make([][]int, n-1)
	for l := range modify {
		modify[l] = make([]int, n)
		for r := l + 1; r < n; r++ {
			modify[l][r] = calc(s[l : r+1])
		}
	}

	f := modify[0]
	for i := 1; i < k; i++ {
		for j := n - 1 - (k-1-i)*2; j > i*2; j-- { // 左右都要预留空间
			f[j] = n
			for L := i * 2; L < j; L++ {
				f[j] = min(f[j], f[L-1]+modify[L][j])
			}
		}
	}
	return f[n-1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3\log n)$，其中 $n$ 为 $s$ 的长度。时间主要在预处理上，有 $\mathcal{O}(n^2)$ 个子串，平均每个子串有 $\mathcal{O}(\log n)$ 个因子，每个因子需要 $\mathcal{O}(n)$ 的时间计算修改次数。
- 空间复杂度：$\mathcal{O}(n^2)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
