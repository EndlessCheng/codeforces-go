## 前置知识：动态规划

请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://b23.tv/72onpYq)

## 子序列 DP 的思考套路

- 子序列 + 不考虑相邻元素：选或不选。代表题目：[494. 目标和（0-1 背包）](https://leetcode.cn/problems/target-sum/)
- 子序列 + 考虑相邻元素：枚举选哪个。代表题目：[300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)

## 本题思路

本题属于「子序列 + 考虑相邻元素」，用枚举选哪个解决，状态定义类似最长递增子序列。

定义 $f[i]$ 表示从 $i$ 到 $n-1$ 中，我们选出的最长子序列的长度。定义成后缀是为了方便后面输出具体方案。

初始值 $f[i]=1$，表示选择它自己作为子序列。

如果 $\textit{groups}[j] \ne \textit{groups}[i]$ 并且 $\textit{words}[j]$ 和 $\textit{words}[i]$ 满足题目要求，并且 $f[j]+1 > f[i]$，那么更新

$$
f[i] = f[j] + 1
$$

并且记录转移来源 $\textit{from}[i] = j$。

那么最长子序列的长度就是 $\max(f)$。

## 如何输出方案

设 $\textit{mx}$ 是 $\max(f)$ 的下标，即 $f[\textit{mx}]=\max(f)$。

从 $\textit{mx}$ 开始不断循环，每次把 $\textit{words}[mx]$ 加入答案，然后更新

$$
mx = \textit{from}[mx]
$$

表示顺着转移来源往右走。

当找到了 $\max(f)$ 个字符串时停止循环。

```py [sol-Python3]
class Solution:
    def getWordsInLongestSubsequence(self, n: int, words: List[str], groups: List[int]) -> List[str]:
        def ok(s: str, t: str) -> bool:
            return len(s) == len(t) and sum(x != y for x, y in zip(s, t)) == 1

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
    public List<String> getWordsInLongestSubsequence(int n, String[] words, int[] groups) {
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
        var diff = false;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) != t.charAt(i)) {
                if (diff) return false;
                diff = true;
            }
        }
        return diff;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool ok(string &s, string &t) {
        if (s.length() != t.length()) {
            return false;
        }
        bool diff = false;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] != t[i]) {
                if (diff) return false;
                diff = true;
            }
        }
        return diff;
    }
public:
    vector<string> getWordsInLongestSubsequence(int n, vector<string>& words, vector<int>& groups) {
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

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
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
