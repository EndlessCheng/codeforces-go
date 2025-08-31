## 理解题意

注意字符串长度只有 $2$。

题目要求所选字符串必须包含 $\texttt{x}$，所以我们能选的字符串有两组：

- 第一组：字符串的第一个字母是 $\texttt{x}$，即 $\texttt{xa},\texttt{xb},\texttt{xc},\ldots$ 
- 第二组：字符串的第二个字母是 $\texttt{x}$，即 $\texttt{ax},\texttt{bx},\texttt{cx},\ldots$

比如从第一组中选一个 $\texttt{xa}$ 和 $\texttt{xb}$，这两个字符串恰好有一个下标上的字母不同，根据题意，它们是兼容的，可以消除，得分加一。

比如从第一组中选一个 $\texttt{xa}$，从第二组中选一个 $\texttt{bx}$，这两个字符串有两个下标上的字母不同，不符合要求，所以不能这么选。

看上去，第一组和第二组互相独立？然而字符串 $\texttt{xx}$ 在两组中都有。

## 没有 xx 的情况

从特殊到一般。首先思考，在没有 $\texttt{xx}$ 字符串的情况下，要怎么做。

在这种情况下，第一组和第二组是互相独立的，可以分别计算。

以第一组为例说明。

统计第二个字母 $s[1]$ 的出现次数，记在一个 $\textit{cnt}$ 数组中。

问题相当于：

- 给定数组 $\textit{cnt}$，每次操作，选两个下标不同的正整数，各减少一。目标：最大化操作次数。

由于每次操作的都是两个下标不同的数，把这些下标按顺序拼接，可以构造出一个**相邻元素不同**的序列。例如 $(1,2),(2,3),(3,4)$ 这三个操作，可以拼接成 $[1,2,3,2,3,4]$。注意为了保证相邻不同，把 $(2,3)$ 调整为等价的 $(3,2)$，都表示选一个下标 $2$ 和一个下标 $3$。

设 $\textit{sum} = \sum\limits_{i}\textit{cnt}[i]$，$\textit{mx} = \max(\textit{cnt})$。

**定理**：如果 $\textit{mx}$ 比其余元素个数 $\textit{sum}- \textit{mx}$ 还多，那么操作次数为其余元素个数 $\textit{sum}- \textit{mx}$。否则操作次数为 $\left\lfloor\dfrac{sum}{2}\right\rfloor$。

[证明+具体构造方案](https://leetcode.cn/problems/reorganize-string/solutions/2779462/tan-xin-gou-zao-pai-xu-bu-pai-xu-liang-c-h9jg/)

这个结论可以简化，操作次数等于

$$
\min\left(\left\lfloor\dfrac{sum}{2}\right\rfloor, \textit{sum}- \textit{mx}\right)
$$

## 有 xx 的情况

由于 $\texttt{xx}$ 的个数不超过 $\textit{cards}$ 的长度，我们可以枚举分配多少个 $\texttt{xx}$ 给第一组，其余的 $\texttt{xx}$ 给第二组，然后用上面的公式计算答案，取最大值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1aCaGzWEm4/?t=1m20s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    # 计算除了 x 以外的出现次数之和 sum_cnt，出现次数最大值 max_cnt
    def get_sum_and_max(self, cnt: Counter, x: str) -> Tuple[int, int]:
        del cnt[x]
        sum_cnt = sum(cnt.values())
        max_cnt = max(cnt.values(), default=0)
        return sum_cnt, max_cnt

    # 计算这一组在得到 k 个 xx 后的得分
    def calc_score(self, s: int, mx: int, k: int) -> int:
        s += k
        mx = max(mx, k)
        return min(s // 2, s - mx)

    def score(self, cards: List[str], x: str) -> int:
        cnt1 = Counter(b for a, b in cards if a == x)  # 统计 "x?" 中的 ? 的出现次数
        cnt2 = Counter(a for a, b in cards if b == x)  # 统计 "?x" 中的 ? 的出现次数

        cnt_xx = cnt1[x]
        sum1, max1 = self.get_sum_and_max(cnt1, x)
        sum2, max2 = self.get_sum_and_max(cnt2, x)

        ans = 0
        # 枚举分配 k 个 xx 给第一组，其余的 xx 给第二组
        for k in range(cnt_xx + 1):
            score1 = self.calc_score(sum1, max1, k)
            score2 = self.calc_score(sum2, max2, cnt_xx - k)
            ans = max(ans, score1 + score2)
        return ans
```

```java [sol-Java]
class Solution {
    public int score(String[] cards, char x) {
        int[] cnt1 = new int[10]; // 题目保证只有前 10 个小写字母
        int[] cnt2 = new int[10];
        for (String s : cards) {
            // 统计形如 x? 的每个 ? 的出现次数
            char c0 = s.charAt(0);
            char c1 = s.charAt(1);
            if (c0 == x) {
                cnt1[c1 - 'a']++;
            }
            // 统计形如 ?x 的每个 ? 的出现次数
            if (c1 == x) {
                cnt2[c0 - 'a']++;
            }
        }

        int[] res1 = getSumAndMax(cnt1, x);
        int[] res2 = getSumAndMax(cnt2, x);
        int sum1 = res1[0], max1 = res1[1];
        int sum2 = res2[0], max2 = res2[1];

        int cntXX = cnt1[x - 'a'];
        int ans = 0;
        // 枚举分配 k 个 xx 给第一组，其余的 xx 给第二组
        for (int k = 0; k <= cntXX; k++) {
            int score1 = calcScore(sum1, max1, k);
            int score2 = calcScore(sum2, max2, cntXX - k);
            ans = Math.max(ans, score1 + score2);
        }
        return ans;
    }

    // 计算除了 x 以外的出现次数之和 sum，出现次数最大值 mx
    private int[] getSumAndMax(int[] cnt, char x) {
        int sum = 0, mx = 0;
        for (int i = 0; i < cnt.length; i++) {
            if (i != x - 'a') {
                sum += cnt[i];
                mx = Math.max(mx, cnt[i]);
            }
        }
        return new int[]{sum, mx};
    }

    // 计算这一组在得到 k 个 xx 后的得分
    private int calcScore(int sum, int mx, int k) {
        sum += k;
        mx = Math.max(mx, k);
        return Math.min(sum / 2, sum - mx);
    }
}
```

```cpp [sol-C++]
class Solution {
    // 计算除了 x 以外的出现次数之和 sum，出现次数最大值 mx
    pair<int, int> get_sum_and_max(const array<int, 10>& cnt, char x) {
        int sum = 0, mx = 0;
        for (int i = 0; i < cnt.size(); i++) {
            if (i != x - 'a') {
                sum += cnt[i];
                mx = max(mx, cnt[i]);
            }
        }
        return {sum, mx};
    }

    // 计算这一组在得到 k 个 xx 后的得分
    int calc_score(int sum, int mx, int k) {
        sum += k;
        mx = max(mx, k);
        return min(sum / 2, sum - mx);
    }

public:
    int score(vector<string>& cards, char x) {
        array<int, 10> cnt1{}, cnt2{}; // 题目保证只有前 10 个小写字母
        for (auto& s : cards) {
            // 统计形如 x? 的每个 ? 的出现次数
            if (s[0] == x) {
                cnt1[s[1] - 'a']++;
            }
            // 统计形如 ?x 的每个 ? 的出现次数
            if (s[1] == x) {
                cnt2[s[0] - 'a']++;
            }
        }

        auto [sum1, max1] = get_sum_and_max(cnt1, x);
        auto [sum2, max2] = get_sum_and_max(cnt2, x);

        int cnt_xx = cnt1[x - 'a'];
        int ans = 0;
        // 枚举分配 k 个 xx 给第一组，其余的 xx 给第二组
        for (int k = 0; k <= cnt_xx; k++) {
            int score1 = calc_score(sum1, max1, k);
            int score2 = calc_score(sum2, max2, cnt_xx - k);
            ans = max(ans, score1 + score2);
        }
        return ans;
    }
};
```

```go [sol-Go]
// 计算除了 x 以外的出现次数之和 sum，出现次数最大值 mx
func getSumAndMax(cnt []int, x byte) (sum, mx int) {
	for i, c := range cnt {
		if i != int(x-'a') {
			sum += c
			mx = max(mx, c)
		}
	}
	return
}

// 计算这一组在得到 k 个 xx 后的得分
func calcScore(sum, mx, k int) int {
	sum += k
	mx = max(mx, k)
	return min(sum/2, sum-mx)
}

func score(cards []string, x byte) (ans int) {
	var cnt1, cnt2 [10]int // 题目保证只有前 10 个小写字母
	for _, s := range cards {
		// 统计形如 x? 的每个 ? 的出现次数
		if s[0] == x {
			cnt1[s[1]-'a']++
		}
		// 统计形如 ?x 的每个 ? 的出现次数
		if s[1] == x {
			cnt2[s[0]-'a']++
		}
	}

	sum1, max1 := getSumAndMax(cnt1[:], x)
	sum2, max2 := getSumAndMax(cnt2[:], x)

	cntXX := cnt1[x-'a']
	// 枚举分配 k 个 xx 给第一组，其余的 xx 给第二组
	for k := range cntXX + 1 {
		score1 := calcScore(sum1, max1, k)
		score2 := calcScore(sum2, max2, cntXX-k)
		ans = max(ans, score1+score2)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $\textit{cards}$ 的长度，$|\Sigma|=10$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 优化

可以不枚举 $\texttt{xx}$ 的个数，直接计算。

首先计算不考虑 $\texttt{xx}$ 时：

- 第一组的得分（配对个数）$\textit{pairs}_1$，以及第一组配对后的剩余字符串个数 $\textit{left}_1$。
- 第二组的得分（配对个数）$\textit{pairs}_2$，以及第二组配对后的剩余字符串个数 $\textit{left}_2$。

然后考虑 $\texttt{xx}$ 的影响。

首先，把 $\texttt{xx}$ 和剩余的字符串配对，每有 $1$ 个 $\texttt{xx}$，得分就能增加一，但这不能超过剩下的字符串个数 $\textit{left}_1 +\textit{left}_2$。

然后，如果还有 $\texttt{xx}$，就撤销之前的配对，比如 $(\texttt{ax},\texttt{bx})$ 改成 $(\texttt{ax},\texttt{xx})$ 和 $(\texttt{bx},\texttt{xx})$。这样每有 $2$ 个 $\texttt{xx}$，得分就能增加一，但这不能超过之前的配对个数 $\textit{pairs}_1+\textit{pairs}_2$。由于这种方案平均每个 $\texttt{xx}$ 只能增加 $0.5$ 分，不如上面的，所以先考虑把 $\texttt{xx}$ 和剩余的字符串配对，再考虑撤销之前的配对。

```py [sol-Python3]
class Solution:
    # 计算这一组的得分（配对个数），以及剩余元素个数
    def calc(self, cnt: Counter, x: str) -> (int, int):
        del cnt[x]
        sum_cnt = sum(cnt.values())
        max_cnt = max(cnt.values(), default=0)
        pairs = min(sum_cnt // 2, sum_cnt - max_cnt)
        return pairs, sum_cnt - pairs * 2

    def score(self, cards: List[str], x: str) -> int:
        cnt1 = Counter(b for a, b in cards if a == x)
        cnt2 = Counter(a for a, b in cards if b == x)

        cnt_xx = cnt1[x]
        pairs1, left1 = self.calc(cnt1, x)
        pairs2, left2 = self.calc(cnt2, x)
        ans = pairs1 + pairs2  # 不考虑 xx 时的得分

        # 把 xx 和剩下的 x? 和 ?x 配对
        # 每有 1 个 xx，得分就能增加一，但这不能超过剩下的 x? 和 ?x 的个数 left1+left2
        if cnt_xx > 0:
            mn = min(cnt_xx, left1 + left2)
            ans += mn
            cnt_xx -= mn

        # 如果还有 xx，就撤销之前的配对，比如 (ax,bx) 改成 (ax,xx) 和 (bx,xx)
        # 每有 2 个 xx，得分就能增加一，但这不能超过之前的配对个数 pairs1+pairs2
        # 由于这种方案平均每个 xx 只能增加 0.5 分，不如上面的，所以先考虑把 xx 和剩下的 x? 和 ?x 配对，再考虑撤销之前的配对
        if cnt_xx > 0:
            ans += min(cnt_xx // 2, pairs1 + pairs2)

        return ans
```

```java [sol-Java]
class Solution {
    // 计算这一组的得分（配对个数），以及剩余元素个数
    private int[] calc(int[] cnt, char x) {
        int sum = 0, mx = 0;
        for (int i = 0; i < cnt.length; i++) {
            if (i != x - 'a') {
                sum += cnt[i];
                mx = Math.max(mx, cnt[i]);
            }
        }
        int pairs = Math.min(sum / 2, sum - mx);
        return new int[]{pairs, sum - pairs * 2};
    }

    public int score(String[] cards, char x) {
        int[] cnt1 = new int[10];
        int[] cnt2 = new int[10];
        for (String s : cards) {
            char c0 = s.charAt(0);
            char c1 = s.charAt(1);
            if (c0 == x) {
                cnt1[c1 - 'a']++;
            }
            if (c1 == x) {
                cnt2[c0 - 'a']++;
            }
        }

        int[] res1 = calc(cnt1, x);
        int[] res2 = calc(cnt2, x);
        int pairs1 = res1[0], left1 = res1[1];
        int pairs2 = res2[0], left2 = res2[1];
        int ans = pairs1 + pairs2; // 不考虑 xx 时的得分

        int cntXX = cnt1[x - 'a'];
        // 把 xx 和剩下的 x? 和 ?x 配对
        // 每有 1 个 xx，得分就能增加一，但这不能超过剩下的 x? 和 ?x 的个数 left1+left2
        if (cntXX > 0) {
            int mn = Math.min(cntXX, left1 + left2);
            ans += mn;
            cntXX -= mn;
        }

        // 如果还有 xx，就撤销之前的配对，比如 (ax,bx) 改成 (ax,xx) 和 (bx,xx)
        // 每有 2 个 xx，得分就能增加一，但这不能超过之前的配对个数 pairs1+pairs2
        // 由于这种方案平均每个 xx 只能增加 0.5 分，不如上面的，所以先考虑把 xx 和剩下的 x? 和 ?x 配对，再考虑撤销之前的配对
        if (cntXX > 0) {
            ans += Math.min(cntXX / 2, pairs1 + pairs2);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 计算这一组的得分（配对个数），以及剩余元素个数
    pair<int, int> calc(const array<int, 10>& cnt, char x) {
        int sum = 0, mx = 0;
        for (int i = 0; i < cnt.size(); i++) {
            if (i != x - 'a') {
                sum += cnt[i];
                mx = max(mx, cnt[i]);
            }
        }
        int pairs = min(sum / 2, sum - mx);
        return {pairs, sum - pairs * 2};
    }

public:
    int score(vector<string>& cards, char x) {
        array<int, 10> cnt1{}, cnt2{};
        for (auto& s : cards) {
            if (s[0] == x) {
                cnt1[s[1] - 'a']++;
            }
            if (s[1] == x) {
                cnt2[s[0] - 'a']++;
            }
        }

        auto [pairs1, left1] = calc(cnt1, x);
        auto [pairs2, left2] = calc(cnt2, x);
        int ans = pairs1 + pairs2; // 不考虑 xx 时的得分

        int cnt_xx = cnt1[x - 'a'];
        // 把 xx 和剩下的 x? 和 ?x 配对
        // 每有 1 个 xx，得分就能增加一，但这不能超过剩下的 x? 和 ?x 的个数 left1+left2
        if (cnt_xx > 0) {
            int mn = min(cnt_xx, left1 + left2);
            ans += mn;
            cnt_xx -= mn;
        }

        // 如果还有 xx，就撤销之前的配对，比如 (ax,bx) 改成 (ax,xx) 和 (bx,xx)
        // 每有 2 个 xx，得分就能增加一，但这不能超过之前的配对个数 pairs1+pairs2
        // 由于这种方案平均每个 xx 只能增加 0.5 分，不如上面的，所以先考虑把 xx 和剩下的 x? 和 ?x 配对，再考虑撤销之前的配对
        if (cnt_xx > 0) {
            ans += min(cnt_xx / 2, pairs1 + pairs2);
        }

        return ans;
    }
};
```

```go [sol-Go]
// 计算这一组的得分（配对个数），以及剩余元素个数
func calc(cnt []int, x byte) (int, int) {
	sum, mx := 0, 0
	for i, c := range cnt {
		if i != int(x-'a') {
			sum += c
			mx = max(mx, c)
		}
	}
	pairs := min(sum/2, sum-mx)
	return pairs, sum - pairs*2
}

func score(cards []string, x byte) int {
	var cnt1, cnt2 [10]int
	for _, s := range cards {
		if s[0] == x {
			cnt1[s[1]-'a']++
		}
		if s[1] == x {
			cnt2[s[0]-'a']++
		}
	}

	pairs1, left1 := calc(cnt1[:], x)
	pairs2, left2 := calc(cnt2[:], x)
	ans := pairs1 + pairs2 // 不考虑 xx 时的得分

	cntXX := cnt1[x-'a']
	// 把 xx 和剩下的 x? 和 ?x 配对
	// 每有 1 个 xx，得分就能增加一，但这不能超过剩下的 x? 和 ?x 的个数 left1+left2
	if cntXX > 0 {
		mn := min(cntXX, left1+left2)
		ans += mn
		cntXX -= mn
	}

	// 如果还有 xx，就撤销之前的配对，比如 (ax,bx) 改成 (ax,xx) 和 (bx,xx)
	// 每有 2 个 xx，得分就能增加一，但这不能超过之前的配对个数 pairs1+pairs2
	// 由于这种方案平均每个 xx 只能增加 0.5 分，不如上面的，所以先考虑把 xx 和剩下的 x? 和 ?x 配对，再考虑撤销之前的配对
	if cntXX > 0 {
		ans += min(cntXX/2, pairs1+pairs2)
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $\textit{cards}$ 的长度，$|\Sigma|=10$ 是字符集合的大小。瓶颈在遍历 $\textit{cards}$ 上。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 专题训练

见下面贪心与思维题单的「**§1.8 相邻不同**」。

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
