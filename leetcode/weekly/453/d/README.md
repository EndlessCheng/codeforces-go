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

如何计算子串 $[j,i]$ 的最小操作次数 $\text{op}(j,i)$？

先只考虑操作 1 和操作 2。

注意到，操作 2 等价两次操作 1，所以应该最大化操作 2 的次数。

为方便描述，下文把 $\textit{word}_1$ 和 $\textit{word}_2$ 简记为 $s$ 和 $t$。

对于区间 $[j,i]$ 中的一对下标 $(p,q)$，如果 $s_p=t_q$ 且 $s_q=t_p$，那么就可以执行一次操作 2。这个条件可以重新表述为：

- 字母对 $(s_p,t_p)$ 等于字母对 $(t_q,s_q)$。

用一个 $26\times 26$ 的数组 $\textit{cnt}$，在遍历这段子串的同时，统计字母对 $(s_p,t_p)$ 的个数：

- 如果 $s_p=t_p$，无需操作。
- 如果 $\textit{cnt}[t_p][s_p] = 0$，那么把 $\textit{cnt}[s_p][t_p]$ 加一，操作次数加一（暂时使用操作 1）。
- 如果 $\textit{cnt}[t_p][s_p] > 0$，那么找到了一对可以交换的下标，把 $\textit{cnt}[t_p][s_p]$ 减一（把之前执行的操作 1 改成操作 2）。

然后考虑操作 3，这个其实简单，只需把上面的条件改成：

- 字母对 $(s_p,t_{i+j-p})$ 等于字母对 $(t_{i+j-q},s_q)$。

重新统计一遍即可。注意操作次数要额外加一（因为执行了一次操作 3）。

具体请看 [视频讲解](https://www.bilibili.com/video/BV113T9zFEjQ/?t=28m17s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str, t: str) -> int:
        n = len(s)
        f = [0] * (n + 1)
        for i in range(n):
            res = inf
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
                int x = s[j] - 'a';
                int y = t[j] - 'a';
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
                    char x = s[p] - 'a', y = t[i + j - p] - 'a';
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

## 优化：预处理 revOp

预处理区间 $[j,i]$ 对应的 $\textit{revOp}$。

上面计算 $\textit{op}$ 是增量计算的，那么 $\textit{revOp}$ 能否也地增量计算呢？

关键在于，要**保证子串的中心是不变的**，这可以用**中心扩展法**（见 [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)）实现，从区间 $[j+1,i-1]$ 增量地计算区间 $[j,i]$，这样计算所有 $\textit{revOp}$ 一共只需要 $\mathcal{O}(n^2)$ 的时间。

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str, t: str) -> int:
        n = len(s)

        # 预处理 rev_op
        rev_op = [[0] * n for _ in range(n)]
        # 中心扩展法
        # i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
        for i in range(2 * n - 1):
            cnt = defaultdict(int)
            op = 1
            # 从闭区间 [l, r] 开始向左右扩展
            l, r = i // 2, (i + 1) // 2
            while l >= 0 and r < n:
                x, y = s[l], t[r]
                if x != y:
                    if cnt[(y, x)] > 0:
                        cnt[(y, x)] -= 1
                    else:
                        cnt[(x, y)] += 1
                        op += 1

                x, y = s[r], t[l]
                if l != r and x != y:
                    if cnt[(y, x)] > 0:
                        cnt[(y, x)] -= 1
                    else:
                        cnt[(x, y)] += 1
                        op += 1

                rev_op[l][r] = op
                l -= 1
                r += 1

        f = [0] * (n + 1)
        for i in range(n):
            res = inf
            cnt = defaultdict(int)
            op = 0  # 不反转时的最小操作次数
            for j in range(i, -1, -1):
                x, y = s[j], t[j]
                if x != y:
                    if cnt[(y, x)] > 0:
                        cnt[(y, x)] -= 1
                    else:
                        cnt[(x, y)] += 1
                        op += 1
                res = min(res, f[j] + min(op, rev_op[j][i]))
            f[i + 1] = res
        return f[n]
```

```java [sol-Java]
class Solution {
    public int minOperations(String S, String T) {
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;

        // 预处理 revOp
        int[][] revOp = new int[n][n];
        // 中心扩展法
        // i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
        for (int i = 0; i < 2 * n - 1; i++) {
            int[][] cnt = new int[26][26];
            int op = 1;
            // 从闭区间 [l, r] 开始向左右扩展
            int l = i / 2;
            int r = (i + 1) / 2;
            while (l >= 0 && r < n) {
                int x = s[l] - 'a';
                int y = t[r] - 'a';
                if (x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }

                x = s[r] - 'a';
                y = t[l] - 'a';
                if (l != r && x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }

                revOp[l][r] = op;
                l--;
                r++;
            }
        }

        int[] f = new int[n + 1];
        for (int i = 0; i < n; i++) {
            int res = Integer.MAX_VALUE;
            int[][] cnt = new int[26][26];
            int op = 0; // 不反转时的最小操作次数
            for (int j = i; j >= 0; j--) {
                int x = s[j] - 'a';
                int y = t[j] - 'a';
                if (x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }
                res = Math.min(res, f[j] + Math.min(op, revOp[j][i]));
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
        // 预处理 rev_op
        vector rev_op(n, vector<int>(n));
        // 中心扩展法
        // i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
        for (int i = 0; i < 2 * n - 1; i++) {
            int cnt[26][26]{};
            int op = 1;
            // 从闭区间 [l, r] 开始向左右扩展
            int l = i / 2, r = (i + 1) / 2;
            while (l >= 0 && r < n) {
                char x = s[l] - 'a', y = t[r] - 'a';
                if (x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }

                x = s[r] - 'a', y = t[l] - 'a';
                if (l != r && x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }

                rev_op[l][r] = op;
                l--;
                r++;
            }
        }

        vector<int> f(n + 1);
        for (int i = 0; i < n; i++) {
            int res = INT_MAX;
            int cnt[26][26]{};
            int op = 0; // 不反转时的最小操作次数
            for (int j = i; j >= 0; j--) {
                char x = s[j] - 'a', y = t[j] - 'a';
                if (x != y) {
                    if (cnt[y][x] > 0) {
                        cnt[y][x]--;
                    } else {
                        cnt[x][y]++;
                        op++;
                    }
                }
                res = min(res, f[j] + min(op, rev_op[j][i]));
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
	// 预处理 revOp
	revOp := make([][]int, n)
	for i := range revOp {
		revOp[i] = make([]int, n)
	}
	// 中心扩展法
	// i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
	for i := range 2*n - 1 {
		cnt := [26][26]int{}
		op := 1
		// 从闭区间 [l,r] 开始向左右扩展
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n {
			x, y := s[l]-'a', t[r]-'a'
			if x != y {
				if cnt[y][x] > 0 {
					cnt[y][x]--
				} else {
					cnt[x][y]++
					op++
				}
			}

			x, y = s[r]-'a', t[l]-'a'
			if l != r && x != y {
				if cnt[y][x] > 0 {
					cnt[y][x]--
				} else {
					cnt[x][y]++
					op++
				}
			}

			revOp[l][r] = op
			l--
			r++
		}
	}

	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt
		cnt := [26][26]int{}
		op := 0 // 不反转时的最小操作次数
		for j := i; j >= 0; j-- {
			x, y := s[j]-'a', t[j]-'a'
			if x != y {
				if cnt[y][x] > 0 {
					cnt[y][x]--
				} else {
					cnt[x][y]++
					op++
				}
			}
			res = min(res, f[j]+min(op, revOp[j][i]))
		}
		f[i+1] = res
	}
	return f[n]
}
```

有很多重复代码，重构一下：

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str, t: str) -> int:
        def update(x: str, y: str) -> None:
            if x == y:
                return
            if cnt[(y, x)] > 0:
                cnt[(y, x)] -= 1
            else:
                cnt[(x, y)] += 1
                nonlocal op
                op += 1

        n = len(s)
        # 预处理 rev_op
        rev_op = [[0] * n for _ in range(n)]
        # 中心扩展法
        # i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
        for i in range(2 * n - 1):
            cnt = defaultdict(int)
            op = 1
            # 从闭区间 [l, r] 开始向左右扩展
            l, r = i // 2, (i + 1) // 2
            while l >= 0 and r < n:
                update(s[l], t[r])
                if l != r:
                    update(s[r], t[l])
                rev_op[l][r] = op
                l -= 1
                r += 1

        f = [0] * (n + 1)
        for i in range(n):
            res = inf
            cnt = defaultdict(int)
            op = 0  # 不反转时的最小操作次数
            for j in range(i, -1, -1):
                update(s[j], t[j])
                res = min(res, f[j] + min(op, rev_op[j][i]))
            f[i + 1] = res
        return f[n]
```

```java [sol-Java]
class Solution {
    public int minOperations(String S, String T) {
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;

        // 预处理 revOp
        int[][] revOp = new int[n][n];
        // 中心扩展法
        // i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
        for (int i = 0; i < 2 * n - 1; i++) {
            cnt = new int[26][26];
            op = 1;
            // 从闭区间 [l, r] 开始向左右扩展
            int l = i / 2;
            int r = (i + 1) / 2;
            while (l >= 0 && r < n) {
                update(s[l], t[r]);
                if (l != r) {
                    update(s[r], t[l]);
                }
                revOp[l][r] = op;
                l--;
                r++;
            }
        }

        int[] f = new int[n + 1];
        for (int i = 0; i < n; i++) {
            int res = Integer.MAX_VALUE;
            cnt = new int[26][26];
            op = 0; // 不反转时的最小操作次数
            for (int j = i; j >= 0; j--) {
                update(s[j], t[j]);
                res = Math.min(res, f[j] + Math.min(op, revOp[j][i]));
            }
            f[i + 1] = res;
        }
        return f[n];
    }

    private int op;
    private int[][] cnt;

    private void update(char x, char y) {
        if (x == y) {
            return;
        }
        x -= 'a';
        y -= 'a';
        if (cnt[y][x] > 0) {
            cnt[y][x]--;
        } else {
            cnt[x][y]++;
            op++;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    void update(array<array<int, 26>, 26>& cnt, int& op, char x, char y) {
        if (x == y) {
            return;
        }
        x -= 'a';
        y -= 'a';
        if (cnt[y][x] > 0) {
            cnt[y][x]--;
        } else {
            cnt[x][y]++;
            op++;
        }
    }

public:
    int minOperations(string s, string t) {
        int n = s.size();
        // 预处理 rev_op
        vector rev_op(n, vector<int>(n));
        // 中心扩展法
        // i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
        for (int i = 0; i < 2 * n - 1; i++) {
            array<array<int, 26>, 26> cnt{};
            int op = 1;
            // 从闭区间 [l, r] 开始向左右扩展
            int l = i / 2, r = (i + 1) / 2;
            while (l >= 0 && r < n) {
                update(cnt, op, s[l], t[r]);
                if (l != r) {
                    update(cnt, op, s[r], t[l]);
                }
                rev_op[l][r] = op;
                l--;
                r++;
            }
        }

        vector<int> f(n + 1);
        for (int i = 0; i < n; i++) {
            int res = INT_MAX;
            array<array<int, 26>, 26> cnt{};
            int op = 0; // 不反转时的最小操作次数
            for (int j = i; j >= 0; j--) {
                update(cnt, op, s[j], t[j]);
                res = min(res, f[j] + min(op, rev_op[j][i]));
            }
            f[i + 1] = res;
        }
        return f[n];
    }
};
```

```go [sol-Go]
func minOperations(s, t string) int {
	var cnt [26][26]int
	var op int
	update := func(x, y byte) {
		if x == y {
			return
		}
		x -= 'a'
		y -= 'a'
		if cnt[y][x] > 0 {
			cnt[y][x]--
		} else {
			cnt[x][y]++
			op++
		}
	}

	n := len(s)
	// 预处理 revOp
	revOp := make([][]int, n)
	for i := range revOp {
		revOp[i] = make([]int, n)
	}
	// 中心扩展法
	// i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
	for i := range 2*n - 1 {
		cnt = [26][26]int{}
		op = 1
		// 从闭区间 [l,r] 开始向左右扩展
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n {
			update(s[l], t[r])
			if l != r {
				update(s[r], t[l])
			}
			revOp[l][r] = op
			l--
			r++
		}
	}

	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt
		cnt = [26][26]int{}
		op = 0 // 不反转时的最小操作次数
		for j := i; j >= 0; j-- {
			update(s[j], t[j])
			res = min(res, f[j]+min(op, revOp[j][i]))
		}
		f[i+1] = res
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

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
