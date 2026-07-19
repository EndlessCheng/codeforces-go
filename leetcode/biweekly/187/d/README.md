设 $n$ 是 $\textit{source}$ 的长度。

从 $\textit{source}$ 的最后一个字母开始思考：

- 不替换，前提是 $\textit{source}[n-1] = \textit{target}[n-1]$。问题变成把 $\textit{source}$ 的子串 $[0,n-2]$ 变成 $\textit{target}$ 的子串 $[0,n-2]$ 的最小总成本。
- 替换，枚举替换规则。设 $\textit{pattern}$ 的长度为 $m$，如果 $\textit{pattern}$ 匹配 $\textit{source}$ 的子串 $[n-m,n-1]$，且 $\textit{replacement}$ 等于 $\textit{target}$ 的子串 $[n-m,n-1]$，那么问题变成把 $\textit{source}$ 的 $[0,n-m-1]$ 变成 $\textit{target}$ 的 $[0,n-m-1]$ 的最小总成本。

根据上述讨论，定义 $\textit{dfs}(i)$ 表示把 $\textit{source}$ 的子串 $[0,i]$ 变成 $\textit{target}$ 的子串 $[0,i]$ 的最小总成本。

用「选或不选」思考，讨论是否替换 $\textit{source}[i]$：

- 不替换，前提是 $\textit{source}[i] = \textit{target}[i]$。问题变成把 $\textit{source}$ 的子串 $[0,i-1]$ 变成 $\textit{target}$ 的子串 $[0,i-1]$ 的最小总成本，即 $\textit{dfs}(i-1)$。
- 替换，枚举替换规则。设 $\textit{pattern}$ 的长度为 $m$，如果 $\textit{pattern}$ 匹配 $\textit{source}$ 的子串 $[i-m+1,i]$，且 $\textit{replacement}$ 等于 $\textit{target}$ 的子串 $[i-m+1,i]$，那么问题变成把 $\textit{source}$ 的 $[0,i-m]$ 变成 $\textit{target}$ 的 $[0,i-m]$ 的最小总成本，即 $\textit{dfs}(i-m)$，再加上替换的成本，更新 $\textit{dfs}(i)$ 的最小值。
- 如果 $\textit{source}[i] \ne \textit{target}[i]$ 且无任何规则可以匹配，则 $\textit{dfs}(i) = \infty$。

**递归边界**：$\textit{dfs}(-1)=0$。此时剩余子串是空串，无需替换，成本为 $0$。

**递归入口**：$\textit{dfs}(n-1)$，这是原问题，也是答案。

[本题视频讲解](https://www.bilibili.com/video/BV1mJK66VEbN/?t=12m30s)，欢迎点赞关注~

## 写法一：记忆化搜索

关于记忆化搜索的原理，请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

**注**：题目没有保证 $\textit{source}$ 和 $\textit{target}$ 长度相等，我估计是出题人漏写了（试了一下，没有长度不等的数据）。

```py [sol-Python3]
class Solution:
    def minCost(self, source: str, target: str, rules: list[list[str]], costs: list[int]) -> int:
        # 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
        for i, (pattern, _) in enumerate(rules):
            costs[i] += pattern.count('*')

        # pattern 是否匹配 s[left: left+len(pattern)]
        def is_match(pattern: str, s: str, left: int) -> bool:
            for j, ch in enumerate(pattern):
                if ch != '*' and ch != s[left + j]:
                    return False
            return True

        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int) -> int:
            if i < 0:
                return 0

            # 不替换下标 i
            res = dfs(i - 1) if source[i] == target[i] else inf

            # 替换子串 [i-m+1, i]
            for (pattern, replacement), cost in zip(rules, costs):
                left = i - len(replacement) + 1  # 子串左端点
                if left >= 0 and replacement == target[left: i + 1] and is_match(pattern, source, left):
                    res = min(res, dfs(left - 1) + cost)

            return res

        ans = dfs(len(source) - 1)
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minCost(String source, String target, List<List<String>> Rules, int[] costs) {
        List<String>[] rules = Rules.toArray(List[]::new); // 转成数组更方便

        // 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
        for (int i = 0; i < rules.length; i++) {
            costs[i] += count(rules[i].get(0), '*');
        }

        int n = source.length();
        int[] memo = new int[n];
        Arrays.fill(memo, -1); // -1 表示该状态没有计算过

        int ans = dfs(n - 1, source, target, rules, costs, memo);
        return ans < Integer.MAX_VALUE / 2 ? ans : -1;
    }

    private int dfs(int i, String s, String t, List<String>[] rules, int[] costs, int[] memo) {
        if (i < 0) {
            return 0;
        }
        if (memo[i] != -1) { // 之前算过了
            return memo[i];
        }

        int res = Integer.MAX_VALUE / 2; // 防止加法溢出

        // 不替换下标 i
        if (s.charAt(i) == t.charAt(i)) {
            res = dfs(i - 1, s, t, rules, costs, memo);
        }

        // 替换子串 [i-m+1, i]
        for (int k = 0; k < rules.length; k++) {
            String replacement = rules[k].get(1);
            int left = i - replacement.length() + 1; // 子串左端点
            if (left >= 0 && replacement.equals(t.substring(left, i + 1)) && isMatch(rules[k].get(0), s, left)) {
                res = Math.min(res, dfs(left - 1, s, t, rules, costs, memo) + costs[k]);
            }
        }

        memo[i] = res; // 记忆化
        return res;
    }

    // pattern 是否匹配 s 的子串 [left, left+pattern.length()-1]
    private boolean isMatch(String pattern, String s, int left) {
        for (int j = 0; j < pattern.length(); j++) {
            char ch = pattern.charAt(j);
            if (ch != '*' && ch != s.charAt(left + j)) {
                return false;
            }
        }
        return true;
    }

    private int count(String s, char target) {
        int cnt = 0;
        for (char ch : s.toCharArray()) {
            if (ch == target) {
                cnt++;
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_match(const string& pattern, const string_view& s) {
        for (int j = 0; j < pattern.size(); j++) {
            if (pattern[j] != '*' && pattern[j] != s[j]) {
                return false;
            }
        }
        return true;
    }

public:
    int minCost(string source, string target, vector<vector<string>>& rules, vector<int>& costs) {
        // 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
        for (int i = 0; i < rules.size(); i++) {
            costs[i] += ranges::count(rules[i][0], '*');
        }

        string_view s(source); // string_view 的 substr 是 O(1) 的时间和空间
        string_view t(target);

        int n = s.size();
        vector<int> memo(n, INT_MAX / 2);

        auto dfs = [&](this auto&& dfs, int i) -> int {
            if (i < 0) {
                return 0;
            }

            int& res = memo[i]; // 注意这里是引用
            if (res != INT_MAX / 2) { // 之前算过了
                return res;
            }

            // 不替换下标 i
            if (s[i] == t[i]) {
                res = dfs(i - 1);
            }

            // 替换子串 [i-m+1, i]
            for (int k = 0; k < rules.size(); k++) {
                auto& replacement = rules[k][1];
                int left = i - replacement.size() + 1; // 子串左端点
                if (left >= 0 && replacement == t.substr(left, replacement.size()) && is_match(rules[k][0], s.substr(left))) {
                    res = min(res, dfs(left - 1) + costs[k]);
                }
            }

            return res;
        };

        int ans = dfs(n - 1);
        return ans < INT_MAX / 2 ? ans : -1;
    }
};
```

```go [sol-Go]
func minCost(source, target string, rules [][]string, costs []int) int {
	// 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
	for i, rule := range rules {
		costs[i] += strings.Count(rule[0], "*")
	}

	isMatch := func(pattern, s string) bool {
		for j, ch := range pattern {
			if ch != '*' && byte(ch) != s[j] {
				return false
			}
		}
		return true
	}

	n := len(source)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示该状态没有计算过
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前算过了
			return *p
		}

		res := math.MaxInt / 2 // 防止加法溢出

		// 不替换下标 i
		if source[i] == target[i] {
			res = dfs(i - 1)
		}

		// 替换子串 [i-m+1, i]
		for k, rule := range rules {
			replacement := rule[1]
			left := i - len(replacement) + 1 // 子串左端点
			if left >= 0 && replacement == target[left:i+1] && isMatch(rule[0], source[left:]) {
				res = min(res, dfs(left-1)+costs[k])
			}
		}

		*p = res // 记忆化
		return res
	}

	ans := dfs(n - 1)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
```

## 写法二：1:1 翻译成递推

$f[i+1]$ 对应 $\textit{dfs}(i)$。

$f[0] = 0$ 对应 $\textit{dfs}(-1) = 0$。

**最优性优化**：如果 $f[\textit{left}]+\textit{cost}\ge f[i+1]$，则不需要更新 $f[i+1]$，也就不需要花时间判断字符串是否匹配了。

```py [sol-Python3]
class Solution:
    def minCost(self, source: str, target: str, rules: list[list[str]], costs: list[int]) -> int:
        # 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
        for i, rule in enumerate(rules):
            costs[i] += rule[0].count('*')

        # pattern 是否匹配 s[left: left+len(pattern)]
        def is_match(pattern: str, s: str, left: int) -> bool:
            for j, ch in enumerate(pattern):
                if ch != '*' and ch != s[left + j]:
                    return False
            return True

        n = len(source)
        f = [0] + [inf] * n

        for i in range(n):
            # 不替换下标 i
            if source[i] == target[i]:
                f[i + 1] = f[i]

            # 替换子串 [i-m+1, i]
            for (pattern, replacement), cost in zip(rules, costs):
                left = i - len(replacement) + 1  # 子串左端点
                if left >= 0 and f[left] + cost < f[i + 1] and replacement == target[left: i + 1] and is_match(pattern, source, left):
                    f[i + 1] = f[left] + cost

        return f[n] if f[n] < inf else -1
```

```java [sol-Java]
class Solution {
    public int minCost(String s, String t, List<List<String>> Rules, int[] costs) {
        List<String>[] rules = Rules.toArray(List[]::new); // 转成数组更方便

        // 先加上 '*' 的个数，这样在 DP 中无需统计 '*' 的个数
        for (int i = 0; i < rules.length; i++) {
            costs[i] += count(rules[i].get(0), '*');
        }

        int n = s.length();
        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2);
        f[0] = 0;

        for (int i = 0; i < n; i++) {
            // 不替换下标 i
            if (s.charAt(i) == t.charAt(i)) {
                f[i + 1] = f[i];
            }

            // 替换子串 [i-m+1, i]
            for (int k = 0; k < rules.length; k++) {
                String replacement = rules[k].get(1);
                int left = i - replacement.length() + 1; // 子串左端点
                if (left >= 0 && f[left] + costs[k] < f[i + 1] && replacement.equals(t.substring(left, i + 1)) && isMatch(rules[k].get(0), s, left)) {
                    f[i + 1] = f[left] + costs[k];
                }
            }
        }

        return f[n] < Integer.MAX_VALUE / 2 ? f[n] : -1;
    }

    // pattern 是否匹配 s 的子串 [left, left+pattern.length()-1]
    private boolean isMatch(String pattern, String s, int left) {
        for (int j = 0; j < pattern.length(); j++) {
            char ch = pattern.charAt(j);
            if (ch != '*' && ch != s.charAt(left + j)) {
                return false;
            }
        }
        return true;
    }

    private int count(String s, char target) {
        int cnt = 0;
        for (char ch : s.toCharArray()) {
            if (ch == target) {
                cnt++;
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_match(const string& pattern, const string_view& s) {
        for (int j = 0; j < pattern.size(); j++) {
            if (pattern[j] != '*' && pattern[j] != s[j]) {
                return false;
            }
        }
        return true;
    }

public:
    int minCost(string source, string target, vector<vector<string>>& rules, vector<int>& costs) {
        // 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
        for (int i = 0; i < rules.size(); i++) {
            costs[i] += ranges::count(rules[i][0], '*');
        }

        string_view s(source); // string_view 的 substr 是 O(1) 的时间和空间
        string_view t(target);

        int n = s.size();
        vector<int> f(n + 1, INT_MAX / 2);
        f[0] = 0;

        for (int i = 0; i < n; i++) {
            // 不替换下标 i
            if (s[i] == t[i]) {
                f[i + 1] = f[i];
            }

            // 替换子串 [i-m+1, i]
            for (int k = 0; k < rules.size(); k++) {
                auto& replacement = rules[k][1];
                int left = i - replacement.size() + 1; // 子串左端点
                if (left >= 0 && f[left] + costs[k] < f[i + 1] && replacement == t.substr(left, replacement.size()) && is_match(rules[k][0], s.substr(left))) {
                    f[i + 1] = f[left] + costs[k];
                }
            }
        }

        return f[n] < INT_MAX / 2 ? f[n] : -1;
    }
};
```

```go [sol-Go]
func minCost(source, target string, rules [][]string, costs []int) int {
	// 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
	for i, rule := range rules {
		costs[i] += strings.Count(rule[0], "*")
	}

	isMatch := func(pattern, s string) bool {
		for j, ch := range pattern {
			if ch != '*' && byte(ch) != s[j] {
				return false
			}
		}
		return true
	}

	n := len(source)
	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt / 2 // 防止加法溢出

		// 不替换下标 i
		if source[i] == target[i] {
			res = f[i]
		}

		// 替换子串 [i-m+1, i]
		for k, rule := range rules {
			replacement := rule[1]
			left := i - len(replacement) + 1 // 子串左端点
			if left >= 0 && f[left]+costs[k] < res && replacement == target[left:i+1] && isMatch(rule[0], source[left:]) {
				res = f[left] + costs[k]
			}
		}

		f[i+1] = res
	}

	ans := f[n]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nL)$，其中 $n$ 是 $\textit{source}$ 的长度，$L$ 是所有 $\textit{pattern}_i$ 的长度之和。
- 空间复杂度：$\mathcal{O}(n)$。

## 更快的做法

由于 $\textit{replacement}_i$ 不含通配符，可以构建由 $\textit{replacement}_i$ 组成的 **AC 自动机**，从而快速求出每个 $\textit{replacement}_i$ 在 $\textit{target}$ 中的所有出现位置。

- [3213. 最小代价构造字符串](https://leetcode.cn/problems/construct-string-with-minimum-cost/)
- [3292. 形成目标字符串需要的最少字符串数 II](https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-ii/)

## 专题训练

见下面动态规划题单的「**§7.1 一维 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
