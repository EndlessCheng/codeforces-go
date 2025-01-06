## 方法一：记忆化搜索+记录字符集合

定义 $\textit{dfs}(i,\textit{mask}, \textit{changed})$ 表示当前遍历到 $s[i]$，当前这一段在 $i$ 之前的字符集合是 $\textit{mask}$，是否已经修改了字符（$\textit{changed}$），后续可以得到的最大分割数。

讨论是否修改 $s[i]$，以及当前字母能否加到 $\textit{mask}$ 中：

- 如果不改 $s[i]$：
   - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $s[i]$ 必须划分到下一段子串中。答案为 $\textit{dfs}(i+1, \{s[i]\},\textit{changed}) + 1$。
   - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $s[i]$ 必须在当前这一段中。答案为 $\textit{dfs}(i+1, \textit{mask}\cup \{s[i]\},\textit{changed})$。
- 如果 $\textit{changed}=\texttt{false}$，那么可以改 $s[i]$，枚举改成第 $j$ 个字母。
    - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $j$ 必须划分到下一段子串中。答案为 $\textit{dfs}(i+1, \{j\},\texttt{true}) + 1$。
    - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $j$ 必须在当前这一段中。答案为 $\textit{dfs}(i+1, \textit{mask}\cup \{j\},\texttt{true})$。

这些情况取最大值，就得到了 $\textit{dfs}(i,\textit{mask}, \textit{changed})$。

递归边界：$\textit{dfs}(n,*,*) = 1$。注意当 $i>0$ 时，$\textit{mask}\ne 0$，表示一段子串的字符集合。所以递归到 $i=n$ 时，$\textit{mask}$ 就是最后一段的字符集合了，返回 $1$。

递归入口：$\textit{dfs}(0,0,\texttt{false})$，也就是答案。

代码实现时，用二进制表示集合，用位运算实现集合相关操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        @cache
        def dfs(i: int, mask: int, changed: bool) -> int:
            if i == len(s):
                return 1

            # 不改 s[i]
            bit = 1 << (ord(s[i]) - ord('a'))
            new_mask = mask | bit
            if new_mask.bit_count() > k:
                # 分割出一个子串，这个子串的最后一个字母在 i-1
                # s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
                res = dfs(i + 1, bit, changed) + 1
            else:  # 不分割
                res = dfs(i + 1, new_mask, changed)
            if changed:
                return res

            # 枚举把 s[i] 改成 a,b,c,...,z
            for j in range(26):
                new_mask = mask | (1 << j)
                if new_mask.bit_count() > k:
                    # 分割出一个子串，这个子串的最后一个字母在 i-1
                    # j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
                    res = max(res, dfs(i + 1, 1 << j, True) + 1)
                else:  # 不分割
                    res = max(res, dfs(i + 1, new_mask, True))
            return res

        return dfs(0, 0, False)
```

```java [sol-Java]
class Solution {
    private final Map<Long, Integer> memo = new HashMap<>();

    public int maxPartitionsAfterOperations(String s, int k) {
        return dfs(0, 0, 0, s.toCharArray(), k);
    }

    private int dfs(int i, int mask, int changed, char[] s, int k) {
        if (i == s.length) {
            return 1;
        }

        long argsMask = (long) i << 32 | mask << 1 | changed;
        if (memo.containsKey(argsMask)) { // 之前计算过
            return memo.get(argsMask);
        }

        int res;
        // 不改 s[i]
        int bit = 1 << (s[i] - 'a');
        int newMask = mask | bit;
        if (Integer.bitCount(newMask) > k) {
            // 分割出一个子串，这个子串的最后一个字母在 i-1
            // s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
            res = dfs(i + 1, bit, changed, s, k) + 1;
        } else { // 不分割
            res = dfs(i + 1, newMask, changed, s, k);
        }

        if (changed == 0) {
            // 枚举把 s[i] 改成 a,b,c,...,z
            for (int j = 0; j < 26; j++) {
                newMask = mask | (1 << j);
                if (Integer.bitCount(newMask) > k) {
                    // 分割出一个子串，这个子串的最后一个字母在 i-1
                    // j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
                    res = Math.max(res, dfs(i + 1, 1 << j, 1, s, k) + 1);
                } else { // 不分割
                    res = Math.max(res, dfs(i + 1, newMask, 1, s, k));
                }
            }
        }

        memo.put(argsMask, res); // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPartitionsAfterOperations(string s, int k) {
        unordered_map<long long, int> memo;
        auto dfs = [&](this auto&& dfs, int i, int mask, bool changed) -> int {
            if (i == s.length()) {
                return 1;
            }

            long long args_mask = (long long) i << 32 | mask << 1 | changed;
            auto it = memo.find(args_mask);
            if (it != memo.end()) { // 之前计算过
                return it->second;
            }

            int res;
            // 不改 s[i]
            int bit = 1 << (s[i] - 'a');
            int new_mask = mask | bit;
            if (__builtin_popcount(new_mask) > k) {
                // 分割出一个子串，这个子串的最后一个字母在 i-1
                // s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
                res = dfs(i + 1, bit, changed) + 1;
            } else { // 不分割
                res = dfs(i + 1, new_mask, changed);
            }

            if (!changed) {
                // 枚举把 s[i] 改成 a,b,c,...,z
                for (int j = 0; j < 26; j++) {
                    new_mask = mask | (1 << j);
                    if (__builtin_popcount(new_mask) > k) {
                        // 分割出一个子串，这个子串的最后一个字母在 i-1
                        // j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
                        res = max(res, dfs(i + 1, 1 << j, true) + 1);
                    } else { // 不分割
                        res = max(res, dfs(i + 1, new_mask, true));
                    }
                }
            }

            return memo[args_mask] = res; // 记忆化
        };
        return dfs(0, 0, false);
    }
};
```

```go [sol-Go]
func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	type args struct {
		i, mask int
		changed bool
	}
	memo := map[args]int{}
	var dfs func(int, int, bool) int
	dfs = func(i, mask int, changed bool) (res int) {
		if i == n {
			return 1
		}

		a := args{i, mask, changed}
		if v, ok := memo[a]; ok { // 之前计算过
			return v
		}

		// 不改 s[i]
		bit := 1 << (s[i] - 'a')
		newMask := mask | bit
		if bits.OnesCount(uint(newMask)) > k {
			// 分割出一个子串，这个子串的最后一个字母在 i-1
			// s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
			res = dfs(i+1, bit, changed) + 1
		} else { // 不分割
			res = dfs(i+1, newMask, changed)
		}

		if !changed {
			// 枚举把 s[i] 改成 a,b,c,...,z
			for j := 0; j < 26; j++ {
				newMask := mask | 1<<j
				if bits.OnesCount(uint(newMask)) > k {
					// 分割出一个子串，这个子串的最后一个字母在 i-1
					// j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
					res = max(res, dfs(i+1, 1<<j, true)+1)
				} else { // 不分割
					res = max(res, dfs(i+1, newMask, true))
				}
			}
		}

		memo[a] = res // 记忆化
		return res
	}
	return dfs(0, 0, false)
}
```

#### 复杂度分析

有多少个状态呢？可能你会觉得非常多，但是我们可以换个角度看：考虑到 $i$ 为止，有多少种不同的字符集合，即有多少个不同的 $\textit{mask}$。

- 如果没有修改，那么 $\textit{mask}$ 是唯一的。
- 如果中途有修改，我们可以从 $i$ 开始向左扩展，每遇到一个字符，$\textit{mask}$ 集合要么不变，要么增加一个字符。那么有 $\mathcal{O}(|\Sigma|)$ 个本质不同的修改位置，每个位置有 $\mathcal{O}(|\Sigma|)$ 个不同的修改方式，所以有 $\mathcal{O}(|\Sigma|^2)$ 个不同的 $\textit{mask}$。这里 $|\Sigma|=26$。

所以只有 $\mathcal{O}(n|\Sigma|^2)$ 个状态。

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。分类讨论：
   - 如果 $\textit{mask}$ 之前没有修改，这样的状态有 $\mathcal{O}(n)$ 个，单个状态的计算时间为 $\mathcal{O}(|\Sigma|)$，即枚举修改成什么字母的时间。
   - 如果 $\textit{mask}$ 之前有修改，这样的状态有 $\mathcal{O}(n|\Sigma|^2)$ 个，单个状态的计算时间为 $\mathcal{O}(1)$，因为我们只能不修改。
   - 所以时间复杂度为 $\mathcal{O}(n|\Sigma|^2)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|^2)$。

## 方法二：前后缀分解

### 提示 1

能否从右往左分割？

我们需要证明，对于 $s$ 的任意**后缀**，从左往右分割出的段数，等于从右往左分割出的段数。

考虑 $s$ 的某个后缀，假设从左往右分成了 $n$ 段，从左往右分别记作 $L_1,L_2,\cdots,L_n$，每一段最左边的字母下标分别记作 $p_1,p_2,\cdots,p_n$。再假设从右往左分成了 $m$ 段，**从右往左**分别记作 $R_1,R_2,\cdots,R_m$，每一段最左边（终点）的字母下标分别记作 $q_1,q_2,\cdots,q_m$，最右边（起点）的字母下标分别记作 $r_1,r_2,\cdots,r_m$。

![w379d.png](https://pic.leetcode.cn/1704712474-pevDnh-w379d.png)

对于 $R_1$，显然它包含 $L_n$，所以 $R_1$ 的终点 $q_1\le p_n$。如果 $q_1\le p_{n-1}$，那么 $R_1$ 会包含 $L_{n-1}$ 和 $L_n$，这是不可能的，因为这两段的字符种类是大于 $k$ 的。所以 $q_1 > p_{n-1}$。

这意味着对于 $R_2$，它的起点 $r_2$ 在 $L_{n-1}$ 中。同样地，$R_2$ 的终点 $q_2$ 必须大于 $p_{n-2}$，否则 $R_2$ 除了完整地包含 $L_{n-2}$，还包含 $L_{n-1}$ 的字母，这会让 $R_2$ 的字母种类数超过 $k$。

依此类推，每个 $R_i$ 的起点 $r_i$ 都在 $L_{n+1-i}$ 中。而最后一段 $R_m$ 的起点 $r_m$ 在 $L_1$ 中。这意味着 $n=m$。所以从左往右分割出的段数，等于从右往左分割出的段数。

### 提示 2

遍历 $s$，枚举要修改的字母 $s[i]$。

设从左往右分割到 $i-1$ 时，分割出了 $\textit{preSeg}$ 段，最新一段（记作 $L$）的字符集合为 $\textit{preMask}$，其大小为 $\textit{preSize}$。

设从右往左分割到 $i+1$ 时，分割出了 $\textit{sufSeg}$ 段，最新一段（记作 $R$）的字符集合为 $\textit{sufMask}$，其大小为 $\textit{sufSize}$。

设 $\textit{preMask}$ 和 $\textit{sufMask}$ 的并集大小为 $\textit{unionSize}$。

分类讨论：

- 情况 1：如果 $\textit{unionSize}<k$，那么无论把 $s[i]$ 改成什么，并集的大小都不会超过 $k$，这意味着 $L$ 和 $R$ 必须合并成一段。
- 情况 2：如果 $\textit{unionSize}<26$，并且 $\textit{preSize}$ 和 $\textit{sufSize}$ 都等于 $k$，那么把 $s[i]$ 改成一个既不在 $L$ 又不在 $R$ 中的字母，从而使 $L$ 和 $R$ 都是最大分割，这样就多出了一段。
- 情况 3：不是情况 1 也不是情况 2，那么总段数不变。

代码实现时，可以从右往左遍历 $s$，把 $\textit{sufSeg}$ 和 $\textit{sufMask}$ 都记录下来。然后从左往右遍历 $s$，在计算 $\textit{preSeg}$ 和 $\textit{preMask}$ 的同时，按照上述分类讨论，计算出修改后的段数，更新答案的最大值。

```py [sol-Python3]
class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        if k == 26:
            return 1

        seg, mask, size = 1, 0, 0
        def update(i: int) -> None:
            nonlocal seg, mask, size
            bit = 1 << (ord(s[i]) - ord('a'))
            if mask & bit:
                return
            size += 1
            if size > k:
                seg += 1  # s[i] 在新的一段中
                mask = bit
                size = 1
            else:
                mask |= bit

        n = len(s)
        suf = [None] * n + [(0, 0)]
        for i in range(n - 1, -1, -1):
            update(i)
            suf[i] = (seg, mask)

        ans = seg  # 不修改任何字母
        seg, mask, size = 1, 0, 0
        for i in range(n):
            suf_seg, suf_mask = suf[i + 1]
            res = seg + suf_seg  # 情况 3
            union_size = (mask | suf_mask).bit_count()
            if union_size < k:
                res -= 1  # 情况 1
            elif union_size < 26 and size == k and suf_mask.bit_count() == k:
                res += 1  # 情况 2
            ans = max(ans, res)
            update(i)
        return ans
```

```java [sol-Java]
class Solution {
    private int seg = 1, mask = 0, size = 0;

    public int maxPartitionsAfterOperations(String S, int k) {
        if (k == 26) {
            return 1;
        }

        char[] s = S.toCharArray();
        int n = s.length;
        int[] sufSeg = new int[n + 1];
        int[] sufMask = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            update(s[i], k);
            sufSeg[i] = seg;
            sufMask[i] = mask;
        }

        int ans = seg; // 不修改任何字母
        seg = 1; mask = 0; size = 0;
        for (int i = 0; i < n; i++) {
            int res = seg + sufSeg[i + 1]; // 情况 3
            int unionSize = Integer.bitCount(mask | sufMask[i + 1]);
            if (unionSize < k) {
                res--; // 情况 1
            } else if (unionSize < 26 && size == k && Integer.bitCount(sufMask[i + 1]) == k) {
                res++; // 情况 2
            }
            ans = Math.max(ans, res);
            update(s[i], k);
        }
        return ans;
    }

    private void update(char c, int k) {
        int bit = 1 << (c - 'a');
        if ((mask & bit) != 0) {
            return;
        }
        if (++size > k) {
            seg++; // c 在新的一段中
            mask = bit;
            size = 1;
        } else {
            mask |= bit;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPartitionsAfterOperations(string s, int k) {
        if (k == 26) {
            return 1;
        }

        int seg = 1, mask = 0, size = 0;
        auto update = [&](int i) {
            int bit = 1 << (s[i] - 'a');
            if (mask & bit) return;
            if (++size > k) {
                seg++; // s[i] 在新的一段中
                mask = bit;
                size = 1;
            } else {
                mask |= bit;
            }
        };

        int n = s.length();
        vector<pair<int, int>> suf(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            update(i);
            suf[i] = {seg, mask};
        }

        int ans = seg; // 不修改任何字母
        seg = 1, mask = 0, size = 0;
        for (int i = 0; i < n; i++) {
            auto [suf_seg, suf_mask] = suf[i + 1];
            int res = seg + suf_seg; // 情况 3
            int union_size = __builtin_popcount(mask | suf_mask);
            if (union_size < k) {
                res--; // 情况 1
            } else if (union_size < 26 && size == k && __builtin_popcount(suf_mask) == k) {
                res++; // 情况 2
            }
            ans = max(ans, res);
            update(i);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPartitionsAfterOperations(s string, k int) int {
	if k == 26 {
		return 1
	}

	seg, mask, size := 1, 0, 0
	update := func(i int) {
		bit := 1 << (s[i] - 'a')
		if mask&bit > 0 {
			return
		}
		size++
		if size > k {
			seg++ // s[i] 在新的一段中
			mask = bit
			size = 1
		} else {
			mask |= bit
		}
	}

	n := len(s)
	type pair struct{ seg, mask int }
	suf := make([]pair, n+1)
	for i := n - 1; i >= 0; i-- {
		update(i)
		suf[i] = pair{seg, mask}
	}

	ans := seg // 不修改任何字母
	seg, mask, size = 1, 0, 0
	for i := range s {
		p := suf[i+1]
		res := seg + p.seg // 情况 3
		unionSize := bits.OnesCount(uint(mask | p.mask))
		if unionSize < k {
			res-- // 情况 1
		} else if unionSize < 26 && size == k && bits.OnesCount(uint(p.mask)) == k {
			res++ // 情况 2
		}
		ans = max(ans, res)
		update(i)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
