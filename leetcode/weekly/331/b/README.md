如果 $\textit{words}[i]$ 符合要求，视作 $1$，否则视作 $0$。计算这个 01 数组的**前缀和**数组 $s$，就可以 $\mathcal{O}(1)$ 回答每个询问了。

关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

```py [sol-Python3]
class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        is_valid = lambda s: s[0] in "aeiou" and s[-1] in "aeiou"
        s = list(accumulate(map(is_valid, words), initial=0))
        return [s[r + 1] - s[l] for l, r in queries]
```

```java [sol-Java]
class Solution {
    private static final String VOWEL = "aeiou";

    public int[] vowelStrings(String[] words, int[][] queries) {
        int[] sum = new int[words.length + 1];
        for (int i = 0; i < words.length; i++) {
            String w = words[i];
            sum[i + 1] = sum[i];
            if (VOWEL.indexOf(w.charAt(0)) != -1 && VOWEL.indexOf(w.charAt(w.length() - 1)) != -1) {
                sum[i + 1]++;
            }
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            ans[i] = sum[q[1] + 1] - sum[q[0]];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> vowelStrings(vector<string>& words, vector<vector<int>>& queries) {
        const string VOWEL = "aeiou";
        vector<int> sum(words.size() + 1);
        for (int i = 0; i < words.size(); i++) {
            auto& w = words[i];
            sum[i + 1] = sum[i];
            if (VOWEL.find(w[0]) != string::npos && VOWEL.find(w.back()) != string::npos) {
                sum[i + 1]++;
            }
        }

        vector<int> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            auto& q = queries[i];
            ans[i] = sum[q[1] + 1] - sum[q[0]];
        }
        return ans;
    }
};
```

```go [sol-Go]
func vowelStrings(words []string, queries [][]int) []int {
	sum := make([]int, len(words)+1)
	for i, w := range words {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", w[:1]) && strings.Contains("aeiou", w[len(w)-1:]) {
			sum[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] - sum[q[0]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{words}$ 的长度, $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**一、前缀和**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
