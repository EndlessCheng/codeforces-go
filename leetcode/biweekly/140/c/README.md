本题可以修改一个字母，推荐先完成**不修改版本** [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/)（[我的题解](https://leetcode.cn/problems/subsequence-with-the-minimum-score/solutions/2107010/qian-hou-zhui-fen-jie-san-zhi-zhen-pytho-6cmr/)）。

做完 2565 后，你知道本题也可以用前后缀分解，但难点在于计算字典序最小的下标序列。

为方便描述，下文把 $\textit{word}_1$ 记作 $s$，把 $\textit{word}_2$ 记作 $t$。

定义 $\textit{suf}[i]$ 为 $s[i:]$ 对应的 $t$ 的最长后缀的开始下标 $j$，即 $t[j:]$ 是 $s[i:]$ 的子序列。

预处理 $\textit{suf}$，然后从左到右遍历 $s$，分类讨论：

- 如果 $s[i]=t[j]$，既然能匹配上，那么就立刻匹配，直接把 $i$ 加入答案。如果不匹配，可能后面就没机会找到子序列了，或者答案的第 $j$ 个下标比 $i$ 大，不是字典序最小的下标序列。
- 如果 $s[i]\ne t[j]$ 且 $\textit{suf}[i+1]\le j+1$，说明修改 $s[i]$ 为 $t[j]$ 后，$t[j+1:]$ 是 $s[i+1:]$ 的子序列。此时**一定要修改**，如果不修改，那么答案的第 $j$ 个下标比 $i$ 大，不是字典序最小的下标序列。
- 修改后，继续向后匹配，在 $s[i]=t[j]$ 时把 $i$ 加入答案。

循环中，如果发现 $j$ 等于 $t$ 的长度，说明匹配完成，立刻返回答案。

如果循环中没有返回，那么循环结束后返回空数组。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1bjxyewEQV/) 第三题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def validSequence(self, s: str, t: str) -> List[int]:
        n, m = len(s), len(t)
        suf = [0] * (n + 1)
        suf[n] = m
        j = m - 1
        for i in range(n - 1, -1, -1):
            if j >= 0 and s[i] == t[j]:
                j -= 1
            suf[i] = j + 1

        ans = []
        changed = False  # 是否修改过
        j = 0
        for i, c in enumerate(s):
            if c == t[j] or not changed and suf[i + 1] <= j + 1:
                if c != t[j]:
                    changed = True
                ans.append(i)
                j += 1
                if j == m:
                    return ans
        return []
```

```java [sol-Java]
class Solution {
    public int[] validSequence(String word1, String word2) {
        char[] s = word1.toCharArray();
        char[] t = word2.toCharArray();
        int n = s.length;
        int m = t.length;

        int[] suf = new int[n + 1];
        suf[n] = m;
        int j = m - 1;
        for (int i = n - 1; i >= 0; i--) {
            if (j >= 0 && s[i] == t[j]) {
                j--;
            }
            suf[i] = j + 1;
        }

        int[] ans = new int[m];
        boolean changed = false; // 是否修改过
        j = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == t[j] || !changed && suf[i + 1] <= j + 1) {
                if (s[i] != t[j]) {
                    changed = true;
                }
                ans[j++] = i;
                if (j == m) {
                    return ans;
                }
            }
        }
        return new int[]{};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> validSequence(string s, string t) {
        int n = s.length(), m = t.length();
        vector<int> suf(n + 1);
        suf[n] = m;
        for (int i = n - 1, j = m - 1; i >= 0; i--) {
            if (j >= 0 && s[i] == t[j]) {
                j--;
            }
            suf[i] = j + 1;
        }

        vector<int> ans(m);
        bool changed = false; // 是否修改过
        for (int i = 0, j = 0; i < n; i++) {
            if (s[i] == t[j] || !changed && suf[i + 1] <= j + 1) {
                if (s[i] != t[j]) {
                    changed = true;
                }
                ans[j++] = i;
                if (j == m) {
                    return ans;
                }
            }
        }
        return {};
    }
};
```

```go [sol-Go]
func validSequence(s, t string) []int {
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = j + 1
	}

	ans := make([]int, m)
	changed := false // 是否修改过
	j := 0
	for i := range s {
		if s[i] == t[j] || !changed && suf[i+1] <= j+1 {
			if s[i] != t[j] {
				changed = true
			}
			ans[j] = i
			j++
			if j == m {
				return ans
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果改的是 $t$ 中的字母呢？

欢迎在评论区分享你的思路/代码。

## 相似题目

- [双指针题单](https://leetcode.cn/circle/discuss/0viNMK/) 中的「**§4.2 判断子序列**」。
- [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**专题：前后缀分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
