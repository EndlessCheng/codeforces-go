## 前言

消除字符满足如下性质：

1. 可以消除相邻的「连续」字符。
2. 相邻字符消除后，原本不相邻的字符会变成相邻，可以继续消除。换句话说，设子串 $A = x + B + y$，如果 $x$ 和 $y$ 是「连续」字符，且子串 $B$ 可以完全消除，那么子串 $A$ 可以完全消除。
3. 设子串 $A = B + C$，如果子串 $B$ 和 $C$ 可以完全消除，那么子串 $A$ 可以完全消除。

> **注**：如果把一对「连续」字符看成一对左右**括号**，那么消除的子串一定是一个**合法括号字符串**。

满足上述性质的题目，可以用**区间 DP** 解决。不了解区间 DP 的同学可以看 [区间 DP【基础算法精讲 22】](https://www.bilibili.com/video/BV1Gs4y1E7EU/)。

## 区间 DP

怎么判断子串 $s[i]$ 到 $s[j]$ 可以完全消除？

根据性质 2，如果 $s[i]$ 和 $s[j]$ 是「连续」字符，且子串 $s[i+1]$ 到 $s[j-1]$ 可以完全消除，那么子串 $s[i]$ 到 $s[j]$ 可以完全消除。

根据性质 3，枚举子串 $B$ 的右端点，即枚举 $k=i+1,i+2,\ldots,j-2$，把子串 $s[i]$ 到 $s[j]$ 分割成子串 $s[i]$ 到 $s[k]$ 和子串 $s[k+1]$ 到 $s[j]$。如果子串 $s[i]$ 到 $s[k]$ 和子串 $s[k+1]$ 到 $s[j]$ 都可以完全消除，那么子串 $s[i]$ 到 $s[j]$ 可以完全消除。

这个过程可以用区间 DP 实现。

边界值：空串可以完全消除。

## 线性 DP

**性质**：设 $a$ 是字典序最小的字符串，那么去掉 $a[0]$，剩余的 $a[1:]$ 的字典序也是最小的。

**证明**：反证法。如果存在 $t < a[1:]$，那么 $a[0] + t$ 是字典序更小的字符串，矛盾。故原命题成立。

根据该性质，分类讨论：

- 如果 $s[0]$ 在答案中，那么答案等于 $s[0]$ 加上后缀 $s[1:]$ 中能得到的字典序最小字符串。
- 如果 $s[0]$ 不在答案中，且子串 $s[0]$ 到 $s[j]$ 可以完全消除，那么答案等于后缀 $s[j+1:]$ 中能得到的字典序最小字符串。

提取子问题：

- 后缀 $s[i:]$ 中能得到的字典序最小字符串，记作 $f[i]$。

用「选或不选」讨论：

- $f[i]$ 包含 $s[i]$。此时 $f[i]$ 等于 $s[i]$ 加上后缀 $s[i+1:]$ 中能得到的字典序最小字符串，即 $f[i+1]$。
- $f[i]$ 不包含 $s[i]$。前提条件是存在 $j$ 使得 $s[i]$ 到 $s[j]$ 可以完全消除，此时 $f[i]$ 等于后缀 $s[j+1:]$ 中能得到的字典序最小字符串，即 $f[j+1]$。我们可以枚举所有满足该条件的 $j$。

这些字符串取最小值，即为 $f[i]$。

初始值：$f[n]$ 等于空串。

答案：$f[0]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1o1jgzJE51/?t=39m31s)，欢迎点赞关注~

## 细节

由于每次消除 $2$ 个字符，所以被消除的子串的长度一定是偶数。这可以用来减少枚举次数（枚举步长是 $2$）。

## 写法一：记忆化搜索

```py [sol-Python3]
def is_consecutive(x: str, y: str) -> bool:
    d = abs(ord(x) - ord(y))
    return d == 1 or d == 25

class Solution:
    def lexicographicallySmallestString(self, s: str) -> str:
        n = len(s)

        @cache
        def can_be_empty(i: int, j: int) -> bool:
            if i > j:  # 空串
                return True
            # 性质 2
            if is_consecutive(s[i], s[j]) and can_be_empty(i + 1, j - 1):
                return True
            # 性质 3
            for k in range(i + 1, j - 1, 2):
                if can_be_empty(i, k) and can_be_empty(k + 1, j):
                    return True
            return False

        @cache
        def dfs(i: int) -> str:
            if i == n:
                return ""
            # 包含 s[i]
            res = s[i] + dfs(i + 1)
            # 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
            for j in range(i + 1, n, 2):
                if can_be_empty(i, j):  # 消除 s[i] 到 s[j]
                    res = min(res, dfs(j + 1))
            return res

        return dfs(0)
```

```java [sol-Java]
class Solution {
    private int[][] memoEmpty;
    private String[] memoDfs;
    private char[] s;

    public String lexicographicallySmallestString(String S) {
        s = S.toCharArray();
        int n = s.length;
        memoEmpty = new int[n][n];
        for (int[] row : memoEmpty) {
            Arrays.fill(row, -1); // -1 表示尚未计算
        }
        memoDfs = new String[n];
        return dfs(0);
    }

    private boolean isConsecutive(char x, char y) {
        int d = Math.abs(x - y);
        return d == 1 || d == 25;
    }

    private int canBeEmpty(int i, int j) {
        if (i > j) { // 空串
            return 1;
        }
        if (memoEmpty[i][j] != -1) {
            return memoEmpty[i][j];
        }

        // 性质 2
        if (isConsecutive(s[i], s[j]) && canBeEmpty(i + 1, j - 1) > 0) {
            return memoEmpty[i][j] = 1;
        }

        // 性质 3
        for (int k = i + 1; k < j; k += 2) {
            if (canBeEmpty(i, k) > 0 && canBeEmpty(k + 1, j) > 0) {
                return memoEmpty[i][j] = 1;
            }
        }

        return memoEmpty[i][j] = 0;
    }

    private String dfs(int i) {
        if (i == s.length) {
            return "";
        }
        if (memoDfs[i] != null) {
            return memoDfs[i];
        }

        // 包含 s[i]
        String res = s[i] + dfs(i + 1);

        // 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
        for (int j = i + 1; j < s.length; j += 2) {
            if (canBeEmpty(i, j) > 0) { // 消除 s[i] 到 s[j]
                String t = dfs(j + 1);
                if (t.compareTo(res) < 0) {
                    res = t;
                }
            }
        }

        return memoDfs[i] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_consecutive(char x, char y) {
        int d = abs(x - y);
        return d == 1 || d == 25;
    }

public:
    string lexicographicallySmallestString(string s) {
        int n = s.size();
        vector memo_empty(n, vector<int8_t>(n, -1));
        auto can_be_empty = [&](this auto&& can_be_empty, int i, int j) -> bool {
            if (i > j) { // 空串
                return true;
            }
            auto& res = memo_empty[i][j]; // 注意这是引用
            if (res != -1) {
                return res;
            }
            // 性质 2
            if (is_consecutive(s[i], s[j]) && can_be_empty(i + 1, j - 1)) {
                return res = true; // 记忆化
            }
            // 性质 3
            for (int k = i + 1; k < j; k += 2) {
                if (can_be_empty(i, k) && can_be_empty(k + 1, j)) {
                    return res = true;
                }
            }
            return res = false;
        };

        vector<string> memo_dfs(n, "?");
        auto dfs = [&](this auto&& dfs, int i) -> string {
            if (i == n) {
                return "";
            }
            auto& res = memo_dfs[i];
            if (res != "?") {
                return res;
            }
            // 包含 s[i]
            res = s[i] + dfs(i + 1);
            // 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
            for (int j = i + 1; j < n; j += 2) {
                if (can_be_empty(i, j)) { // 消除 s[i] 到 s[j]
                    res = min(res, dfs(j + 1));
                }
            }
            return res;
        };

        return dfs(0);
    }
};
```

```go [sol-Go]
func isConsecutive(x, y byte) bool {
	d := abs(int(x) - int(y))
	return d == 1 || d == 25
}

func lexicographicallySmallestString(s string) string {
	n := len(s)
	memoEmpty := make([][]int8, n)
	for i := range memoEmpty {
		memoEmpty[i] = make([]int8, n)
		for j := range memoEmpty[i] {
			memoEmpty[i][j] = -1
		}
	}

	var canBeEmpty func(int, int) int8
	canBeEmpty = func(i, j int) (res int8) {
		if i > j { // 空串
			return 1
		}
		p := &memoEmpty[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		// 性质 2
		if isConsecutive(s[i], s[j]) && canBeEmpty(i+1, j-1) > 0 {
			return 1
		}
		// 性质 3
		for k := i + 1; k < j; k += 2 {
			if canBeEmpty(i, k) > 0 && canBeEmpty(k+1, j) > 0 {
				return 1
			}
		}
		return 0
	}

	memoDfs := make([]string, n)
	for i := range memoDfs {
		memoDfs[i] = "?"
	}
	var dfs func(int) string
	dfs = func(i int) string {
		if i == n {
			return ""
		}
		p := &memoDfs[i]
		if *p != "?" {
			return *p
		}

		// 包含 s[i]
		res := string(s[i]) + dfs(i+1)
		// 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
		for j := i + 1; j < n; j += 2 {
			if canBeEmpty(i, j) > 0 { // 消除 s[i] 到 s[j]
				res = min(res, dfs(j+1))
			}
		}
		*p = res
		return res
	}

	return dfs(0)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

## 写法二：递推

```py [sol-Python3]
def is_consecutive(x: int, y: int) -> bool:
    d = abs(x - y)
    return d == 1 or d == 25

class Solution:
    def lexicographicallySmallestString(self, s: str) -> str:
        a = list(map(ord, s))  # 避免反复调用 ord
        n = len(s)

        can_be_empty = [[False] * n for _ in range(n)]
        for i in range(n - 2, -1, -1):
            can_be_empty[i + 1][i] = True  # 空串
            for j in range(i + 1, n, 2):
                # 性质 2
                if is_consecutive(a[i], a[j]) and can_be_empty[i + 1][j - 1]:
                    can_be_empty[i][j] = True
                    continue
                # 性质 3
                for k in range(i + 1, j - 1, 2):
                    if can_be_empty[i][k] and can_be_empty[k + 1][j]:
                        can_be_empty[i][j] = True
                        break

        f = [''] * (n + 1)
        for i in range(n - 1, -1, -1):
            # 包含 s[i]
            res = s[i] + f[i + 1]
            # 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
            for j in range(i + 1, n, 2):
                if can_be_empty[i][j]:  # 消除 s[i] 到 s[j]
                    res = min(res, f[j + 1])
            f[i] = res
        return f[0]
```

```java [sol-Java]
class Solution {
    public String lexicographicallySmallestString(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        boolean[][] canBeEmpty = new boolean[n][n];
        for (int i = n - 2; i >= 0; i--) {
            canBeEmpty[i + 1][i] = true; // 空串
            for (int j = i + 1; j < n; j += 2) {
                // 性质 2
                if (isConsecutive(s[i], s[j]) && canBeEmpty[i + 1][j - 1]) {
                    canBeEmpty[i][j] = true;
                    continue;
                }
                // 性质 3
                for (int k = i + 1; k < j - 1; k += 2) {
                    if (canBeEmpty[i][k] && canBeEmpty[k + 1][j]) {
                        canBeEmpty[i][j] = true;
                        break;
                    }
                }
            }
        }

        String[] f = new String[n + 1];
        f[n] = "";
        for (int i = n - 1; i >= 0; i--) {
            // 包含 s[i]
            String res = s[i] + f[i + 1];
            // 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
            for (int j = i + 1; j < n; j += 2) { // 枚举消除 s[i] 到 s[j]
                if (canBeEmpty[i][j] && f[j + 1].compareTo(res) < 0) {
                    res = f[j + 1];
                }
            }
            f[i] = res;
        }
        return f[0];
    }

    private boolean isConsecutive(char x, char y) {
        int d = Math.abs(x - y);
        return d == 1 || d == 25;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_consecutive(char x, char y) {
        int d = abs(x - y);
        return d == 1 || d == 25;
    }

public:
    string lexicographicallySmallestString(string s) {
        int n = s.size();
        vector can_be_empty(n, vector<uint8_t>(n));
        for (int i = n - 2; i >= 0; i--) {
            can_be_empty[i + 1][i] = true; // 空串
            for (int j = i + 1; j < n; j += 2) {
                // 性质 2
                if (is_consecutive(s[i], s[j]) && can_be_empty[i + 1][j - 1]) {
                    can_be_empty[i][j] = true;
                    continue;
                }
                // 性质 3
                for (int k = i + 1; k < j - 1; k += 2) {
                    if (can_be_empty[i][k] && can_be_empty[k + 1][j]) {
                        can_be_empty[i][j] = true;
                        break;
                    }
                }
            }
        }

        vector<string> f(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            // 包含 s[i]
            f[i] = s[i] + f[i + 1];
            // 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
            for (int j = i + 1; j < n; j += 2) {
                if (can_be_empty[i][j]) { // 消除 s[i] 到 s[j]
                    f[i] = min(f[i], f[j + 1]);
                }
            }
        }
        return f[0];
    }
};
```

```go [sol-Go]
func isConsecutive(x, y byte) bool {
	d := abs(int(x) - int(y))
	return d == 1 || d == 25
}

func lexicographicallySmallestString(s string) (ans string) {
	n := len(s)
	canBeEmpty := make([][]bool, n)
	for i := range canBeEmpty {
		canBeEmpty[i] = make([]bool, n)
	}
	for i := n - 2; i >= 0; i-- {
		canBeEmpty[i+1][i] = true // 空串
		for j := i + 1; j < n; j += 2 {
			// 性质 2
			if isConsecutive(s[i], s[j]) && canBeEmpty[i+1][j-1] {
				canBeEmpty[i][j] = true
				continue
			}
			// 性质 3
			for k := i + 1; k < j-1; k += 2 {
				if canBeEmpty[i][k] && canBeEmpty[k+1][j] {
					canBeEmpty[i][j] = true
					break
				}
			}
		}
	}

	f := make([]string, n+1)
	for i := n - 1; i >= 0; i-- {
		// 包含 s[i]
		res := string(s[i]) + f[i+1]
		// 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
		for j := i + 1; j < n; j += 2 {
			if canBeEmpty[i][j] { // 消除 s[i] 到 s[j]
				res = min(res, f[j+1])
			}
		}
		f[i] = res
	}
	return f[0]
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

更多相似题目，见下面动态规划题单的「**八、区间 DP**」。

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
