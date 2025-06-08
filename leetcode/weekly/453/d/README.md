## 划分型 DP

本题是标准的划分型 DP，见 [DP 题单](https://leetcode.cn/circle/discuss/tXLS3i/) 的「§5.2 最优划分」。

定义 $f[i+1]$ 表示前缀 $[0,i]$ 为了达成题目要求，需要的最少操作次数。这里 $i+1$ 是为了把 $f[0]$ 当作初始值。

枚举最后一个子串的左端点 $j$，那么问题变成前缀 $[0,j-1]$ 为了达成题目要求，需要的最少操作次数，即 $f[j]$。取最小值，得

$$
f[i+1] = \min_{j=0}^{i} f[j] + \text{op}(j,i)
$$

初始值 $f[0] = 0$，空串无需操作。

答案为 $f[n]$。

## 贪心

如何计算子串 $[j,i]$ 的操作次数 $\text{op}(j,i)$？

先只考虑操作 1 和操作 2。

注意到，操作 2 等价两次操作 1，所以应该最大化操作 2 的次数。

为方便描述，下文把 $\textit{word}_1$ 和 $\textit{word}_2$ 简记为 $s$ 和 $t$。

对于区间 $[j,i]$ 中的一对下标 $(p,q)$，如果 $s_p=t_q$ 且 $s_q=t_p$，那么就可以执行一次操作 2。这个条件可以重新表述为：

- 字母对 $(s_p,t_p)$ 等于字母对 $(t_q,s_q)$。

用一个 $26\times 26$ 的数组 $\textit{cnt}$，在遍历这段子串的同时，统计字母对 $(s_p,t_p)$ 的个数：

- 如果 $s_p=t_p$，无需操作。
- 如果 $\textit{cnt}[t_p][s_p] = 0$，那么把 $\textit{cnt}[s_p][t_p]$ 加一，操作次数加一（暂时使用操作 1）
- 如果 $\textit{cnt}[t_p][s_p] > 0$，那么找到了一对可以交换的下标，把 $\textit{cnt}[t_p][s_p]$ 减一（把之前执行的操作 1 改成操作 2）。

然后考虑操作 3，这个其实简单，只需把上面的条件改成：

- 字母对 $(s_p,t_{i+j-p})$ 等于字母对 $(t_{i+j-q},s_q)$。

重新统计一遍即可。注意操作次数要额外加一（因为执行了一次操作 3）。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str, t: str) -> int:
        n = len(s)
        f = [0] * (n + 1)
        for i in range(n):
            res = sys.maxsize
            cnt = defaultdict(int)
            op = 0
            for j in range(i, -1, -1):
                # 不反转
                x, y = s[j], t[j]
                if x != y:
                    if cnt[(y, x)] > 0:
                        cnt[(y, x)] -= 1
                    else:
                        cnt[(x, y)] += 1
                        op += 1

                # 反转
                rev_cnt = defaultdict(int)
                rev_op = 1
                for p in range(j, i + 1):
                    x, y = s[p], t[i + j - p]
                    if x == y:
                        continue
                    if rev_cnt[(y, x)] > 0:
                        rev_cnt[(y, x)] -= 1
                    else:
                        rev_cnt[(x, y)] += 1
                        rev_op += 1

                res = min(res, f[j] + min(op, rev_op))
            f[i + 1] = res
        return f[n]
```

```java [sol-Java]
class Solution {
    public int minOperations(String S, String T) {
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;
        int[] f = new int[n + 1];

        for (int i = 0; i < n; i++) {
            int res = Integer.MAX_VALUE;
            int[][] cnt = new int[26][26];
            int op = 0;

            for (int j = i; j >= 0; j--) {
                // 不反转
                int x = s[j] - 'a', y = t[j] - 'a';
                if (x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }

                // 反转
                int[][] revCnt = new int[26][26];
                int revOp = 1;
                for (int p = j; p <= i; p++) {
                    x = s[p] - 'a';
                    y = t[i + j - p] - 'a';
                    if (x == y) {
                        continue;
                    }
                    if (revCnt[y][x] > 0) {
                        revCnt[y][x]--;
                    } else {
                        revCnt[x][y]++;
                        revOp++;
                    }
                }

                res = Math.min(res, f[j] + Math.min(op, revOp));
            }

            f[i + 1] = res;
        }

        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s, string t) {
        int n = s.size();
        vector<int> f(n + 1);
        for (int i = 0; i < n; i++) {
            int res = INT_MAX;
            int cnt[26][26]{};
            int op = 0;
            for (int j = i; j >= 0; j--) {
                // 不反转
                int x = s[j] - 'a', y = t[j] - 'a';
                if (x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }

                // 反转
                int rev_cnt[26][26]{};
                int rev_op = 1;
                for (int p = j; p <= i; p++) {
                    int x = s[p] - 'a', y = t[i + j - p] - 'a';
                    if (x == y) {
                        continue;
                    }
                    if (rev_cnt[y][x] > 0) {
                        rev_cnt[y][x]--;
                    } else {
                        rev_cnt[x][y]++;
                        rev_op++;
                    }
                }

                res = min(res, f[j] + min(op, rev_op));
            }
            f[i + 1] = res;
        }
        return f[n];
    }
};
```

```go [sol-Go]
func minOperations(s, t string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt
		cnt := [26][26]int{}
		op := 0
		for j := i; j >= 0; j-- {
			// 不反转
			x, y := s[j]-'a', t[j]-'a'
			if x != y {
				if cnt[y][x] > 0 {
					cnt[y][x]--
				} else {
					cnt[x][y]++
					op++
				}
			}

			// 反转
			revCnt := [26][26]int{}
			revOp := 1
			for p := j; p <= i; p++ {
				x, y := s[p]-'a', t[i+j-p]-'a'
				if x == y {
					continue
				}
				if revCnt[y][x] > 0 {
					revCnt[y][x]--
				} else {
					revCnt[x][y]++
					revOp++
				}
			}

			res = min(res, f[j]+min(op, revOp))
		}
		f[i+1] = res
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单的「**§5.2 最优划分**」。

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
