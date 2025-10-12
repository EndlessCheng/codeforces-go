## 写法一：暴力枚举

枚举子串左端点 $i$，然后枚举子串右端点 $j=i,i+1,i+2,\ldots,n-1$。

在枚举右端点 $j$ 的过程中，统计子串 $[i,j]$ 每种字母的出现次数 $\textit{cnt}$。

遍历 $\textit{cnt}$，如果所有字母的出现次数均相同，用子串长度 $j-i+1$ 更新答案的最大值。

[本题视频讲解](https://www.bilibili.com/video/BV1FJ4uz1EkN/?t=1m20s)，欢迎点赞关注~

### 优化前

```py [sol-Python3]
class Solution:
    def longestBalanced(self, s: str) -> int:
        ans = 0
        n = len(s)
        for i in range(n):
            cnt = defaultdict(int)
            for j in range(i, n):
                cnt[s[j]] += 1
                if len(set(cnt.values())) == 1:
                    ans = max(ans, j - i + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestBalanced(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int[] cnt = new int[26];
            next:
            for (int j = i; j < n; j++) {
                int base = ++cnt[s[j] - 'a'];
                for (int c : cnt) {
                    if (c > 0 && c != base) {
                        continue next;
                    }
                }
                ans = Math.max(ans, j - i + 1);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int cnt[26]{};
            for (int j = i; j < n; j++) {
                int base = ++cnt[s[j] - 'a'];
                for (int c : cnt) {
                    if (c && c != base) {
                        base = -1;
                        break;
                    }
                }
                if (base != -1) {
                    ans = max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	for i := range s {
		cnt := make([]int, 26)
	next:
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			base := cnt[s[j]-'a']
			for _, c := range cnt {
				if c > 0 && c != base {
					continue next
				}
			}
			ans = max(ans, j-i+1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

### 优化

设 $\textit{mx} = \max(\textit{cnt})$，设 $\textit{kinds}$ 为子串中的不同字母个数。

如果 $\textit{mx}\cdot \textit{kinds} = j-i+1$，说明子串所有字母的出现次数均为 $\textit{mx}$，均相等。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def longestBalanced(self, s: str) -> int:
        ans = 0
        for i in range(len(s)):
            cnt = defaultdict(int)
            mx = 0
            for j in range(i, len(s)):
                cnt[s[j]] += 1
                mx = max(mx, cnt[s[j]])
                if mx * len(cnt) == j - i + 1:
                    ans = max(ans, j - i + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestBalanced(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int[] cnt = new int[26];
            int mx = 0, kinds = 0;
            for (int j = i; j < n; j++) {
                int b = s[j] - 'a';
                if (cnt[b] == 0) {
                    kinds++;
                }
                mx = Math.max(mx, ++cnt[b]);
                if (mx * kinds == j - i + 1) {
                    ans = Math.max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int cnt[26]{};
            int mx = 0, kinds = 0;
            for (int j = i; j < n; j++) {
                int b = s[j] - 'a';
                if (cnt[b] == 0) {
                    kinds++;
                }
                mx = max(mx, ++cnt[b]);
                if (mx * kinds == j - i + 1) {
                    ans = max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		mx, kinds := 0, 0
		for j := i; j < len(s); j++ {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				kinds++
			}
			cnt[b]++
			mx = max(mx, cnt[b])
			if mx*kinds == j-i+1 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|=26$ 是字符集合的大小。

## 写法二：枚举右，维护左

推荐先完成 [3714. 最长的平衡子串 II](https://leetcode.cn/problems/longest-balanced-substring-ii/)，再来理解这个做法。

能否只用一个哈希表，同时解决子串包含 $1$ 个、$2$ 个……$26$ 个字母的情况？

定义 $S[i][j]$ 表示前缀 $[0,i]$ 中的字母 $j$ 的出现次数。$S[-1][j] = 0$。

如果要找只包含字母 $\texttt{abc}$ 的平衡子串 $(l,r]$，3714 题告诉我们：

- 用 $S[r][1] - S[r][0] = S[l][1] - S[l][0]$ 判断子串中的 $\texttt{a}$ 和 $\texttt{b}$ 的个数是否相等。
- 用 $S[r][2] - S[r][0] = S[l][2] - S[l][0]$ 判断子串中的 $\texttt{a}$ 和 $\texttt{c}$ 的个数是否相等。
- 此外，子串不能包含 $\ge \texttt{d}$ 的字母，即 $S[r][j] - S[l][j] = 0\ (j\ge 3)$，也就是 $S[r][j] = S[l][j]\ (j\ge 3)$。

据此，定义

$$
d[i][j] =
\begin{cases}
S[i][j] - S[i][\textit{minCh}], & j\ 在子串中     \\
S[i][j], & j\ 不在子串中     \\
\end{cases}
$$

其中 $\textit{minCh}$ 是子串中的最小字母，用来作为减法的基准。

然而，这个定义面临一个尴尬的问题：

- 我怎么知道 $j$ 是否在子串中？

必须先知道子串包含哪些字母，才能准确地算出 $d[i][j]$。

难道要像 3714 题那样，枚举所有非空字母子集，即 $2^{26}-1$ 种情况？

实际上，考虑以 $i$ 为右端点的子串，当子串左端点从 $i$ 开始向左移动（扩展）时，子串中的字母种类数要么不变，要么加一，所以（对于固定的 $i$ 来说）只有至多 $26$ 种字母集合。（这有点像 [LogTrick](https://zhuanlan.zhihu.com/p/1933215367158830792)）

于是，对于每个右端点 $i$，我们至多枚举 $26$ 种字母集合。

对于一个固定的字母集合，$d[i][j]$ 的值就是固定的了。问题相当于：

- 找到一对距离最远的 $(d[l],d[i])$，满足 $d[l] = d[i]$。
- 用 $i-l$ 更新答案的最大值。

现在，「枚举右维护左」中的「枚举右」解决了，「维护左」怎么做？

对于子串 $(l,r]$，我们需要：

- 枚举以 $l+1$ 为左端点的字母集合（至多 $26$ 种）。
- 计算 $d[l]$，作为哈希表的 key。哈希表的 value 为 $l$。

如果 $d[l] = d[r]$，那么子串 $(l,r]$ 就是平衡子串吗？

不一定。存在 $d[i]$ 相同，但字母集合不同的情况。所以我们还需要在哈希表的 key 中添加一个 $\textit{mask}$，表示字母集合（实现时用 [二进制数](https://leetcode.cn/circle/discuss/CaOJ45/) 压缩表示）。

```py [sol-Python3]
class Solution:
    def longestBalanced(self, s: str) -> int:
        s = [ord(c) - ord('a') for c in s]

        n = len(s)
        suf_orders = [None] * n
        order = []
        for i in range(n - 1, -1, -1):
            # 把最近出现的字母移到 order 末尾
            try: order.remove(s[i])
            except: pass
            order.append(s[i])
            suf_orders[i] = order[:]

        order = []
        cnt = [0] * 26
        pos = {}
        ans = 0
        for i, b in enumerate(s):
            suf_order = suf_orders[i]
            min_ch = inf
            mask = 0
            for j in range(len(suf_order) - 1, -1, -1):
                min_ch = min(min_ch, suf_order[j])
                # 注意此时 cnt 并不包含 s[i]，我们计算的是前缀 s[:i] 的信息
                # 在子串中的字母，计算差值
                # 不在子串中的字母，维持原样
                d = cnt[:]
                for ch in suf_order[j:]:
                    d[ch] -= cnt[min_ch]
                mask |= 1 << suf_order[j]
                p = (tuple(d), mask)  # mask 用来区分 d[ch] 是差值还是原始值
                # 记录 p 首次出现的位置
                if p not in pos:
                    pos[p] = i - 1

            # 把最近出现的字母移到 order 末尾
            try: order.remove(b)
            except: pass
            order.append(b)

            cnt[b] += 1
            min_ch = inf
            mask = 0
            for j in range(len(order) - 1, -1, -1):
                min_ch = min(min_ch, order[j])
                d = cnt[:]
                for ch in order[j:]:
                    d[ch] -= cnt[min_ch]
                mask |= 1 << order[j]
                p = (tuple(d), mask)
                # 再次遇到完全一样的 p，说明我们找到了一个平衡子串，左端点为 pos[p]+1，右端点为 i
                if p in pos:
                    ans = max(ans, i - pos[p])

        return ans
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	n := len(s)
	sufOrders := make([][]byte, n)
	order := []byte{}
	move := func(b byte) {
		// 把最近出现的字母 b 移到 order 末尾
		j := bytes.IndexByte(order, b)
		if j >= 0 {
			order = append(order[:j], order[j+1:]...)
		}
		order = append(order, b)
	}
	for i := n - 1; i >= 0; i-- {
		move(s[i] - 'a')
		sufOrders[i] = slices.Clone(order)
	}

	order = []byte{}
	cnt := [27]int{} // cnt[26] 作为 mask，用来区分 tmp[ch] 是差值还是原始值
	pos := map[[27]int]int{}
	for i, b := range s {
		sufOrder := sufOrders[i]
		minCh := byte(25)
		cnt[26] = 0
		for j := len(sufOrder) - 1; j >= 0; j-- {
			cnt[26] |= 1 << sufOrder[j]
			minCh = min(minCh, sufOrder[j])
			// 注意此时 cnt 并不包含 s[i]，我们计算的是前缀 s[:i] 的信息
			// 在子串中的字母，计算差值
			// 不在子串中的字母，维持原样
			d := cnt
			for _, ch := range sufOrder[j:] {
				d[ch] -= cnt[minCh]
			}
			// 记录 d 首次出现的位置
			if _, ok := pos[d]; !ok {
				pos[d] = i - 1
			}
		}

		// 把最近出现的字母移到 order 末尾
		move(byte(b - 'a'))

		cnt[b-'a']++
		minCh = byte(25)
		cnt[26] = 0
		for j := len(order) - 1; j >= 0; j-- {
			cnt[26] |= 1 << order[j]
			minCh = min(minCh, order[j])
			d := cnt
			for _, ch := range order[j:] {
				d[ch] -= cnt[minCh]
			}
			// 再次遇到完全一样的状态，说明找到了一个平衡子串，左端点为 l+1，右端点为 i
			if l, ok := pos[d]; ok {
				ans = max(ans, i-l)
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。

**注**：每次 $\textit{minCh}$ 变小时，我们都要重新算一遍 $\textit{tmp}$。如果计算 $\textit{tmp}$ 的多项式哈希值（代替哈希表的 key），我们可以 $\mathcal{O}(1)$ 计算哈希值的变化量，从而做到 $\mathcal{O}(n|\Sigma|)$ 时间。但该做法无法保证 100% 正确。

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
