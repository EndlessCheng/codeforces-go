## 子序列 DP 的思考套路

- 子序列 + 不考虑相邻元素：**选或不选**。代表题目：[494. 目标和（0-1 背包）](https://leetcode.cn/problems/target-sum/)
- 子序列 + 考虑相邻元素：**枚举选哪个**。代表题目：[300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)

## 本题思路

本题属于「子序列 + 考虑相邻元素」，用枚举选哪个解决，状态定义类似最长递增子序列。

定义 $f[i]$ 表示从 $i$ 到 $n-1$ 中，我们选出的最长子序列的长度（第一个下标一定是 $i$）。定义成后缀是为了方便后面输出具体方案。

初始值 $f[i]=1$，表示选择它自己作为子序列。

如果 $\textit{groups}[j] \ne \textit{groups}[i]$ 并且 $\textit{words}[j]$ 和 $\textit{words}[i]$ 满足题目要求，并且 $f[j]+1 > f[i]$，那么更新

$$
f[i] = f[j] + 1
$$

并且记录转移来源 $\textit{from}[i] = j$。

那么最长子序列的长度就是 $\max(f)$。

## 如何输出方案

设 $\textit{mx}$ 是 $\max(f)$ 在 $f$ 中的下标，即 $f[\textit{mx}]=\max(f)$。如果有多个这样的下标，随便取哪个都行。

从 $\textit{mx}$ 开始不断循环，每次把 $\textit{words}[mx]$ 加入答案，然后更新

$$
mx = \textit{from}[mx]
$$

表示顺着转移来源往右走。

当找到了 $\max(f)$ 个字符串时停止循环。

```py [sol-Python3]
class Solution:
    def getWordsInLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        def ok(s: str, t: str) -> bool:
            return len(s) == len(t) and sum(x != y for x, y in zip(s, t)) == 1

        n = len(words)
        f = [0] * n
        from_idx = [0] * n
        mx = n - 1
        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                if f[j] > f[i] and groups[j] != groups[i] and ok(words[i], words[j]):
                    f[i] = f[j]
                    from_idx[i] = j
            f[i] += 1  # 加一写在这里
            if f[i] > f[mx]:
                mx = i

        ans = [''] * f[mx]
        for i in range(f[mx]):
            ans[i] = words[mx]
            mx = from_idx[mx]
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> getWordsInLongestSubsequence(String[] words, int[] groups) {
        int n = words.length;
        int[] f = new int[n];
        int[] from = new int[n];
        int mx = n - 1;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = i + 1; j < n; j++) {
                if (f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j])) {
                    f[i] = f[j];
                    from[i] = j;
                }
            }
            f[i]++; // 加一写在这里
            if (f[i] > f[mx]) {
                mx = i;
            }
        }

        int m = f[mx];
        List<String> ans = new ArrayList<>(m);
        for (int i = 0; i < m; i++) {
            ans.add(words[mx]);
            mx = from[mx];
        }
        return ans;
    }

    private static boolean ok(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        boolean diff = false;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) != t.charAt(i)) {
                if (diff) {
                    return false;
                }
                diff = true;
            }
        }
        return diff;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool ok(string& s, string& t) {
        if (s.length() != t.length()) {
            return false;
        }
        bool diff = false;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] != t[i]) {
                if (diff) {
                    return false;
                }
                diff = true;
            }
        }
        return diff;
    }

public:
    vector<string> getWordsInLongestSubsequence(vector<string>& words, vector<int>& groups) {
        int n = words.size();
        vector<int> f(n);
        vector<int> from(n);
        int mx = n - 1;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = i + 1; j < n; j++) {
                if (f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j])) {
                    f[i] = f[j];
                    from[i] = j;
                }
            }
            f[i]++; // 加一写在这里
            if (f[i] > f[mx]) {
                mx = i;
            }
        }

        int m = f[mx];
        vector<string> ans(m);
        for (int i = 0; i < m; i++) {
            ans[i] = words[mx];
            mx = from[mx];
        }
        return ans;
    }
};
```

```go [sol-Go]
func ok(s, t string) (diff bool) {
	if len(s) != len(t) {
		return
	}
	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return
}

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	f := make([]int, n)
	from := make([]int, n)
	mx := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j]) {
				f[i] = f[j]
				from[i] = j
			}
		}
		f[i]++ // 加一写在这里
		if f[i] > f[mx] {
			mx = i
		}
	}

	ans := make([]string, f[mx])
	for i := range ans {
		ans[i] = words[mx]
		mx = from[mx]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2L)$，其中 $L$ 为 $\textit{words}[i]$ 的长度，至多为 $10$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
