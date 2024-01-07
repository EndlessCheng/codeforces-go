**前置知识**：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

定义 $\textit{dfs}(i,\textit{mask}, \textit{changed})$ 表示当前遍历到 $s[i]$，当前这一段的字符集合是 $\textit{mask}$，是否已经修改了字符（$\textit{changed}$），后续可以得到的最大分割数。

分类讨论：

- 如果不改 $s[i]$：
   - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $s[i]$ 必须划分到下一段子串中。答案为 $\textit{dfs}(i+1, \{s[i]\},\textit{changed}) + 1$。
   - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $s[i]$ 必须在当前这一段中。答案为 $\textit{dfs}(i+1, \textit{mask}\cup \{s[i]\},\textit{changed})$。
- 如果 $\textit{changed}=\texttt{false}$，那么可以改 $s[i]$，枚举改成第 $j$ 个字母。
    - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $j$ 必须划分到下一段子串中。答案为 $\textit{dfs}(i+1, \{j\},\texttt{true}) + 1$。
    - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $j$ 必须在当前这一段中。答案为 $\textit{dfs}(i+1, \textit{mask}\cup \{j\},\texttt{true})$。

这些情况取最大值，就得到了 $\textit{dfs}(i,\textit{mask}, \textit{changed})$。

递归边界：$\textit{dfs}(n,*,*) = 1$。注意当 $i>0$ 时，$\textit{mask}\ne 0$，表示一段子串的字符集合。所以递归到 $i=n$ 时，$\textit{mask}$ 就是最后一段的字符集合了，返回 $1$。

递归入口：$\textit{dfs}(0,0,\texttt{false})$，也就是答案。

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
public class Solution {
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
        function<int(int, int, bool)> dfs = [&](int i, int mask, bool changed) -> int {
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

我们可以从 $i$ 开始向左扩展，每遇到一个字符，$\textit{mask}$ 集合要么不变，要么增加一个字符。

所以到 $i$ 为止的 $\textit{mask}$ 至多有 $|\Sigma|$ 个不同的值。这里 $|\Sigma|=26$。

所以只有 $\mathcal{O}(n|\Sigma|)$ 个状态。

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n|\Sigma|)$，单个状态的计算时间为 $\mathcal{O}(|\Sigma|)$，所以动态规划的时间复杂度为 $\mathcal{O}(n|\Sigma|^2)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
