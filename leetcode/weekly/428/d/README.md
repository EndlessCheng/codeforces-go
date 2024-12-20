## 核心思路

1. 统计 $s$ 每种字母的个数，记到一个长为 $26$ 的 $\textit{cnt}$ 数组中。
2. 设 $m = \max(\textit{cnt})$。枚举每种字母都变成 $\textit{target}$，其中 $\textit{target}=0,1,2,\ldots,m$。
3. 每种字母的个数要么变成 $0$，要么变成 $\textit{target}$。
4. 分析性质（下文细说），发现需要写一个线性 DP，决策：
    1. 当前字母的个数单独变成 $0$ 或者 $\textit{target}$；
    2. 当前字母和下一个字母的个数一起考虑，变成 $0$ 或者 $\textit{target}$。

## 分析性质

比如 a 出现了 $6$ 次，b 出现了 $2$ 次，$\textit{target} = 3$。

- 如果 $6$ 和 $2$ 都单独变成 $3$（使用操作一和操作二），那么需要操作 $(6-3)+(3-2)=4$ 次。
- 使用 $1$ 次操作三，先把 $6$ 变成 $5$，把 $2$ 变成 $3$，然后再使用 $5-3=2$ 次操作一，一共使用 $1+2=3$ 次操作，就可以把 $6$ 和 $2$ 都变成 $3$。

所以关键在于，通过恰当地使用操作三，节省操作次数。

> 注意操作三只能用于相邻的两种字母，不能用于连续三种或更多种字母。比如把一个 a 变成 b，然后再把这个 b 变成 c，共执行 $2$ 次操作三，这等价于 $1$ 次操作二和 $1$ 次操作一，我们不会得到更优的结果。

设当前字母出现了 $x$ 次，下一个字母出现了 $y$ 次。

分类讨论：

- 单独操作 $x$：讨论变成 $0$ 还是 $\textit{target}$，需要操作 $\min(x,|x-\textit{target}|)$ 次。
- 如果 $y\ge \textit{target}$，由于操作三会把 $y$ 变大，我们后面还要再把 $y$ 变小，这不如直接单独操作 $x$ 和 $y$。所以这种情况不考虑操作三。
- 如果 $y< \textit{target}$：
  - 如果 $x> \textit{target}$，那么可以把 $x$ 和 $y$ 都变成 $\textit{target}$，操作 $\max(x-\textit{target}, \textit{target}-y)$ 次。
  - 如果 $x\le \textit{target}$，那么可以把 $x$ 变成 $0$，$y$ 变成 $\textit{target}$，操作 $\max(x, \textit{target}-y)$ 次。

## 动态规划

我们不知道哪些字母之间要使用操作三，所以**把所有情况全部枚举一遍**。由于有大量重复子问题，用动态规划减少计算量。

比如单独操作字母 a，那么子问题为操作字母 b 到 z 的最小操作次数。

比如同时操作字母 a 和 b，那么子问题为操作字母 c 到 z 的最小操作次数。

故定义 $f[i]$ 表示操作第 $i$ 到 $25$ 种字母的最小操作次数。这里 $i$ 从 $0$ 开始。

设 $x=\textit{cnt}[i]$。单独操作 $x$，有 $f[i] = f[i+1] + \min(x,|x-\textit{target}|)$。

设 $y=\textit{cnt}[i+1]$。如果 $y< \textit{target}$，分类讨论：

- 如果 $x> \textit{target}$，那么可以把 $x$ 和 $y$ 都变成 $\textit{target}$，有 $f[i] = f[i+2] + \max(x-\textit{target}, \textit{target}-y)$。
- 如果 $x\le \textit{target}$，那么可以把 $x$ 变成 $0$，$y$ 变成 $\textit{target}$，有 $f[i] = f[i+2] + \max(x, \textit{target}-y)$。

初始值：$f[26]=0,\ f[25]=\min(\textit{cnt}[25],|\textit{cnt}[25]-\textit{target}|)$。

答案：$f[0]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1pnqZYKEqr/?t=40m52s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def makeStringGood(self, s: str) -> int:
        cnt = Counter(s)
        cnt = [cnt[c] for c in ascii_lowercase]

        ans = len(s)  # target = 0 时的答案
        f = [0] * 27
        for target in range(1, max(cnt) + 1):
            f[25] = min(cnt[25], abs(cnt[25] - target))
            for i in range(24, -1, -1):
                x, y = cnt[i], cnt[i + 1]
                # 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + min(x, abs(x - target))
                # x 变成 target 或 0，y 变成 target
                if y < target:  # 只有当 y 需要变大时，才去执行第三种操作
                    t = target if x > target else 0
                    f[i] = min(f[i], f[i + 2] + max(x - t, target - y))
            ans = min(ans, f[0])
        return ans
```

```java [sol-Java]
class Solution {
    public int makeStringGood(String s) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }
        int m = Arrays.stream(cnt).max().getAsInt();

        int ans = s.length(); // target = 0 时的答案
        int[] f = new int[27];
        for (int target = 1; target <= m; target++) {
            f[25] = Math.min(cnt[25], Math.abs(cnt[25] - target));
            for (int i = 24; i >= 0; i--) {
                int x = cnt[i];
                int y = cnt[i + 1];
                // 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + Math.min(x, Math.abs(x - target));
                // x 变成 target 或 0，y 变成 target
                if (y < target) { // 只有当 y 需要变大时，才去执行第三种操作
                    int t = x > target ? target : 0;
                    f[i] = Math.min(f[i], f[i + 2] + Math.max(x - t, target - y));
                }
            }
            ans = Math.min(ans, f[0]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int makeStringGood(string s) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }
        int m = ranges::max(cnt);

        int ans = s.length(); // target = 0 时的答案
        int f[27]{};
        for (int target = 1; target <= m; target++) {
            f[25] = min(cnt[25], abs(cnt[25] - target));
            for (int i = 24; i >= 0; i--) {
                int x = cnt[i], y = cnt[i + 1];
                // 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + min(x, abs(x - target));
                // x 变成 target 或 0，y 变成 target
                if (y < target) { // 只有当 y 需要变大时，才去执行第三种操作
                    int t = x > target ? target : 0;
                    f[i] = min(f[i], f[i + 2] + max(x - t, target - y));
                }
            }
            ans = min(ans, f[0]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func makeStringGood(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	m := slices.Max(cnt[:])

	ans := len(s) // target = 0 时的答案
	f := [27]int{}
	for target := 1; target <= m; target++ {
		f[25] = min(cnt[25], abs(cnt[25]-target))
		for i := 24; i >= 0; i-- {
			x, y := cnt[i], cnt[i+1]
			// 单独操作 x（变成 target 或 0）
			f[i] = f[i+1] + min(x, abs(x-target))
			// x 变成 target 或 0，y 变成 target
			if y < target { // 只有当 y 需要变大时，才去执行第三种操作
				if x > target { // x 变成 target
					f[i] = min(f[i], f[i+2]+max(x-target, target-y))
				} else { // x 变成 0
					f[i] = min(f[i], f[i+2]+max(x, target-y))
				}
			}
		}
		ans = min(ans, f[0])
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

## 优化一

如果 $x = \textit{cnt}[i]=0$，无需操作，$f[i]=f[i+1]$。

如果 $y = \textit{cnt}[i+1]=0$，只需单独操作 $x$，$f[i] = f[i+1] + \min(x,|x-\textit{target}|)$。

这个优化对 Python 十分显著，其他语言不明显。

```py [sol-Python3]
class Solution:
    def makeStringGood(self, s: str) -> int:
        cnt = Counter(s)
        cnt = [cnt[c] for c in ascii_lowercase]

        ans = len(s)  # target = 0 时的答案
        f = [0] * 27
        for target in range(1, max(cnt) + 1):
            f[25] = min(cnt[25], abs(cnt[25] - target))
            for i in range(24, -1, -1):
                x = cnt[i]
                if x == 0:
                    f[i] = f[i + 1]
                    continue
                # 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + min(x, abs(x - target))
                # x 变成 target 或 0，y 变成 target
                y = cnt[i + 1]
                if 0 < y < target:  # 只有当 y 需要变大时，才去执行第三种操作
                    t = target if x > target else 0
                    f[i] = min(f[i], f[i + 2] + max(x - t, target - y))
            ans = min(ans, f[0])
        return ans
```

```java [sol-Java]
class Solution {
    public int makeStringGood(String s) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }
        int m = Arrays.stream(cnt).max().getAsInt();

        int ans = s.length(); // target = 0 时的答案
        int[] f = new int[27];
        for (int target = 1; target <= m; target++) {
            f[25] = Math.min(cnt[25], Math.abs(cnt[25] - target));
            for (int i = 24; i >= 0; i--) {
                int x = cnt[i];
                if (x == 0) {
                    f[i] = f[i + 1];
                    continue;
                }
                // 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + Math.min(x, Math.abs(x - target));
                // x 变成 target 或 0，y 变成 target
                int y = cnt[i + 1];
                if (0 < y && y < target) { // 只有当 y 需要变大时，才去执行第三种操作
                    int t = x > target ? target : 0;
                    f[i] = Math.min(f[i], f[i + 2] + Math.max(x - t, target - y));
                }
            }
            ans = Math.min(ans, f[0]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int makeStringGood(string s) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }
        int m = ranges::max(cnt);

        int ans = s.length(); // target = 0 时的答案
        int f[27]{};
        for (int target = 1; target <= m; target++) {
            f[25] = min(cnt[25], abs(cnt[25] - target));
            for (int i = 24; i >= 0; i--) {
                int x = cnt[i];
                if (x == 0) {
                    f[i] = f[i + 1];
                    continue;
                }
                // 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + min(x, abs(x - target));
                // x 变成 target 或 0，y 变成 target
                int y = cnt[i + 1];
                if (0 < y && y < target) { // 只有当 y 需要变大时，才去执行第三种操作
                    int t = x > target ? target : 0;
                    f[i] = min(f[i], f[i + 2] + max(x - t, target - y));
                }
            }
            ans = min(ans, f[0]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func makeStringGood(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	m := slices.Max(cnt[:])

	ans := len(s) // target = 0 时的答案
	f := [27]int{}
	for target := 1; target <= m; target++ {
		f[25] = min(cnt[25], abs(cnt[25]-target))
		for i := 24; i >= 0; i-- {
			x := cnt[i]
			if x == 0 {
				f[i] = f[i+1]
				continue
			}
			// 单独操作 x（变成 target 或 0）
			f[i] = f[i+1] + min(x, abs(x-target))
			// x 变成 target 或 0，y 变成 target
			y := cnt[i+1]
			if 0 < y && y < target { // 只有当 y 需要变大时，才去执行第三种操作
				if x > target { // x 变成 target
					f[i] = min(f[i], f[i+2]+max(x-target, target-y))
				} else { // x 变成 0
					f[i] = min(f[i], f[i+2]+max(x, target-y))
				}
			}
		}
		ans = min(ans, f[0])
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 优化二

不需要枚举那么多种 $\textit{target}$，只需考虑如下数字：

1. $\textit{cnt}[i]$。变成 $\textit{cnt}$ 中的某个数。
2. $\textit{cnt}[i] + \textit{cnt}[i+1]$。操作三把 $x$ 变成 $0$，把 $y$ 变成 $x+y$。
3. $\left\lfloor\dfrac{\textit{cnt}[i]+\textit{cnt}[i+1]}{2}\right\rfloor$。操作三把 $x$ 和 $y$ 都变成 $\textit{target}$，取平均值最优。
4. $\left\lceil\dfrac{\textit{cnt}[i]+\textit{cnt}[i+1]}{2}\right\rceil$。同上。

如此修改后，即使题目给你一堆询问，每个询问要你计算 $s$ 的一个子串的答案，也能轻松应对。（提示：前缀和）

```py [sol-Python3]
class Solution:
    def makeStringGood(self, s: str) -> int:
        cnt = Counter(s)
        cnt = [cnt[c] for c in ascii_lowercase]

        targets = set(cnt)
        for x, y in pairwise(cnt):
            targets.add(x + y)
            targets.add((x + y) // 2)
            targets.add((x + y + 1) // 2)

        ans = len(s)  # target = 0 时的答案
        f = [0] * 27
        for target in targets:
            f[25] = min(cnt[25], abs(cnt[25] - target))
            for i in range(24, -1, -1):
                x = cnt[i]
                if x == 0:
                    f[i] = f[i + 1]
                    continue
                # 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + min(x, abs(x - target))
                # x 变成 target 或 0，y 变成 target
                y = cnt[i + 1]
                if 0 < y < target:  # 只有当 y 需要变大时，才去执行第三种操作
                    t = target if x > target else 0
                    f[i] = min(f[i], f[i + 2] + max(x - t, target - y))
            ans = min(ans, f[0])
        return ans
```

```java [sol-Java]
class Solution {
    public int makeStringGood(String s) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }

        Set<Integer> targets = new HashSet<>();
        targets.add(cnt[0]);
        for (int i = 1; i < 26; i++) {
            int x = cnt[i - 1];
            int y = cnt[i];
            targets.add(x);
            targets.add(x + y);
            targets.add((x + y) / 2);
            targets.add((x + y + 1) / 2);
        }

        int ans = s.length(); // target = 0 时的答案
        int[] f = new int[27];
        for (int target : targets) {
            f[25] = Math.min(cnt[25], Math.abs(cnt[25] - target));
            for (int i = 24; i >= 0; i--) {
                int x = cnt[i];
                if (x == 0) {
                    f[i] = f[i + 1];
                    continue;
                }
                // 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + Math.min(x, Math.abs(x - target));
                // x 变成 target 或 0，y 变成 target
                int y = cnt[i + 1];
                if (0 < y && y < target) { // 只有当 y 需要变大时，才去执行第三种操作
                    int t = x > target ? target : 0;
                    f[i] = Math.min(f[i], f[i + 2] + Math.max(x - t, target - y));
                }
            }
            ans = Math.min(ans, f[0]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int makeStringGood(string s) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }

        unordered_set<int> targets(cnt, cnt + 26);
        for (int i = 1; i < 26; i++) {
            int x = cnt[i - 1], y = cnt[i];
            targets.insert(x + y);
            targets.insert((x + y) / 2);
            targets.insert((x + y + 1) / 2);
        }

        int ans = s.length(); // target = 0 时的答案
        int f[27]{};
        for (int target : targets) {
            f[25] = min(cnt[25], abs(cnt[25] - target));
            for (int i = 24; i >= 0; i--) {
                int x = cnt[i];
                if (x == 0) {
                    f[i] = f[i + 1];
                    continue;
                }
                // 单独操作 x（变成 target 或 0）
                f[i] = f[i + 1] + min(x, abs(x - target));
                // x 变成 target 或 0，y 变成 target
                int y = cnt[i + 1];
                if (0 < y && y < target) { // 只有当 y 需要变大时，才去执行第三种操作
                    int t = x > target ? target : 0;
                    f[i] = min(f[i], f[i + 2] + max(x - t, target - y));
                }
            }
            ans = min(ans, f[0]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func makeStringGood(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	targets := map[int]struct{}{}
	targets[cnt[0]] = struct{}{}
	for i := 1; i < 26; i++ {
		x, y := cnt[i-1], cnt[i]
		targets[y] = struct{}{}
		targets[x+y] = struct{}{}
		targets[(x+y)/2] = struct{}{}
		targets[(x+y+1)/2] = struct{}{}
	}

	ans := len(s) // target = 0 时的答案
	f := [27]int{}
	for target := range targets {
		f[25] = min(cnt[25], abs(cnt[25]-target))
		for i := 24; i >= 0; i-- {
			x := cnt[i]
			if x == 0 {
				f[i] = f[i+1]
				continue
			}
			// 单独操作 x（变成 target 或 0）
			f[i] = f[i+1] + min(x, abs(x-target))
			// x 变成 target 或 0，y 变成 target
			y := cnt[i+1]
			if 0 < y && y < target { // 只有当 y 需要变大时，才去执行第三种操作
				if x > target { // x 变成 target
					f[i] = min(f[i], f[i+2]+max(x-target, target-y))
				} else { // x 变成 0
					f[i] = min(f[i], f[i+2]+max(x, target-y))
				}
			}
		}
		ans = min(ans, f[0])
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^2)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 变形题

1. 写代码，实现优化二中提到的变形。
2. 如果可以把 z 改成 a，要怎么做？提示：[213. 打家劫舍 II](https://leetcode.cn/problems/house-robber-ii/)。
3. 如果每种操作都有相应的次数限制呢？即额外输入 $\textit{limit}_1,\textit{limit}_2,\textit{limit}_3$。
4. 如果每种操作都有相应的代价呢？即额外输入 $\textit{cost}_1,\textit{cost}_2,\textit{cost}_3$。

更多相似题目，见下面动态规划题单中的「**§7.1 一维 DP**」。

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
