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

- $f[i]$ 包含 $s[i]$，那么 $f[i]$ 等于 $s[i]$ 加上后缀 $s[i+1:]$ 中能得到的字典序最小字符串，即 $f[i+1]$。
- $f[i]$ 不包含 $s[i]$，且子串 $s[i]$ 到 $s[j]$ 可以完全消除，那么 $f[i]$ 等于后缀 $s[j+1:]$ 中能得到的字典序最小字符串，即 $f[j+1]$。

这些字符串取最小值，即为 $f[i]$。

初始值：$f[n]$ 等于空串。

答案：$f[0]$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
def is_consecutive(x: str, y: str) -> bool:
    d = abs(ord(x) - ord(y))
    return d == 1 or d == 25

class Solution:
    def lexicographicallySmallestString(self, s: str) -> str:
        n = len(s)
        can_be_empty = [[False] * n for _ in range(n)]
        for i in range(n - 2, -1, -1):
            can_be_empty[i + 1][i] = True  # 空串
            for j in range(i + 1, n):
                # 性质 2
                if is_consecutive(s[i], s[j]) and can_be_empty[i + 1][j - 1]:
                    can_be_empty[i][j] = True
                    continue
                # 性质 3
                for k in range(i + 1, j - 1):
                    if can_be_empty[i][k] and can_be_empty[k + 1][j]:
                        can_be_empty[i][j] = True
                        break

        f = [''] * (n + 1)
        for i in range(n - 1, -1, -1):
            res = s[i] + f[i + 1]
            for j in range(i + 1, n):
                if can_be_empty[i][j]:
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
            for (int j = i + 1; j < n; j++) {
                // 性质 2
                if (isConsecutive(s[i], s[j]) && canBeEmpty[i + 1][j - 1]) {
                    canBeEmpty[i][j] = true;
                    continue;
                }
                // 性质 3
                for (int k = i + 1; k < j - 1; k++) {
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
            String res = s[i] + f[i + 1];
            for (int j = i + 1; j < n; j++) {
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
            for (int j = i + 1; j < n; j++) {
                // 性质 2
                if (is_consecutive(s[i], s[j]) && can_be_empty[i + 1][j - 1]) {
                    can_be_empty[i][j] = true;
                    continue;
                }
                // 性质 3
                for (int k = i + 1; k < j - 1; k++) {
                    if (can_be_empty[i][k] && can_be_empty[k + 1][j]) {
                        can_be_empty[i][j] = true;
                        break;
                    }
                }
            }
        }

        vector<string> f(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            f[i] = s[i] + f[i + 1];
            for (int j = i + 1; j < n; j++) {
                if (can_be_empty[i][j]) {
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
		for j := i + 1; j < n; j++ {
			// 性质 2
			if isConsecutive(s[i], s[j]) && canBeEmpty[i+1][j-1] {
				canBeEmpty[i][j] = true
				continue
			}
			// 性质 3
			for k := i + 1; k < j-1; k++ {
				if canBeEmpty[i][k] && canBeEmpty[k+1][j] {
					canBeEmpty[i][j] = true
					break
				}
			}
		}
	}

	f := make([]string, n+1)
	for i := n - 1; i >= 0; i-- {
		res := string(s[i]) + f[i+1]
		for j := i + 1; j < n; j++ {
			if canBeEmpty[i][j] {
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
