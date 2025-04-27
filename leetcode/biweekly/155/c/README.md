## 前置知识

1. [KMP 原理讲解](https://www.zhihu.com/question/21923021/answer/37475572)。
2. [差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

## 思路

对于水平子串，把 $\textit{grid}$ 每一行首尾相接，得到一个长为 $mn$ 的文本串 $\textit{text}$，我们需要在文本串 $\textit{text}$ 中找到所有模式串 $\textit{pattern}$ 的出现位置，这正是 **KMP 算法**的标准应用。

标记所有在 $\textit{pattern}$ 中的单元格。假设我们在 $\textit{text}[i]$ 处匹配完成，那么用**差分数组**把下标区间 $[i-k+1,i]$ 加一，其中 $k$ 是 $\textit{pattern}$ 的长度。计算差分数组的前缀和，值大于 $0$ 的下标就对应着在 $\textit{pattern}$ 中的单元格。

对于垂直子串，计算方法同理。可以把相关逻辑封装成一个函数，方便垂直子串**复用**。

## 细节

对于水平子串，我们计算差分数组的前缀和，得到一个长为 $mn$ 的数组 $\textit{inPatternH}$。

对于垂直子串，我们计算差分数组的前缀和，得到一个长为 $mn$ 的数组 $\textit{inPatternV}$。

如果 $\textit{inPatternH}[i]>0$，则表示单元格 $\textit{grid}[\left\lfloor i/n \right\rfloor][i\bmod n]$ 在 $\textit{pattern}$ 中。

单元格 $\textit{grid}[\left\lfloor i/n \right\rfloor][i\bmod n]$ 在 $\textit{inPatternV}$ 中的哪个位置？由于垂直子串是竖着扫描的，基于一个 $n$ 行 $m$ 列的矩阵，$\textit{grid}[\left\lfloor i/n \right\rfloor][i\bmod n]$ 在这个矩阵的 $i\bmod n$ 行 $\left\lfloor i/n \right\rfloor$ 列，所以对应到 $\textit{inPatternV}$ 中的下标为

$$
(i\bmod n)\cdot m + \left\lfloor i/n \right\rfloor
$$

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def calc_pi(self, pattern: str) -> List[int]:
        pi = [0] * len(pattern)
        cnt = 0
        for i in range(1, len(pi)):
            b = pattern[i]
            while cnt > 0 and pattern[cnt] != b:
                cnt = pi[cnt - 1]
            if pattern[cnt] == b:
                cnt += 1
            pi[i] = cnt
        return pi

    def kmp_search(self, text: List[str], pattern: str, pi: List[int]) -> List[int]:
        n, k = len(text), len(pattern)
        diff = [0] * (n + 1)
        cnt = 0
        for i, b in enumerate(text):
            while cnt > 0 and pattern[cnt] != b:
                cnt = pi[cnt - 1]
            if pattern[cnt] == b:
                cnt += 1
            if cnt == k:
                diff[i - k + 1] += 1
                diff[i + 1] -= 1
                cnt = pi[cnt - 1]
        return list(accumulate(diff[:n]))

    def countCells(self, grid: List[List[str]], pattern: str) -> int:
        h_text = [c for row in grid for c in row]
        v_text = [c for col in zip(*grid) for c in col]

        pi = self.calc_pi(pattern)
        in_pattern_h = self.kmp_search(h_text, pattern, pi)
        in_pattern_v = self.kmp_search(v_text, pattern, pi)

        m, n = len(grid), len(grid[0])
        ans = 0
        for i, in_h in enumerate(in_pattern_h):
            if in_h and in_pattern_v[i % n * m + i // n]:
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countCells(char[][] grid, String pattern) {
        int m = grid.length;
        int n = grid[0].length;

        char[] hText = new char[m * n];
        int idx = 0;
        for (char[] row : grid) {
            for (char c : row) {
                hText[idx++] = c;
            }
        }

        char[] vText = new char[m * n];
        idx = 0;
        for (int j = 0; j < n; j++) {
            for (char[] row : grid) {
                vText[idx++] = row[j];
            }
        }

        char[] pat = pattern.toCharArray();
        int[] pi = calcPi(pat);
        int[] inPatternH = kmpSearch(hText, pat, pi);
        int[] inPatternV = kmpSearch(vText, pat, pi);

        int ans = 0;
        for (int i = 0; i < m * n; i++) {
            if (inPatternH[i] > 0 && inPatternV[i % n * m + i / n] > 0) {
                ans++;
            }
        }
        return ans;
    }

    private int[] calcPi(char[] pattern) {
        int[] pi = new int[pattern.length];
        int match = 0;
        for (int i = 1; i < pi.length; i++) {
            char b = pattern[i];
            while (match > 0 && pattern[match] != b) {
                match = pi[match - 1];
            }
            if (pattern[match] == b) {
                match++;
            }
            pi[i] = match;
        }
        return pi;
    }

    private int[] kmpSearch(char[] text, char[] pattern, int[] pi) {
        int n = text.length;
        int[] diff = new int[n + 1];
        int match = 0;
        for (int i = 0; i < n; i++) {
            char b = text[i];
            while (match > 0 && pattern[match] != b) {
                match = pi[match - 1];
            }
            if (pattern[match] == b) {
                match++;
            }
            if (match == pi.length) {
                diff[i - pi.length + 1]++;
                diff[i + 1]--;
                match = pi[match - 1];
            }
        }
        for (int i = 1; i < n; i++) {
            diff[i] += diff[i - 1];
        }
        return diff;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc_pi(string& pattern) {
        vector<int> pi(pattern.size());
        int match = 0;
        for (int i = 1; i < pi.size(); i++) {
            char b = pattern[i];
            while (match > 0 && pattern[match] != b) {
                match = pi[match - 1];
            }
            if (pattern[match] == b) {
                match++;
            }
            pi[i] = match;
        }
        return pi;
    }

    vector<int> kmp_search(string& text, string& pattern, vector<int>& pi) {
        int n = text.size(), k = pattern.size();
        vector<int> diff(n + 1);
        int match = 0;
        for (int i = 0; i < n; i++) {
            int b = text[i];
            while (match > 0 && pattern[match] != b) {
                match = pi[match - 1];
            }
            if (pattern[match] == b) {
                match++;
            }
            if (match == k) {
                diff[i - k + 1]++;
                diff[i + 1]--;
                match = pi[match - 1];
            }
        }
        partial_sum(diff.begin(), diff.end(), diff.begin());
        return diff;
    }

public:
    int countCells(vector<vector<char>>& grid, string pattern) {
        int m = grid.size(), n = grid[0].size();
        string h_text;
        for (auto& row : grid) {
            h_text.insert(h_text.end(), row.begin(), row.end());
        }
        string v_text;
        for (int j = 0; j < n; j++) {
            for (auto& row : grid) {
                v_text += row[j];
            }
        }

        vector<int> pi = calc_pi(pattern);
        vector<int> in_pattern_h = kmp_search(h_text, pattern, pi);
        vector<int> in_pattern_v = kmp_search(v_text, pattern, pi);

        int ans = 0;
        for (int i = 0; i < m * n; i++) {
            if (in_pattern_h[i] && in_pattern_v[i % n * m + i / n]) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func calcPi(pattern string) []int {
	pi := make([]int, len(pattern))
	match := 0
	for i := 1; i < len(pi); i++ {
		b := pattern[i]
		for match > 0 && pattern[match] != b {
			match = pi[match-1]
		}
		if pattern[match] == b {
			match++
		}
		pi[i] = match
	}
	return pi
}

func kmpSearch(text []byte, pattern string, pi []int) []int {
	n := len(text)
	diff := make([]int, n+1)
	match := 0
	for i, b := range text {
		for match > 0 && pattern[match] != b {
			match = pi[match-1]
		}
		if pattern[match] == b {
			match++
		}
		if match == len(pi) {
			diff[i-len(pi)+1]++
			diff[i+1]--
			match = pi[match-1]
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff[:n]
}

func countCells(grid [][]byte, pattern string) (ans int) {
	m, n := len(grid), len(grid[0])
	hText := slices.Concat(grid...)
	vText := make([]byte, 0, m*n)
	for j := range n {
		for _, row := range grid {
			vText = append(vText, row[j])
		}
	}

	pi := calcPi(pattern)
	inPatternH := kmpSearch(hText, pattern, pi)
	inPatternV := kmpSearch(vText, pattern, pi)

	for i, x := range inPatternH {
		if x > 0 && inPatternV[i%n*m+i/n] > 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

更多相似题目，见下面字符串题单的「**一、KMP**」和数据结构题单的「**§2.1 一维差分**」。

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
