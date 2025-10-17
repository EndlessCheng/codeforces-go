## 题意解读

我们要把 $s$ 分割成若干段。

「选择 $s$ 的最长前缀」的意思是，不能随意分割，每一段都要尽可能地长。

所以**在不修改字母的情况下，分割方案是唯一的**。

示例 1 的 $s = \texttt{accca}$，$k=2$。如果不修改字母，那么不能分割，只有一段。

在修改一处字母的情况下，为了让分割段数尽量多，最优做法是修改 $s[2]$ 为任意不等于 $\texttt{a}$ 和 $\texttt{c}$ 的字母，比如修改成 $\texttt{b}$，得到 $s'=\texttt{acbca}$。按照题目要求，$s'$ 的唯一分割方案为 $\texttt{ac} + \texttt{bc} + \texttt{a}$。

## 方法一：记忆化搜索

考虑递归枚举所有分割方案，即依次枚举 $s[0]$ 改成什么字母，$s[1]$ 改成什么字母，$s[2]$ 改成什么字母，依此类推。在递归过程中，我们需要跟踪如下信息：

- 当前位于 $s[i]$。
- 当前这段子串，包含哪些字母。不同字母的个数与 $k$ 比大小，如果超过 $k$，那么必须分割，即 $s[i]$ 是下一段的首字母。
- 之前是否修改过。这决定了我们能否修改 $s[i]$。

于是定义 $\textit{dfs}(i,\textit{mask}, \textit{changed})$ 表示当前遍历到 $s[i]$，在 $i$ 之前（不含 $i$）的字符集合是 $\textit{mask}$，之前是否修改过（$\textit{changed}$），返回在这一状态下，我们从剩余未分割的后缀中，可以得到的最大分割数。注意 $\textit{mask}$ 对应的那一段因为还不知道在哪结束，所以尚未分割，未计入分割数。

讨论是否修改 $s[i]$，以及当前字母能否加到 $\textit{mask}$ 中：

- 如果不改 $s[i]$：
   - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $s[i]$ 必须划分到下一段子串中。分割数为 $\textit{dfs}(i+1, \{s[i]\},\textit{changed}) + 1$。
   - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $s[i]$ 必须在当前这一段中。分割数为 $\textit{dfs}(i+1, \textit{mask}\cup \{s[i]\},\textit{changed})$。
- 如果 $\textit{changed}=\texttt{false}$，那么可以改 $s[i]$，枚举改成第 $j$ 个字母。
    - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $j$ 必须划分到下一段子串中。分割数为 $\textit{dfs}(i+1, \{j\},\texttt{true}) + 1$。
    - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $j$ 必须在当前这一段中。分割数为 $\textit{dfs}(i+1, \textit{mask}\cup \{j\},\texttt{true})$。

所有情况取最大值，就得到了 $\textit{dfs}(i,\textit{mask}, \textit{changed})$。

**递归边界**：$\textit{dfs}(n,*,*) = 1$。根据状态定义，$i=n$ 时 $\textit{mask}$ 对应的那一段尚未计入，所以递归到终点时返回的是 $1$ 而不是 $0$。

**递归入口**：$\textit{dfs}(0,0,\texttt{false})$，也就是答案。

代码实现时，用二进制表示集合，用位运算实现集合相关操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
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
    public int maxPartitionsAfterOperations(String s, int k) {
        Map<Long, Integer> memo = new HashMap<>();
        return dfs(0, 0, 0, memo, s.toCharArray(), k);
    }

    private int dfs(int i, int mask, int changed, Map<Long, Integer> memo, char[] s, int k) {
        if (i == s.length) {
            return 1;
        }

        // 把参数压缩到一个 long 中，方便作为哈希表的 key
        long args = (long) i << 32 | mask << 1 | changed;
        if (memo.containsKey(args)) { // 之前计算过
            return memo.get(args);
        }

        int res;
        // 不改 s[i]
        int bit = 1 << (s[i] - 'a');
        int newMask = mask | bit;
        if (Integer.bitCount(newMask) > k) {
            // 分割出一个子串，这个子串的最后一个字母在 i-1
            // s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
            res = dfs(i + 1, bit, changed, memo, s, k) + 1;
        } else { // 不分割
            res = dfs(i + 1, newMask, changed, memo, s, k);
        }

        if (changed == 0) {
            // 枚举把 s[i] 改成 a,b,c,...,z
            for (int j = 0; j < 26; j++) {
                newMask = mask | (1 << j);
                if (Integer.bitCount(newMask) > k) {
                    // 分割出一个子串，这个子串的最后一个字母在 i-1
                    // j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
                    res = Math.max(res, dfs(i + 1, 1 << j, 1, memo, s, k) + 1);
                } else { // 不分割
                    res = Math.max(res, dfs(i + 1, newMask, 1, memo, s, k));
                }
            }
        }

        memo.put(args, res); // 记忆化
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

            // 把参数压缩到一个 long long 中，方便作为哈希表的 key
            long long args = (long long) i << 32 | mask << 1 | changed;
            auto it = memo.find(args);
            if (it != memo.end()) { // 之前计算过
                return it->second;
            }

            int res;
            // 不改 s[i]
            int bit = 1 << (s[i] - 'a');
            int new_mask = mask | bit;
            if (popcount((uint32_t) new_mask) > k) {
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
                    if (popcount((uint32_t) new_mask) > k) {
                        // 分割出一个子串，这个子串的最后一个字母在 i-1
                        // j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
                        res = max(res, dfs(i + 1, 1 << j, true) + 1);
                    } else { // 不分割
                        res = max(res, dfs(i + 1, new_mask, true));
                    }
                }
            }

            return memo[args] = res; // 记忆化
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

有多少个状态呢？可能你会觉得非常多，因为 $\textit{mask}$ 至多有 $2^{26}$ 个不同的值。

但本题的约束是很强的，$\textit{mask}$ 必须对应一个**连续子串**（至多修改一次），而不是子序列。我们不能跳着选字母，如果不修改字母，$s = \texttt{abc}$ 不可能出现只包含 $\texttt{a}$ 和 $\texttt{c}$ 的集合，这样的状态是不存在的。

设子串的起点为 $j$，终点为 $i$。考虑当 $i$ 固定时，有多少个不同的 $\textit{mask}$。

- 当我们把左端点 $j$ 从 $i$ 向左移动（扩展）时，相当于把一个字符加入 $\textit{mask}$，那么 $\textit{mask}$ 集合要么不变，要么增加一个字符。这个过程只会产生 $\mathcal{O}(k)$ 个不同的 $\textit{mask}$。**注**：这让人想到了 [LogTrick](https://zhuanlan.zhihu.com/p/1933215367158830792)。
- 在 $\textit{mask}$ 集合不变的那一段，从中选择任意位置修改，本质上是相同的（递归到 $i$ 的时候 $\textit{mask}$ 是一样的），所以只有 $\mathcal{O}(k)$ 个本质不同的修改位置。每个位置至多有 $|\Sigma|=26$ 种不同的修改方式。因此，只有 $\mathcal{O}(k|\Sigma|)$ 个不同的 $\textit{mask}$。

有 $n$ 个不同的 $i$，每个 $i$ 有 $\mathcal{O}(k|\Sigma|)$ 个不同的 $\textit{mask}$，所以一共有 $\mathcal{O}(nk|\Sigma|)$ 个状态。

- 时间复杂度：$\mathcal{O}(nk|\Sigma|)$，其中 $n$ 是 $s$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。分类讨论：
   - 如果 $\textit{changed}=\texttt{false}$，没有修改过字母，那么分割方案是唯一的（注意我们只在集合大小大于 $k$ 时才分割），所以这样的状态一共有 $\mathcal{O}(n)$ 个，单个状态的计算时间为 $\mathcal{O}(|\Sigma|)$，枚举修改成什么字母。
   - 如果 $\textit{changed}=\texttt{true}$，根据上面的分析，这样的状态一共有 $\mathcal{O}(nk|\Sigma|)$ 个，单个状态的计算时间为 $\mathcal{O}(1)$。
   - 所以时间复杂度为 $\mathcal{O}(nk|\Sigma|)$。
- 空间复杂度：$\mathcal{O}(nk|\Sigma|)$。保存多少状态，就需要多少空间。

## 方法二：前后缀分解

枚举修改的字母是 $s[i]$，那么前缀 $[0,i-1]$ 和后缀 $[i+1,n-1]$ 就不能修改了。

如果预处理每个前缀 $[0,i-1]$ 的分割段数，以及每个后缀 $[i+1,n-1]$ 的分割段数，是不是就能快速地算出修改 $s[i]$ 后的分割段数呢？

前缀 $[0,i-1]$ 不能修改时的分割段数，可以从左到右遍历 $s$ 计算。

后缀 $[i+1,n-1]$ 呢？

我们需要想清楚两个问题：

1. 后缀可以独立计算吗？
2. 后缀可以从右到左遍历 $s$ 计算吗？如果可以，我们就能像计算前缀那样，计算后缀的分割段数。

对于第一个问题，设修改后，划分方式如下：

1. 把前缀 $[0,j]$ 分成若干段。
2. 下一段是 $[j+1,k]$，包含 $i$。
3. 接着，把后缀 $[k+1,n-1]$ 分成若干段。

由于后缀 $[k+1,n-1]$ 不包含 $i$，我们**可以独立计算** $[k+1,n-1]$ 的分割段数。

对于第二个问题，即证明如下定理。

**定理**：对于 $s$ 的任意后缀，从左到右分割出的段数，等于从右到左分割出的段数。

**证明**：考虑 $s$ 的某个后缀，假设从左到右分成了 $N$ 段，从左到右分别记作 $L_1,L_2,\ldots,L_N$，每一段最左边的字母下标分别记作 $p_1,p_2,\ldots,p_N$。再假设从右到左分成了 $M$ 段，**从右到左**分别记作 $R_1,R_2,\ldots,R_M$，每一段最左边（终点）的字母下标分别记作 $q_1,q_2,\ldots,q_M$，最右边（起点）的字母下标分别记作 $r_1,r_2,\ldots,r_M$。

![w379d.png](https://pic.leetcode.cn/1704712474-pevDnh-w379d.png)

对于 $R_1$，显然它包含 $L_N$，所以 $R_1$ 的终点 $q_1\le p_N$。如果 $q_1\le p_{N-1}$，那么 $R_1$ 会包含 $L_{N-1}$ 和 $L_N$，这是不可能的，因为这两段的字符种类是大于 $k$ 的。所以 $q_1 > p_{N-1}$。

这意味着对于 $R_2$，它的起点 $r_2$ 在 $L_{N-1}$ 中。同样地，$R_2$ 的终点 $q_2$ 必须大于 $p_{N-2}$，否则 $R_2$ 除了完整地包含 $L_{N-2}$，还包含 $L_{N-1}$ 的字母，这会让 $R_2$ 的字母种类数超过 $k$。

依此类推，每个 $R_i$ 的起点 $r_i$ 都在 $L_{N+1-i}$ 中。而最后一段 $R_M$ 的起点 $r_M$ 在 $L_1$ 中。这意味着 $N=M$。所以从左到右分割出的段数，等于从右到左分割出的段数。

### 算法

遍历 $s$，枚举要修改的字母 $s[i]$。

设从左到右分割到 $i-1$ 时，分割出了 $\textit{preSeg}$ 段，最新一段（记作 $L$）的字符集合为 $\textit{preMask}$，其大小为 $\textit{preSize}$。

设从右到左分割到 $i+1$ 时，分割出了 $\textit{sufSeg}$ 段，最新一段（记作 $R$）的字符集合为 $\textit{sufMask}$，其大小为 $\textit{sufSize}$。

设 $\textit{preMask}$ 和 $\textit{sufMask}$ 的并集大小为 $\textit{unionSize}$。

分类讨论：

- **情况 1**：如果 $\textit{unionSize}<k$，那么无论把 $s[i]$ 改成什么，并集的大小都不会超过 $k$，这意味着 $L$ 和 $R$ 必须合并成一段。
- **情况 2**：如果 $\textit{unionSize}<26$，并且 $\textit{preSize}$ 和 $\textit{sufSize}$ 都等于 $k$，那么把 $s[i]$ 改成一个既不在 $L$ 又不在 $R$ 中的字母，从而使 $L$ 和 $R$ 都是最大分割，这样就多出了一段。
- **情况 3**：不是情况 1 也不是情况 2，那么总段数不变。

代码实现时，可以先从右到左遍历 $s$，把 $\textit{sufSeg}$ 和 $\textit{sufMask}$ 都记录下来。然后从左到右遍历 $s$，计算 $\textit{preSeg}$ 和 $\textit{preMask}$，同时可以计算修改后的段数，更新答案的最大值。

特别地，如果 $k=26$，那么无论怎么改，都不能分割，只有一段。

```py [sol-Python3]
class Solution:
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        if k == 26:
            return 1

        seg, mask, size = 1, 0, 0
        def update(i: int) -> None:
            nonlocal seg, mask, size
            bit = 1 << (ord(s[i]) - ord('a'))
            if mask & bit:  # s[i] 已经在当前这一段中，无影响
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
        if ((mask & bit) != 0) { // c 已经在当前这一段中，无影响
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
        auto update = [&](int i) -> void {
            int bit = 1 << (s[i] - 'a');
            if (mask & bit) { // s[i] 已经在当前这一段中，无影响
                return;
            }
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
        seg = 1; mask = 0; size = 0;
        for (int i = 0; i < n; i++) {
            auto [suf_seg, suf_mask] = suf[i + 1];
            int res = seg + suf_seg; // 情况 3
            int union_size = popcount((uint32_t) mask | suf_mask);
            if (union_size < k) {
                res--; // 情况 1
            } else if (union_size < 26 && size == k && popcount((uint32_t) suf_mask) == k) {
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
		if mask&bit > 0 { // s[i] 已经在当前这一段中，无影响
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

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
